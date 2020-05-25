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
	"bytes"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/binarly-atlas/pkg/minio"
)

// DataChunkSenderReceiver defines transport level interface (for both client and server),
// which serves DataChunk streams bi-directionally.
type DataChunkSenderReceiver interface {
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
	DataChunkSenderReceiver DataChunkSenderReceiver

	// Header is mandatory
	Header *Header
	// Metadata is optional
	Metadata *Metadata

	// initialOffset of these chunks
	initialOffset int64
	// currentOffset of the current chunk within file
	currentOffset int64
	// offset of the current chunk within "all chunks" = initialOffset + currentOffset
	offset int64
}

// OpenDataChunkFile opens set of DataChunk(s)
// Inspired by os.OpenFile()
func OpenDataChunkFile(DataChunkSenderReceiver DataChunkSenderReceiver) (*DataChunkFile, error) {
	return &DataChunkFile{
		DataChunkSenderReceiver: DataChunkSenderReceiver,
	}, nil
}

// ensureMetadata ensures DataChunkFile has Metadata in place
func (s *DataChunkFile) ensureMetadata() {
	if s.Metadata == nil {
		s.Metadata = new(Metadata)
	}
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
func (s *DataChunkFile) Write(p []byte) (n int, err error) {
	log.Infof("Write() - start")
	defer log.Infof("Write() - end")

	n = len(p)
	var md *Metadata = nil
	if s.currentOffset == 0 {
		// First chunk in this file, it may have some metadata
		md = s.Metadata
	}
	s.offset = s.initialOffset + s.currentOffset
	chunk := NewDataChunk(md, &s.offset, false, p)
	err = s.DataChunkSenderReceiver.Send(chunk)
	if err != nil {
		// We have some kind of error and were not able to send the data
		n = 0
		if err == io.EOF {
			// Not sure, probably this is not an error, after all
			log.Infof("Send() received EOF")
		} else {
			log.Fatalf("failed to Send() %v", err)
		}
	}

	s.currentOffset += int64(n)

	return
}

// Implements io.WriterTo
//
// WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to w until there's no more data to write or
// when an error occurs. The return value n is the number of bytes
// written. Any error encountered during the write is also returned.
//
// The Copy function uses WriterTo if available.
func (s *DataChunkFile) WriteTo(dst io.Writer) (n int64, err error) {
	log.Infof("WriteTo() - start")
	defer log.Infof("WriteTo() - end")

	n = 0
	for {
		// Whether this chunk is the last one on the stream
		lastChunk := false
		dataChunk, readErr := s.DataChunkSenderReceiver.Recv()

		if dataChunk != nil {
			// We've got data chunk and it has to be processed no matter what.
			// Even in case Recv() reported error, we still need to process obtained data

			// Fetch filename from the chunks stream - it may be in any chunk, actually
			filename := "not specified"
			if md := dataChunk.GetMetadata(); md != nil {
				filename = md.GetFilename()
				if filename != "" {
					s.ensureMetadata()
					s.Metadata.SetFilename(filename)
				}
			}

			// Fetch offset of this chunk within the stream
			offset := "not specified"
			if off, ok := dataChunk.GetOffsetWithAvailabilityReport(); ok {
				offset = fmt.Sprintf("%d", off)
			}

			// How many bytes do we have in this chunk?
			_len := len(dataChunk.GetBytes())
			log.Infof("Data.Recv() got msg. filename: '%s', chunk len: %d, chunk offset: %s, last chunk: %v",
				filename,
				_len,
				offset,
				dataChunk.GetLast(),
			)
			fmt.Printf("%s\n", string(dataChunk.GetBytes()))

			// TODO need to handle write errors
			_, _ = dst.Write(dataChunk.GetBytes())

			n += int64(_len)

			// This is officially last chunk on the chunks stream, so no need to Recv() any more, break the recv loop,
			// but we need to handle possible errors first
			if dataChunk.GetLast() {
				lastChunk = true
			}
		}

		if readErr == nil {
			// All went well, ready to receive more data
		} else if readErr == io.EOF {
			// Correct EOF
			log.Infof("Data.Recv() get EOF")
			err = readErr
			return
		} else {
			// Stream broken
			log.Infof("Data.Recv() got err: %v", err)
			err = readErr
			return
		}

		if lastChunk {
			// This is officially last chunk on the chunks stream, so no need to Recv() any more, break the recv loop
			break
		}
	}

	return
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
func (s *DataChunkFile) Read(p []byte) (n int, err error) {
	log.Infof("Read() - start")
	defer log.Infof("Read() - end")

	return 0, fmt.Errorf("not implemented function: Read()")
}

// Implements io.ReaderFrom
//
// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from r until EOF or error.
// The return value n is the number of bytes read.
// Any error except io.EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
func (s *DataChunkFile) ReadFrom(src io.Reader) (n int64, err error) {
	log.Infof("ReadFrom() - start")
	defer log.Infof("ReadFrom() - end")

	n = 0
	p := make([]byte, 1024)
	for {
		read, readErr := src.Read(p)
		n += int64(read)
		if read > 0 {
			// We've got some data and it has to be processed no matter what
			// Write these bytes to data stream
			_, writeErr := s.Write(p[:read])
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
func (s *DataChunkFile) Close() error {
	log.Infof("Close() - start")
	defer log.Infof("Close() - end")

	if s.currentOffset == 0 {
		// No data were sent via this stream, no need to send finalizer
		return nil
	}

	// Some data were sent via this stream, need to finalize transmission with finalizer packet

	// Send "last" data chunk
	chunk := NewDataChunk(nil, nil, true, nil)
	err := s.DataChunkSenderReceiver.Send(chunk)
	if err != nil {
		if err == io.EOF {
			log.Infof("Send() received EOF")
		} else {
			log.Fatalf("failed to Send() %v", err)
		}
	}

	return err
}

// TODO implement Reset function for DataChunkFile,
//  so the same descriptor can be used for multiple transmissions.

func SendDataChunkFile(
	DataChunkSenderReceiver DataChunkSenderReceiver,
	metadata *Metadata,
	src io.Reader,
) (int64, error) {
	log.Infof("SendDataChunkFile() - start")
	defer log.Infof("SendDataChunkFile() - end")

	f, err := OpenDataChunkFile(DataChunkSenderReceiver)
	if err != nil {
		log.Warnf("got err: %v", err)
		return 0, err
	}
	defer f.Close()
	f.Header = NewHeader(int32(DataChunkType_DATA_CHUNK_TYPE_DATA), "", 0, "", "123", 0, 0, "desc")
	f.Metadata = metadata
	return io.Copy(f, src)
}

func RecvDataChunkFile(DataChunkSenderReceiver DataChunkSenderReceiver, dst io.Writer) (int64, *Metadata, error) {
	log.Infof("RecvDataChunkFile() - start")
	defer log.Infof("RecvDataChunkFile() - end")

	f, err := OpenDataChunkFile(DataChunkSenderReceiver)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()

	written, err := io.Copy(dst, f)
	if err != nil {
		log.Errorf("RecvDataChunkFile() got error: %v", err.Error())
	}

	return written, f.Metadata, err
}

func RecvDataChunkFileIntoBuf(DataChunkSenderReceiver DataChunkSenderReceiver) (int64, *bytes.Buffer, *Metadata, error) {
	log.Infof("RecvDataChunkFileIntoBuf() - start")
	defer log.Infof("RecvDataChunkFileIntoBuf() - end")

	var buf = &bytes.Buffer{}
	written, metadata, err := RecvDataChunkFile(DataChunkSenderReceiver, buf)
	if err != nil {
		log.Errorf("RecvDataChunkFileIntoBuf() got error: %v", err.Error())
	}

	// Debug
	log.Infof("metadata: %s", metadata.String())
	log.Infof("data: %s", buf.String())

	return written, buf, metadata, err
}

func RelayDataChunkFileIntoMinIO(DataChunkSenderReceiver DataChunkSenderReceiver, mi *minio.MinIO, bucketName, objectName string) (int64, *Metadata, error) {
	log.Infof("RelayDataChunkFileIntoMinIO() - start")
	defer log.Infof("RelayDataChunkFileIntoMinIO() - end")

	f, err := OpenDataChunkFile(DataChunkSenderReceiver)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()

	written, err := mi.Put(bucketName, objectName, f)
	if err != nil {
		log.Errorf("RelayDataChunkFileIntoMinIO() got error: %v", err.Error())
	}

	// Debug
	log.Infof("metadata: %s", f.Metadata.String())

	return written, f.Metadata, err
}
