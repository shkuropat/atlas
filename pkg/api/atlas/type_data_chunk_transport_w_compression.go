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
	"io"

	log "github.com/sirupsen/logrus"
)

// DataChunkTransportCompressionOptions
type DataChunkTransportCompressionOptions struct {
	Compress   bool
	Decompress bool
}

// GetCompress
func (opts *DataChunkTransportCompressionOptions) GetCompress() bool {
	if opts == nil {
		return false
	}

	return opts.Compress
}

// GetDecompress
func (opts *DataChunkTransportCompressionOptions) GetDecompress() bool {
	if opts == nil {
		return false
	}

	return opts.Decompress
}

// DataChunkTransportWithCompression
type DataChunkTransportWithCompression struct {
	Transport DataChunkTransport
	Options   *DataChunkTransportCompressionOptions
}

// OpenDataChunkTransportWithCompression
func OpenDataChunkTransportWithCompression(transport DataChunkTransport, options *DataChunkTransportCompressionOptions) *DataChunkTransportWithCompression {
	return &DataChunkTransportWithCompression{
		Transport: transport,
		Options:   options,
	}
}

// Send
func (f *DataChunkTransportWithCompression) Send(src io.Reader, metadata *Metadata) (int64, error) {
	log.Infof("DataChunkTransportWithCompression.Send() - start")
	defer log.Infof("DataChunkTransportWithCompression.Send() - end")

	w, err := OpenDataChunkFileWriter(
		f.Transport,
		NewMetadata(
			int32(DataChunkType_DATA_CHUNK_TYPE_DATA),
			"",
			0,
			"",
			"123",
			0,
			0,
			"desc",
		),
		metadata,
		f.Options.GetCompress(),
	)
	if err != nil {
		log.Warnf("got err: %v", err)
		return 0, err
	}
	defer w.Close()

	return io.Copy(w, src)
}

// Recv
func (f *DataChunkTransportWithCompression) Recv(dst io.Writer) (int64, *Metadata, error) {
	log.Infof("DataChunkTransportWithCompression.Recv() - start")
	defer log.Infof("DataChunkTransportWithCompression.Recv() - end")

	r, err := OpenDataChunkFileReader(f.Transport, f.Options.GetDecompress())
	if err != nil {
		return 0, nil, err
	}
	defer r.Close()

	written, err := io.Copy(dst, r)
	if err != nil {
		log.Errorf("got error: %v", err.Error())
	}

	return written, r.DataChunkFile.PayloadMetadata, err
}

// RecvIntoBuf
func (f *DataChunkTransportWithCompression) RecvIntoBuf() (int64, *bytes.Buffer, *Metadata, error) {
	log.Infof("DataChunkTransportWithCompression.RecvIntoBuf() - start")
	defer log.Infof("DataChunkTransportWithCompression.RecvIntoBuf() - end")

	var buf = &bytes.Buffer{}
	written, metadata, err := f.Recv(buf)
	if err != nil {
		log.Errorf("DataChunkTransportWithCompression.RecvIntoBuf() got error: %v", err.Error())
	}

	// Debug
	log.Infof("metadata: %s", metadata.String())
	log.Infof("data: %s", buf.String())

	return written, buf, metadata, err
}
