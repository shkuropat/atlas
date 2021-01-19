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
	"github.com/binarly-io/atlas/pkg/util"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
)

const defaultMaxWriteChunkSize = 1024

// DataChunkFile is a handler to set of DataChunk(s)
// Inspired by os.File handler and is expected to be used in the same context.
// DataChunkFile implements the following interfaces:
//	- io.Writer
//	- io.WriterTo
// 	- io.ReaderFrom
//	- io.Closer
// and thus can be used in any functions, which operate these interfaces, such as io.Copy()
type DataChunkFile struct {
	// transport provides functions to send/receive DataChunks
	transport DataChunkTransporter

	// Header is mandatory
	Header *Metadata
	// TransportMetadata is optional
	TransportMetadata *Metadata
	// PayloadMetadata is optional
	PayloadMetadata *Metadata

	// initialOffset of these chunks
	initialOffset int64
	// currentOffset of the current chunk within current chunk set (file)
	currentOffset int64
	// offset of the current chunk within "all chunks" = initialOffset + currentOffset
	offset int64

	// MaxWriteChunkSize limits max size of a payload within one data chunk to be sent
	MaxWriteChunkSize int

	// Receive part
	recvBuf []byte
	recvErr error

	// Log part
	// lastTimeDataChunkLogged specifies when last time data chunk was logged
	lastTimeDataChunkLogged time.Time
}

