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

const (
	LZMACompression = "lzma"
)

// DataChunkFileCompression is a compression descriptor
type DataChunkFileCompression struct {
	Type       string
	LZMAReader *lzma.Reader
	LZMAWriter *lzma.Writer
}

// DataChunkFileWriter
type DataChunkFileWriter struct {
	DataChunkFile *DataChunkFile
	Compression   DataChunkFileCompression
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

	if compress {
		f.ensureTransportMetadata()
		f.TransportMetadata.SetCompression(LZMACompression)
		lzmaWriter, err := lzma.NewWriter(f)
		if err != nil {
			return nil, err
		}
		return &DataChunkFileWriter{
			DataChunkFile: f,
			Compression: DataChunkFileCompression{
				Type:       LZMACompression,
				LZMAWriter: lzmaWriter,
			},
		}, nil
	}

	return &DataChunkFileWriter{
		DataChunkFile: f,
	}, nil
}

// Close
func (w *DataChunkFileWriter) Close() error {
	var err error
	if w.Compression.LZMAWriter != nil {
		err = w.Compression.LZMAWriter.Close()
	}
	if w.DataChunkFile != nil {
		err = w.DataChunkFile.Close()
	}
	return err
}

// Write
func (w *DataChunkFileWriter) Write(p []byte) (n int, err error) {
	if w.Compression.LZMAWriter != nil {
		return w.Compression.LZMAWriter.Write(p)
	}
	if w.DataChunkFile != nil {
		return w.DataChunkFile.Write(p)
	}

	return 0, fmt.Errorf("unknown write() entity")
}

// DataChunkFileReader
type DataChunkFileReader struct {
	DataChunkFile *DataChunkFile
	Compression   DataChunkFileCompression
}

// OpenDataChunkFileReader
func OpenDataChunkFileReader(transport DataChunkTransport, decompress bool) (*DataChunkFileReader, error) {
	dcf := &DataChunkFile{
		transport: transport,
	}

	if decompress {
		lzmaReader, err := lzma.NewReader(dcf)
		if err != nil {
			return nil, err
		}
		return &DataChunkFileReader{
			DataChunkFile: dcf,
			Compression: DataChunkFileCompression{
				Type:       LZMACompression,
				LZMAReader: lzmaReader,
			},
		}, nil
	}

	return &DataChunkFileReader{
		DataChunkFile: dcf,
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
	if w.Compression.LZMAReader != nil {
		if !w.DataChunkFile.HasTransportMetadata() {
			w.DataChunkFile.appendDataBuf()
		}

		if w.DataChunkFile.HasTransportMetadata() {
			if w.DataChunkFile.TransportMetadata.GetCompression() != "" {
				return w.Compression.LZMAReader.Read(p)
			}
		}
	}

	if w.DataChunkFile != nil {
		return w.DataChunkFile.Read(p)
	}

	return 0, fmt.Errorf("unknown read() entity")
}
