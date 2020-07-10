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

// DataChunkFileAdapter
type DataChunkFileAdapter struct {
	transport DataChunkTransport
	options   *DataChunkFileAdapterOptions
}

// NewDataChunkFileAdapter
func NewDataChunkFileAdapter(transport DataChunkTransport, options *DataChunkFileAdapterOptions) *DataChunkFileAdapter {
	return &DataChunkFileAdapter{
		transport: transport,
		options:   options,
	}
}

// AcceptFrom sends data into adapter from `src`
func (f *DataChunkFileAdapter) AcceptFrom(src io.Reader, metadata *Metadata) (int64, error) {
	log.Infof("DataChunkTransportWithCompression.Send() - start")
	defer log.Infof("DataChunkTransportWithCompression.Send() - end")

	w, err := OpenDataChunkFileCompressor(
		f.transport,
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
		f.options.GetCompress(),
	)
	if err != nil {
		log.Warnf("got err: %v", err)
		return 0, err
	}
	defer w.Close()

	return io.Copy(w, src)
}

// RelayInto gets data from adapter into `dst`
func (f *DataChunkFileAdapter) RelayInto(dst io.Writer) (int64, *Metadata, error) {
	log.Infof("DataChunkTransportWithCompression.Recv() - start")
	defer log.Infof("DataChunkTransportWithCompression.Recv() - end")

	r, err := OpenDataChunkFileDecompressor(f.transport, f.options.GetDecompress())
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

// RelayIntoBuf gets data from adapter into newly created buffer
func (f *DataChunkFileAdapter) RelayIntoBuf() (int64, *bytes.Buffer, *Metadata, error) {
	log.Infof("DataChunkTransportWithCompression.RecvIntoBuf() - start")
	defer log.Infof("DataChunkTransportWithCompression.RecvIntoBuf() - end")

	var buf = &bytes.Buffer{}
	written, metadata, err := f.RelayInto(buf)
	if err != nil {
		log.Errorf("DataChunkTransportWithCompression.RecvIntoBuf() got error: %v", err.Error())
	}

	// Debug
	log.Infof("metadata: %s", metadata.String())
	log.Infof("data: %s", buf.String())

	return written, buf, metadata, err
}
