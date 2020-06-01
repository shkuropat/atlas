// Copyright 2020 The Atlas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package atlas

import (
	"fmt"
	"github.com/ulikunitz/xz/lzma"
	"io"

	log "github.com/sirupsen/logrus"
)

// DataChunkTransport defines transport level interface (for both client and server),
// which serves DataChunk streams bi-directionally.
type DataChunkTransport interface {
	Send(*DataChunk) error
	Recv() (*DataChunk, error)
}

// DataChunkFile is a handler to set of DataChunk(s)
// Inspired by os.File handler and is expected to be used in the same context.
// DataChunkFile implements the following interfaces:
//	- io.Writer
//	- io.WriterTo
// 	- io.ReaderFrom
//	- io.Closer
// and thus can be used in any functions, which operate these interfaces, such as io.Copy()
type DataChunkFile struct {
	transport DataChunkTransport

	// Header is mandatory
	Header *Metadata
	// TransportMetadata is optional
	TransportMetadata *Metadata
	// PayloadMetadata is optional
	PayloadMetadata *Metadata

	// initialOffset of these chunks
	initialOffset int64
	// currentOffset of the current chunk within file
	currentOffset int64
	// offset of the current chunk within "all chunks" = initialOffset + currentOffset
	offset int64

	// MaxWriteChunkSize limits max size of a payload within one data chunk to be sent
	MaxWriteChunkSize int

	lzmaWriter *lzma.Writer

	// Receive part
	buf []byte
	err error
}

// OpenDataChunkFile opens set of DataChunk(s)
// Inspired by os.OpenFile()
func OpenDataChunkFile(transport DataChunkTransport) (*DataChunkFile, error) {
	return &DataChunkFile{
		transport: transport,
	}, nil
}

// ensureTransportMetadata ensures DataChunkFile has TransportMetadata in place
func (f *DataChunkFile) ensureTransportMetadata() {
	if f.TransportMetadata == nil {
		f.TransportMetadata = new(Metadata)
	}
}

// HasTransportMetadata
func (f *DataChunkFile) HasTransportMetadata() bool {
	return f.TransportMetadata != nil
}

// ensurePayloadMetadata ensures DataChunkFile has PayloadMetadata in place
func (f *DataChunkFile) ensurePayloadMetadata() {
	if f.PayloadMetadata == nil {
		f.PayloadMetadata = new(Metadata)
	}
}

// HasPayloadMetadata
func (f *DataChunkFile) HasPayloadMetadata() bool {
	return f.PayloadMetadata != nil
}

// ensureBuf ensures DataChunkFile has buf in place
func (f *DataChunkFile) ensureBuf() {
	if f.buf == nil {
		f.buf = []byte{}
	}
}

// close
func (f *DataChunkFile) close() {
	f.Header = nil
	f.TransportMetadata = nil
	f.PayloadMetadata = nil

	f.initialOffset = 0
	f.currentOffset = 0
	f.offset = 0

	f.buf = nil
	f.err = nil
}

// acceptHeader
func (f *DataChunkFile) acceptHeader(dataChunk *DataChunk) {
	hdr := dataChunk.GetHeader()
	if hdr == nil {
		return
	}

	f.Header = hdr
}

// acceptTransportMetadata
func (f *DataChunkFile) acceptTransportMetadata(dataChunk *DataChunk) {
	md := dataChunk.GetTransportMetadata()
	if md == nil {
		return
	}

	f.TransportMetadata = md
}

// acceptPayloadMetadata
func (f *DataChunkFile) acceptPayloadMetadata(dataChunk *DataChunk) {
	md := dataChunk.GetPayloadMetadata()
	if md == nil {
		return
	}

	f.PayloadMetadata = md
}

// acceptMeta
func (f *DataChunkFile) acceptMeta(dataChunk *DataChunk) {
	f.acceptTransportMetadata(dataChunk)
	f.acceptPayloadMetadata(dataChunk)
	f.acceptHeader(dataChunk)
}

// logDataChunk
func (f *DataChunkFile) logDataChunk(dataChunk *DataChunk) {
	// Fetch filename from the chunks stream - it may be in any chunk, actually
	filename := "not specified"
	if dataChunk.HasPayloadMetadata() {
		if dataChunk.GetPayloadMetadata().HasFilename() {
			filename = dataChunk.GetPayloadMetadata().GetFilename()
		}
	}

	// Fetch offset of this chunk within the stream
	offset := "not specified"
	if dataChunk.Header.HasOffset() {
		offset = fmt.Sprintf("%d", dataChunk.Header.GetOffset())
	}

	// How many bytes do we have in this chunk?
	_len := len(dataChunk.GetBytes())
	log.Infof("Data.Recv() got msg. filename: '%s', chunk len: %d, chunk offset: %s, last chunk: %v",
		filename,
		_len,
		offset,
		dataChunk.Header.GetLast(),
	)
	fmt.Printf("%s\n", string(dataChunk.GetBytes()))
}

// recvDataChunk
func (f *DataChunkFile) recvDataChunk() (*DataChunk, error) {
	log.Infof("DataChunkFile.recvDataChunk() - start")
	defer log.Infof("DataChunkFile.recvDataChunk() - end")

	// Whether this chunk is the last one within this DataChunkFile
	dataChunk, err := f.transport.Recv()
	if dataChunk != nil {
		f.acceptMeta(dataChunk)
		f.logDataChunk(dataChunk)
	}

	if err == nil {
		// All went well, ready to receive more data
	} else if err == io.EOF {
		// Correct EOF arrived
		log.Infof("DataChunkTransport.Recv() get EOF")
	} else {
		// Stream broken
		log.Infof("DataChunkTransport.Recv() got err: %v", err)
	}

	return dataChunk, err
}

