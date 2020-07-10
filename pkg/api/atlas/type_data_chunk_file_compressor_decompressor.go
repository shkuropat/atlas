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

	log "github.com/sirupsen/logrus"
	"github.com/ulikunitz/xz/lzma"
)

const (
	CompressionNone = "none"
	CompressionLZMA = "lzma"
)

// DataChunkFileReadCompression is a compression descriptor
type DataChunkFileReadCompression struct {
	Type       string
	LZMAReader *lzma.Reader
}

// DataChunkFileWriteCompression is a compression descriptor
type DataChunkFileWriteCompression struct {
	Type       string
	LZMAWriter *lzma.Writer
}

// DataChunkFileWriter
type DataChunkFileWriter struct {
	DataChunkFile *DataChunkFile
	Compression   DataChunkFileWriteCompression
}

// OpenDataChunkFileCompressor
func OpenDataChunkFileCompressor(
	transport DataChunkTransport,
	header *Metadata,
	metadata *Metadata,
	compress bool,
) (*DataChunkFileWriter, error) {
	log.Infof("OpenDataChunkFileCompressor() - start")
	defer log.Infof("OpenDataChunkFileCompressor() - end")

	f := &DataChunkFile{
		transport:       transport,
		Header:          header,
		PayloadMetadata: metadata,
	}

	if compress {
		log.Infof("OpenDataChunkFileCompressor() - requesting LZMA compression")

		f.ensureTransportMetadata()
		f.TransportMetadata.SetCompression(CompressionLZMA)
		lzmaWriter, err := lzma.NewWriter(f)
		if err != nil {
			log.Warnf("FAILED to create lzma writer. err: %v", err)
			return nil, err
		}
		return &DataChunkFileWriter{
			DataChunkFile: f,
			Compression: DataChunkFileWriteCompression{
				Type:       CompressionLZMA,
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
	log.Infof("DataChunkFileWriter.Write() - start")
	defer log.Infof("DataChunkFileWriter.Write() - end")

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
	Compression   DataChunkFileReadCompression
}

// OpenDataChunkFileDecompressor
func OpenDataChunkFileDecompressor(transport DataChunkTransport, decompress bool) (*DataChunkFileReader, error) {
	log.Infof("OpenDataChunkFileDecompressor() - start")
	defer log.Infof("OpenDataChunkFileDecompressor() - end")

	f := &DataChunkFile{
		transport: transport,
	}

	if decompress {
		log.Infof("OpenDataChunkFileDecompressor() - requesting LZMA decompression")

		lzmaReader, err := lzma.NewReader(f)
		if err != nil {
			log.Warnf("FAILED to create lzma reader. err: %v", err)
			return nil, err
		}

		return &DataChunkFileReader{
			DataChunkFile: f,
			Compression: DataChunkFileReadCompression{
				Type:       CompressionLZMA,
				LZMAReader: lzmaReader,
			},
		}, nil
	}

	return &DataChunkFileReader{
		DataChunkFile: f,
	}, nil
}

// Close
func (r *DataChunkFileReader) Close() error {
	var err error
	if r.DataChunkFile != nil {
		err = r.DataChunkFile.Close()
	}
	return err
}

// Read
func (r *DataChunkFileReader) Read(p []byte) (n int, err error) {
	log.Infof("DataChunkFileReader.Read() - start")
	defer log.Infof("DataChunkFileReader.Read() - end")

	if r.Compression.LZMAReader != nil {
		log.Infof("DataChunkFileReader.Read() - decompression requested")

		if !r.DataChunkFile.HasTransportMetadata() {
			log.Infof("DataChunkFileReader.Read() - no TransportMetadata yet, wait for it")
			r.DataChunkFile.recvDataChunkAndAppendBuf()
		}

		if !r.DataChunkFile.HasTransportMetadata() {
			log.Warnf("DataChunkFileReader.Read() - got no TransportMetadata, abort")
			return 0, fmt.Errorf("decompression requested, but no metadata available")
		}

		if r.DataChunkFile.TransportMetadata.GetCompression() != "" {
			log.Infof("DataChunkFileReader.Read() - reading compressed data")
			return r.Compression.LZMAReader.Read(p)
		}

		log.Warnf("DataChunkFileReader.Read() - unknown compression method %v", r.DataChunkFile.TransportMetadata.GetCompression())

		return 0, fmt.Errorf("unknown compression method")
	}

	if r.DataChunkFile != nil {
		return r.DataChunkFile.Read(p)
	}

	return 0, fmt.Errorf("unknown read() entity")
}
