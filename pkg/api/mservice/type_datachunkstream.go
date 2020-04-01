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

package mservice

import (
	log "github.com/golang/glog"
	"io"
)

// DataChunkStream is a handler to open stream of DataChunk's
// Inspired by os.File handler and is expected to be used in the same context
// OpenDataStream()
// Write()
// Close()
// it implements the following interfaces:
//	- io.Writer
//	- io.Closer
//	- io.WriterTo
// and thus can be used in any functions, which operate these interfaces, such as io.Copy()
type DataChunkStream struct {
	client         MServiceControlPlane_DataClient
	_type          uint32
	name           string
	metadata       *Metadata
	version        uint32
	uuid_reference string
	description    string

	offset uint64
}

// OpenDataChunkStream opens DataChunk's stream with specified parameters
// Inspired by os.OpenFile()
func OpenDataChunkStream(
	client MServiceControlPlane_DataClient,
	_type uint32,
	name string,
	metadata *Metadata,
	version uint32,
	uuid_reference string,
	description string,
) (*DataChunkStream, error) {
	return &DataChunkStream{
		client:         client,
		_type:          _type,
		name:           name,
		metadata:       metadata,
		version:        version,
		uuid_reference: uuid_reference,
		description:    description,
	}, nil
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
func (s *DataChunkStream) Write(p []byte) (n int, err error) {
	n = len(p)
	log.Infof("before Send()")
	var md *Metadata
	if s.offset == 0 {
		// First chunk in the stream, it may have some metadata
		md = s.metadata
	}
	dataChunk := NewDataChunk(md, &s.offset, false, p)
	err = s.client.Send(dataChunk)
	if err == io.EOF {
		log.Infof("Send() received EOF, return from func")
		n = 0
	}
	if err != nil {
		log.Fatalf("failed to Send() %v", err)
		n = 0
	}

	s.offset += uint64(n)

	log.Infof("after Send()")
	return
}

// Implements io.Closer
//
// Closer is the interface that wraps the basic Close method.
//
// The behavior of Close after the first call is undefined.
// Specific implementations may document their own behavior.
func (s *DataChunkStream) Close() error {
	dataChunk := NewDataChunk(nil, nil, true, nil)
	err := s.client.Send(dataChunk)
	if err == io.EOF {
		log.Infof("Send() received EOF, return from func")
	}
	if err != nil {
		log.Fatalf("failed to Send() %v", err)
	}
	log.Infof("after Send()")
	return err
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
func (s *DataChunkStream) ReadFrom(dataSource io.Reader) (n int64, err error) {
	n = 0
	p := make([]byte, 1024)
	for {
		read, readErr := dataSource.Read(p)
		n += int64(read)
		if read > 0 {
			_, writeErr := s.Write(p[:read])
			if writeErr != nil {
				err = writeErr
				return
			}
		}
		if readErr == io.EOF {
			readErr = nil
			return
		}
		if readErr != nil {
			return
		}
	}
}
