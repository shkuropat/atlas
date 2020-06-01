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
)

// DataChunkFileWriter
type DataChunkFileWriter struct {
	DataChunkFile *DataChunkFile
	Writer        *lzma.Writer
}

// OpenDataChunkFileWriter
func OpenDataChunkFileWriter(
	transport DataChunkTransport,
	header *Metadata,
	metadata *Metadata,
	compress bool,
) (*DataChunkFileWriter, error) {

	f := &DataChunkFile{
		transport:       transport,
		Header:          header,
		PayloadMetadata: metadata,
	}

	var lzmaWriter *lzma.Writer
	var err error
	if compress {
		f.ensureTransportMetadata()
		f.TransportMetadata.SetCompression("lzma")
		lzmaWriter, err = lzma.NewWriter(f)
		if err != nil {
			return nil, err
		}
	}

	return &DataChunkFileWriter{
		DataChunkFile: f,
		Writer:        lzmaWriter,
	}, nil
}

// Close
func (w *DataChunkFileWriter) Close() error {
	var err error
	if w.Writer != nil {
		err = w.Writer.Close()
	}
	if w.DataChunkFile != nil {
		err = w.DataChunkFile.Close()
	}
	return err
}

// Write
func (w *DataChunkFileWriter) Write(p []byte) (n int, err error) {
	if w.Writer != nil {
		return w.Writer.Write(p)
	}
	if w.DataChunkFile != nil {
		return w.DataChunkFile.Write(p)
	}

	return 0, fmt.Errorf("unknown write() entity")
}

// DataChunkFileReader
type DataChunkFileReader struct {
	DataChunkFile *DataChunkFile
	Reader        *lzma.Reader
}

// OpenDataChunkFileReader
func OpenDataChunkFileReader(transport DataChunkTransport, decompress bool) (*DataChunkFileReader, error) {
	dcf := &DataChunkFile{
		transport: transport,
	}

	var lzmaReader *lzma.Reader
	var err error
	if decompress {
		lzmaReader, err = lzma.NewReader(dcf)
		if err != nil {
			return nil, err
		}
	}

	return &DataChunkFileReader{
		DataChunkFile: dcf,
		Reader:        lzmaReader,
	}, nil
}

// Close
func (w *DataChunkFileReader) Close() error {
	var err error
	if w.DataChunkFile != nil {
		err = w.DataChunkFile.Close()
	}
	return err
}

// Read
func (w *DataChunkFileReader) Read(p []byte) (n int, err error) {
	if w.Reader != nil {
		if !w.DataChunkFile.HasTransportMetadata() {
			w.DataChunkFile.appendDataBuf()
		}

		if w.DataChunkFile.HasTransportMetadata() {
			if w.DataChunkFile.TransportMetadata.GetCompression() != "" {
				return w.Reader.Read(p)
			}
		}
	}
	if w.DataChunkFile != nil {
		return w.DataChunkFile.Read(p)
	}

	return 0, fmt.Errorf("unknown read() entity")
}