// appendDataBuf
func (f *DataChunkFile) appendDataBuf() {
	dataChunk, err := f.recvDataChunk()
	if dataChunk != nil {
		if len(dataChunk.GetBytes()) > 0 {
			f.ensureBuf()
			f.buf = append(f.buf, dataChunk.GetBytes()...)
		}

		if dataChunk.Header.GetLast() {
			f.err = io.EOF
		}
	}

	if err != nil {
		f.err = err
	}
}

// sendDataChunk
func (f *DataChunkFile) sendDataChunk(p []byte) (n int, err error) {
	log.Infof("DataChunkFile.sendDataChunk() - start")
	defer log.Infof("DataChunkFile.sendDataChunk() - end")

	if (len(p) > f.MaxWriteChunkSize) && (f.MaxWriteChunkSize > 0) {
		return 0, fmt.Errorf("attempt to sendDataChunk() with oversized chunk: %d > %d", len(p), f.MaxWriteChunkSize)
	}

	n = len(p)
	var transportMD *Metadata
	var payloadMD *Metadata
	if f.currentOffset == 0 {
		// First chunk in this file, it may have some metadata
		transportMD = f.TransportMetadata
		payloadMD = f.PayloadMetadata
	}
	f.offset = f.initialOffset + f.currentOffset
	chunk := NewDataChunk(transportMD, payloadMD, &f.offset, false, p)
	err = f.transport.Send(chunk)
	if err != nil {
		// We have some kind of error and were not able to send the data
		n = 0
		if err == io.EOF {
			// Not sure, probably this is not an error, after all
			log.Infof("Send() received EOF")
		} else {
			log.Errorf("failed to Send() %v", err)
		}
	}

	f.currentOffset += int64(n)

	return
}

// Implements io.Writer
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
func (f *DataChunkFile) Write(p []byte) (n int, err error) {
	log.Infof("DataChunkFile.Write() - start")
	defer log.Infof("DataChunkFile.Write() - end")

	if (len(p) < f.MaxWriteChunkSize) || (f.MaxWriteChunkSize <= 0) {
		// No need to chunk
		return f.sendDataChunk(p)
	}

	// Need to chunk
	n = 0
	for offset := 0; offset < len(p); offset += f.MaxWriteChunkSize {
		if sent, e := f.sendDataChunk(p[offset:f.MaxWriteChunkSize]); e != nil {
			return n, e
		} else {
			n += sent
		}
	}

	return
}

// Implements io.WriterTo
//
// WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to dst until there's no more data to write or
// when an error occurs. The return value n is the number of bytes
// written. Any error encountered during the write is also returned.
//
// The Copy function uses WriterTo if available.
func (f *DataChunkFile) WriteTo(dst io.Writer) (n int64, err error) {
	log.Infof("DataChunkFile.WriteTo() - start")
	defer log.Infof("DataChunkFile.WriteTo() - end")

	n = 0
	for {
		var dataChunk *DataChunk
		dataChunk, err = f.recvDataChunk()
		if dataChunk != nil {
			n += int64(len(dataChunk.GetBytes()))

			// TODO need to handle write errors
			_, _ = dst.Write(dataChunk.GetBytes())

			if dataChunk.Header.GetLast() {
				return
			}
		}

		if err != nil {
			return
		}
	}
}

// Implements io.Reader
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
func (f *DataChunkFile) Read(p []byte) (n int, err error) {
	log.Infof("DataChunkFile.Read() - start")
	defer log.Infof("DataChunkFile.Read() - end")

	if len(f.buf) == 0 {
		f.appendDataBuf()
	}

	n = 0
	if len(f.buf) > 0 {
		n = copy(p, f.buf)
		f.buf = f.buf[n:]

		if len(f.buf) > 0 {
			return n, nil
		}
	}

	return n, f.err
}

// Implements io.ReaderFrom
//
// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from src until EOF or error.
// The return value n is the number of bytes read.
// Any error except io.EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
func (f *DataChunkFile) ReadFrom(src io.Reader) (n int64, err error) {
	log.Infof("DataChunkFile.ReadFrom() - start")
	defer log.Infof("DataChunkFile.ReadFrom() - end")

	bufSize := 1024
	if f.MaxWriteChunkSize > 0 {
		bufSize = f.MaxWriteChunkSize
	}

	n = 0
	p := make([]byte, bufSize)
	for {
		read, readErr := src.Read(p)
		n += int64(read)
		if read > 0 {
			// We've got some data and it has to be processed no matter what
			// Write these bytes to data stream
			_, writeErr := f.Write(p[:read])
			if writeErr != nil {
				err = writeErr
				return
			}
		}

		if readErr != nil {
			// We have some kind of an error
			if readErr == io.EOF {
				// However, EOF is not an error on read
				readErr = nil
			}
			// In case of an error or EOF, break the read loop
			return
		}
	}
}

// Implements io.Closer
//
// Closer is the interface that wraps the basic Close method.
//
// The behavior of Close after the first call is undefined.
// Specific implementations may document their own behavior.
func (f *DataChunkFile) Close() error {
	log.Infof("DataChunkFile.Close() - start")
	defer log.Infof("DataChunkFile.Close() - end")
	defer f.close()

	if f.currentOffset == 0 {
		// No data were sent via this stream, no need to send finalizer
		return nil
	}

	// Some data were sent via this stream, need to finalize transmission with finalizer packet

	// Send "last" data chunk
	chunk := NewDataChunk(nil, nil, nil, true, nil)
	err := f.transport.Send(chunk)
	if err != nil {
		if err == io.EOF {
			log.Infof("Send() received EOF")
		} else {
			log.Errorf("failed to Send() %v", err)
		}
	}

	return err
}

// TODO implement Reset function for DataChunkFile,
//  so the same descriptor can be used for multiple transmissions.
