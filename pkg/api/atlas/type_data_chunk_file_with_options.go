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
)

// DataChunkFileWithOptions
// Inspired by os.File handler and is expected to be used in the same context.
type DataChunkFileWithOptions struct {
	DataChunkFile *DataChunkFile
	Compression   *DataChunkFileCompression
}

// OpenDataChunkFileWOptions
func OpenDataChunkFileWOptions(
	transport DataChunkTransporter,
	options *DataChunkFileOptions,
) (*DataChunkFileWithOptions, error) {
	log.Tracef("OpenDataChunkFileWOptions() - start")
	defer log.Tracef("OpenDataChunkFileWOptions() - end")

	// Open underlying DataChunkFile
	f, err := OpenDataChunkFile(transport)
	if err != nil {
		log.Warnf("FAILED to open DataChunkFile. err: %v", err)
		return nil, err
	}

	// Build DataChunkFile with options
	fWOpts := &DataChunkFileWithOptions{}
	fWOpts.DataChunkFile = f
	fWOpts.DataChunkFile.Header = options.GetHeader()
	fWOpts.DataChunkFile.PayloadMetadata = options.GetMetadata()

	readCompression := CompressionNone
	if options.GetDecompress() {
		log.Infof("requesting LZMA decompression")
		readCompression = CompressionLZMA
	}

	writeCompression := CompressionNone
	if options.GetCompress() {
		log.Infof("requesting LZMA compression")
		writeCompression = CompressionLZMA
		// Set compression in transport metadata
		fWOpts.DataChunkFile.ensureTransportMetadata()
		fWOpts.DataChunkFile.TransportMetadata.SetCompression(CompressionLZMA.String())
	}
	fWOpts.Compression, _ = NewDataChunkFileCompression(readCompression, fWOpts.DataChunkFile, writeCompression, fWOpts.DataChunkFile)

	return fWOpts, nil
}

// Close
func (f *DataChunkFileWithOptions) Close() error {
	log.Tracef("DataChunkFileWithOptions.Close() - start")
	defer log.Tracef("DataChunkFileWithOptions.Close() - end")

	if f == nil {
		return nil
	}

	err1 := f.Compression.Close()
	err2 := f.DataChunkFile.Close()

	switch {
	case err1 != nil:
		return err1
	case err2 != nil:
		return err2
	default:
		return nil
	}
}

// Write
func (f *DataChunkFileWithOptions) Write(p []byte) (n int, err error) {
	log.Tracef("DataChunkFileWithOptions.Write() - start")
	defer log.Tracef("DataChunkFileWithOptions.Write() - end")

	if f.Compression.WriteEnabled() {
		return f.Compression.Write(p)
	}
	if f.DataChunkFile != nil {
		return f.DataChunkFile.Write(p)
	}

	return 0, fmt.Errorf("unknown write() entity")
}

// WriteTo writes data to dst
func (f *DataChunkFileWithOptions) WriteTo(dst io.Writer) (int64, error) {
	log.Tracef("DataChunkFileWithOptions.WriteTo() - start")
	defer log.Tracef("DataChunkFileWithOptions.WriteTo() - end")

	return cp(dst, f)
}

// Read
func (f *DataChunkFileWithOptions) Read(p []byte) (n int, err error) {
	log.Tracef("DataChunkFileWithOptions.Read() - start")
	defer log.Tracef("DataChunkFileWithOptions.Read() - end")

	if f.Compression.ReadEnabled() {
		// TODO need to read uncompressed data even if compression requested
		log.Debugf("decompression requested")

		if !f.DataChunkFile.HasTransportMetadata() {
			log.Debugf("no TransportMetadata yet, wait for it")
			f.DataChunkFile.recvDataChunkIntoBuf()
		}

		if !f.DataChunkFile.HasTransportMetadata() {
			log.Warnf("got no TransportMetadata, abort")
			return 0, fmt.Errorf("decompression requested, but no metadata available")
		}

		if f.DataChunkFile.TransportMetadata.GetCompression() != "" {
			log.Tracef("reading compressed data")
			return f.Compression.Read(p)
		}

		log.Warnf("unknown compression method %v", f.DataChunkFile.TransportMetadata.GetCompression())

		return 0, fmt.Errorf("unknown compression method")
	}

	if f.DataChunkFile != nil {
		return f.DataChunkFile.Read(p)
	}

	return 0, fmt.Errorf("unknown read() entity")
}

// ReadFrom reads data from src
func (f *DataChunkFileWithOptions) ReadFrom(src io.Reader) (int64, error) {
	log.Tracef("DataChunkFileWithOptions.ReadFrom() - start")
	defer log.Tracef("DataChunkFileWithOptions.ReadFrom() - end")

	n, err := cp(f, src)

	log.Debugf("Accepted data meta:")
	f.DataChunkFile.PayloadMetadata.Log()
	f.DataChunkFile.TransportMetadata.Log()

	return n, err
}

// WriteToBuf writes data to newly created buffer
func (f *DataChunkFileWithOptions) WriteToBuf() (int64, *bytes.Buffer, error) {
	log.Tracef("DataChunkFileWithOptions.WriteToBuf() - start")
	defer log.Tracef("DataChunkFileWithOptions.WriteToBuf() - end")

	var buf = &bytes.Buffer{}
	written, err := f.WriteTo(buf)
	if err != nil {
		log.Errorf("got error: %v", err.Error())
	}

	// Debug
	log.Debugf("metadata: %s", f.DataChunkFile.PayloadMetadata.String())
	log.Debugf("data: %s", buf.String())

	return written, buf, err
}