// OpenDataChunkFile opens set of DataChunk(s)
// Inspired by os.OpenFile()
func OpenDataChunkFile(transport DataChunkTransporter) (*DataChunkFile, error) {
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

// ensureRecvBuf ensures DataChunkFile has buf in place
func (f *DataChunkFile) ensureRecvBuf() {
	if f.recvBuf == nil {
		f.recvBuf = make([]byte, 0)
	}
}

// close does internal job to complete DataChunkFile
func (f *DataChunkFile) close() {
	if f == nil {
		return
	}

	f.Header = nil
	f.TransportMetadata = nil
	f.PayloadMetadata = nil

	f.initialOffset = 0
	f.currentOffset = 0
	f.offset = 0

	f.recvBuf = nil
	f.recvErr = nil

	// TODO flush outgoing transport
	// TODO should incoming transport be drained?
}

// acceptHeader accepts header from DataChunk into file
func (f *DataChunkFile) acceptHeader(dataChunk *DataChunk) {
	if header := dataChunk.GetHeader(); header != nil {
		f.Header = header
	}
}

// acceptTransportMetadata accepts transport metadata from DataChunk into file
func (f *DataChunkFile) acceptTransportMetadata(dataChunk *DataChunk) {
	if md := dataChunk.GetTransportMetadata(); md != nil {
		f.TransportMetadata = md
	}
}

// acceptPayloadMetadata accepts payload metadata from DataChunk into file
func (f *DataChunkFile) acceptPayloadMetadata(dataChunk *DataChunk) {
	if md := dataChunk.GetPayloadMetadata(); md != nil {
		f.PayloadMetadata = md
	}
}

// acceptAllMetadata accepts all metadata from from DataChunk into file
func (f *DataChunkFile) acceptAllMetadata(dataChunk *DataChunk) {
	if dataChunk == nil {
		return
	}
	f.acceptTransportMetadata(dataChunk)
	f.acceptPayloadMetadata(dataChunk)
	f.acceptHeader(dataChunk)
}

// logDataChunk logs DataChunk
func (f *DataChunkFile) logDataChunk(dataChunk *DataChunk) {
	if dataChunk == nil {
		return
	}

	now := time.Now()
	interval := 30 * time.Second
	switch {
	case
		f.lastTimeDataChunkLogged.IsZero(),                 // No chunks logged before, log the first one
		dataChunk.Header.GetLast(),                         // Log the last one to complete transfer logging
		now.After(f.lastTimeDataChunkLogged.Add(interval)), // Log every X seconds
		log.GetLevel() == log.TraceLevel:                   // In case of trace all chunks should be logged
		// Should log this DataChunk
	default:
		// Should not log this DataChunk
		return
	}

	// Fetch filename from the chunks stream - it may be in any chunk, actually
	filename := "not specified"
	if f.HasPayloadMetadata() {
		if f.PayloadMetadata.HasFilename() {
			filename = f.PayloadMetadata.GetFilename()
		}
	}

	compression := CompressionNone
	if f.HasTransportMetadata() {
		if f.TransportMetadata.HasCompression() {
			compression = ParseCompressionType(f.TransportMetadata.GetCompression())
		}
	}

	// Fetch offset of this chunk within the stream
	offset := "not specified"
	if dataChunk.Header.HasOffset() {
		offset = fmt.Sprintf("%d", dataChunk.Header.GetOffset())
	}

	// How many bytes do we have in this chunk?
	_len := len(dataChunk.GetBytes())
	log.Infof("got DataChunk. filename:%s, compression:%s, chunk len:%d, chunk offset:%s, last chunk:%v",
		filename,
		compression,
		_len,
		offset,
		dataChunk.Header.GetLast(),
	)
	// Dump content
	//if compression == CompressionNone {
	//	fmt.Printf("%s\n", string(dataChunk.GetBytes()))
	//} else {
	//	fmt.Printf("got %d %s compressed bytes\n", _len, compression)
	//}

	f.lastTimeDataChunkLogged = now
}

// recvDataChunk
func (f *DataChunkFile) recvDataChunk() (*DataChunk, error) {
	log.Tracef("DataChunkFile.recvDataChunk() - start")
	defer log.Tracef("DataChunkFile.recvDataChunk() - end")

	// Recv DataChunk
	dataChunk, err := f.transport.Recv()
	f.acceptAllMetadata(dataChunk)
	f.logDataChunk(dataChunk)

	switch err {
	case nil:
		// All went well, ready to receive more data
	case io.EOF:
		// Correct EOF arrived
		log.Infof("DataChunkFile.Recv() get EOF")
	default:
		// Stream is somehow broken
		log.Infof("DataChunkFile.Recv() got err: %v", err)
	}

	return dataChunk, err
}

// recvDataChunkIntoBuf
func (f *DataChunkFile) recvDataChunkIntoBuf() {
	dataChunk, err := f.recvDataChunk()
	if dataChunk != nil {
		if len(dataChunk.GetBytes()) > 0 {
			// TODO
			f.recvBuf = append(f.recvBuf, dataChunk.GetBytes()...)
		}

		if dataChunk.Header.GetLast() {
			f.recvErr = io.EOF
		}
	}

	if err != nil {
		f.recvErr = err
	}
}

// newDataChunk
func (f *DataChunkFile) newDataChunk(data []byte) *DataChunk {
	var transportMD *Metadata
	var payloadMD *Metadata
	if f.currentOffset == 0 {
		// First chunk in this file, it may have some metadata
		transportMD = f.TransportMetadata
		payloadMD = f.PayloadMetadata
		log.Tracef("Attaching metadata. transport=%v payload=%v", transportMD, payloadMD)
	}
	// Offset of current chunk
	f.offset = f.initialOffset + f.currentOffset
	return NewDataChunk(transportMD, payloadMD, &f.offset, false, data)
}

// sendDataChunk
func (f *DataChunkFile) sendDataChunk(p []byte) (n int, err error) {
	log.Tracef("DataChunkFile.sendDataChunk() - start")
	defer log.Tracef("DataChunkFile.sendDataChunk() - end")

	if len(p) == 0 {
		return 0, nil
	}

	if (len(p) > f.MaxWriteChunkSize) && (f.MaxWriteChunkSize > 0) {
		return 0, fmt.Errorf("attempt to sendDataChunk() with oversized chunk: %d > %d", len(p), f.MaxWriteChunkSize)
	}

	n = len(p)
	chunk := f.newDataChunk(p)
	err = f.transport.Send(chunk)
	if err != nil {
		// We have some kind of error and were not able to send the data
		n = 0
		if err == io.EOF {
			// Not sure, probably this is not an error, after all
			log.Infof("sendDataChunk.Send() received EOF")
		} else {
			log.Errorf("sendDataChunk.Send() FAILED %v", err)
		}
	}

	// Adjust offset of next chunk to be sent after this one
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
	log.Tracef("DataChunkFile.Write() - start")
	defer log.Tracef("DataChunkFile.Write() - end")

	if (len(p) < f.MaxWriteChunkSize) || (f.MaxWriteChunkSize <= 0) {
		// No need to split into chunks, can send all data as one piece
		return f.sendDataChunk(p)
	}

	// This piece is too big, need to split into several chunks
	n = 0
	for offset := 0; offset < len(p); offset += f.MaxWriteChunkSize {
		if sent, e := f.sendDataChunk(p[offset:util.Min(offset+f.MaxWriteChunkSize, len(p))]); e != nil {
			return n, e
		} else {
			n += sent
		}
	}

	return n, nil
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
	log.Tracef("DataChunkFile.WriteTo() - start")
	defer log.Tracef("DataChunkFile.WriteTo() - end")

	return cp(dst, f, util.IfPositiveValue(f.MaxWriteChunkSize, defaultMaxWriteChunkSize))
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
	log.Tracef("DataChunkFile.Read() - start")
	defer log.Tracef("DataChunkFile.Read() - end")

	if len(f.recvBuf) == 0 {
		// No buffered data available, need to get some
		f.recvDataChunkIntoBuf()
	}

	n = 0
	if len(f.recvBuf) > 0 {
		// Have some buffered data, copy it out
		n = copy(p, f.recvBuf)
		f.recvBuf = f.recvBuf[n:]
	}

	if len(f.recvBuf) > 0 {
		// Still have some data in recvBuf for next read, so not even EOF should be reported right now
		return n, nil
	}

	// No more data in recvBuf left
	return n, f.recvErr
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
	log.Tracef("DataChunkFile.ReadFrom() - start")
	defer log.Tracef("DataChunkFile.ReadFrom() - end")

	return cp(f, src, util.IfPositiveValue(f.MaxWriteChunkSize, defaultMaxWriteChunkSize))
}

// Implements io.Closer
//
// Closer is the interface that wraps the basic Close method.
//
// The behavior of Close after the first call is undefined.
// Specific implementations may document their own behavior.
func (f *DataChunkFile) Close() error {
	log.Tracef("DataChunkFile.Close() - start")
	defer log.Tracef("DataChunkFile.Close() - end")
	if f == nil {
		return nil
	}

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
