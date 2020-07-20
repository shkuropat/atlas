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
	"github.com/ulikunitz/xz/lzma"
)

const (
	CompressionNone = "none"
	CompressionLZMA = "lzma"
)

// DataChunkFileWOptions
// Inspired by os.File handler and is expected to be used in the same context.
type DataChunkFileWOptions struct {
	DataChunkFile *DataChunkFile
	Compression   DataChunkFileCompression
}

// DataChunkFileCompression is a compression descriptor
type DataChunkFileCompression struct {
	Type       string
	LZMAReader *lzma.Reader
	LZMAWriter *lzma.Writer
}

// OpenDataChunkFileWOptions
func OpenDataChunkFileWOptions(transport DataChunkTransport, options *DataChunkFileOptions) (*DataChunkFileWOptions, error) {
	log.Tracef("OpenDataChunkFileWOptions() - start")
	defer log.Tracef("OpenDataChunkFileWOptions() - end")

	// Open underlying DataChunkFile
	f, err := OpenDataChunkFile(transport)
	if err != nil {
		log.Warnf("FAILED to open DataChunkFile. err: %v", err)
		return nil, err
	}

	// Build DataChunkFile with options
	fWOpts := &DataChunkFileWOptions{}
	fWOpts.DataChunkFile = f
	fWOpts.DataChunkFile.Header = options.GetHeader()
	fWOpts.DataChunkFile.PayloadMetadata = options.GetMetadata()

	if options.GetDecompress() {
		log.Infof("requesting LZMA decompression")

		lzmaReader, err := lzma.NewReader(fWOpts.DataChunkFile)
		if err != nil {
			log.Warnf("FAILED to create lzma reader. err: %v", err)
			return nil, err
		}
		fWOpts.Compression.Type = CompressionLZMA
		fWOpts.Compression.LZMAReader = lzmaReader
	}

	if options.GetCompress() {
		log.Infof("requesting LZMA compression")

		fWOpts.DataChunkFile.ensureTransportMetadata()
		fWOpts.DataChunkFile.TransportMetadata.SetCompression(CompressionLZMA)
		lzmaWriter, err := lzma.NewWriter(fWOpts.DataChunkFile)
		if err != nil {
			log.Warnf("FAILED to create lzma writer. err: %v", err)
			return nil, err
		}
		fWOpts.Compression.Type = CompressionLZMA
		fWOpts.Compression.LZMAWriter = lzmaWriter
	}

	return fWOpts, nil
}

// Close
func (f *DataChunkFileWOptions) Close() error {
	log.Tracef("DataChunkFileWOptions.Close() - start")
	defer log.Tracef("DataChunkFileWOptions.Close() - end")

	var err1 error
	var err2 error
	if f.Compression.LZMAWriter != nil {
		err1 = f.Compression.LZMAWriter.Close()
	}
	if f.DataChunkFile != nil {
		err2 = f.DataChunkFile.Close()
	}

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
func (f *DataChunkFileWOptions) Write(p []byte) (n int, err error) {
	log.Tracef("DataChunkFileWOptions.Write() - start")
	defer log.Tracef("DataChunkFileWOptions.Write() - end")

	if f.Compression.LZMAWriter != nil {
		return f.Compression.LZMAWriter.Write(p)
	}
	if f.DataChunkFile != nil {
		return f.DataChunkFile.Write(p)
	}

	return 0, fmt.Errorf("unknown write() entity")
}

// WriteTo writes data to dst
func (f *DataChunkFileWOptions) WriteTo(dst io.Writer) (int64, error) {
	log.Tracef("DataChunkFileWOptions.WriteTo() - start")
	defer log.Tracef("DataChunkFileWOptions.WriteTo() - end")

	return cp(dst, f)
}

// Read
func (f *DataChunkFileWOptions) Read(p []byte) (n int, err error) {
	log.Tracef("DataChunkFileWOptions.Read() - start")
	defer log.Tracef("DataChunkFileWOptions.Read() - end")

	if f.Compression.LZMAReader != nil {
		log.Debugf("decompression requested")

		if !f.DataChunkFile.HasTransportMetadata() {
			log.Debugf("no TransportMetadata yet, wait for it")
			f.DataChunkFile.recvDataChunkAndAppendBuf()
		}

		if !f.DataChunkFile.HasTransportMetadata() {
			log.Warnf("got no TransportMetadata, abort")
			return 0, fmt.Errorf("decompression requested, but no metadata available")
		}

		if f.DataChunkFile.TransportMetadata.GetCompression() != "" {
			log.Tracef("reading compressed data")
			return f.Compression.LZMAReader.Read(p)
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
func (f *DataChunkFileWOptions) ReadFrom(src io.Reader) (int64, error) {
	log.Tracef("DataChunkFileWOptions.ReadFrom() - start")
	defer log.Tracef("DataChunkFileWOptions.ReadFrom() - end")

	n, err := cp(f, src)

	log.Debugf("Accepted data meta:")
	f.DataChunkFile.PayloadMetadata.Log()
	f.DataChunkFile.TransportMetadata.Log()

	return n, err
}

// WriteToBuf writes data to newly created buffer
func (f *DataChunkFileWOptions) WriteToBuf() (int64, *bytes.Buffer, error) {
	log.Tracef("DataChunkFileWOptions.WriteToBuf() - start")
	defer log.Tracef("DataChunkFileWOptions.WriteToBuf() - end")

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

// cp copies from src into dst. Have to use `cp` because `copy` is a built-in function
func cp(dst io.Writer, src io.Reader) (int64, error) {
	log.Tracef("cp() - start")
	defer log.Tracef("cp() - end")

	var written int64
	var err error
	size := 32 * 1024
	buf := make([]byte, size)

	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
