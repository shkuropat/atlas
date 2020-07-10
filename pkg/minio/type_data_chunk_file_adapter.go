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

package minio

import (
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

type DataChunkFileAdapter struct {
	mi        *MinIO
	s3address *atlas.S3Address
	options   *atlas.DataChunkFileAdapterOptions
}

func NewDataChunkFileAdapter(mi *MinIO, s3address *atlas.S3Address, options *atlas.DataChunkFileAdapterOptions) *DataChunkFileAdapter {
	return &DataChunkFileAdapter{
		mi:        mi,
		s3address: s3address,
		options:   options,
	}
}

// AcceptFrom sends data into adapter from `src`
func (f *DataChunkFileAdapter) AcceptFrom(src atlas.DataChunkTransport) (int64, *atlas.Metadata, error) {
	log.Infof("AcceptDataChunkFile() - start")
	defer log.Infof("AcceptDataChunkFile() - end")

	r, err := atlas.OpenDataChunkFileDecompressor(src, f.options.GetDecompress())
	if err != nil {
		log.Errorf("got error: %v", err)
		return 0, nil, err
	}
	defer r.Close()

	written, err := f.mi.Put(f.s3address.Bucket, f.s3address.Object, r)
	if err != nil {
		log.Errorf("AcceptDataChunkFile() got error: %v", err.Error())
	}

	log.Info("Accepted data meta:")
	r.DataChunkFile.PayloadMetadata.Log()
	r.DataChunkFile.TransportMetadata.Log()

	return written, r.DataChunkFile.PayloadMetadata, err
}

// RelayInto gets data from adapter into `dst`
func (f *DataChunkFileAdapter) RelayInto(dst atlas.DataChunkTransport) (int64, error) {
	log.Infof("FetchDataChunkFile() - start")
	defer log.Infof("FetchDataChunkFile() - end")

	r, err := f.mi.Get(f.s3address.Bucket, f.s3address.Object)
	if err != nil {
		log.Errorf("got error from MinIO: %v", err)
		return 0, err
	}

	t := atlas.NewDataChunkFileAdapter(dst, &atlas.DataChunkFileAdapterOptions{
		Decompress: f.options.GetCompress(),
	})
	metadata := new(atlas.Metadata)
	metadata.SetFilename(f.s3address.Object)
	return t.AcceptFrom(r, metadata)
}
