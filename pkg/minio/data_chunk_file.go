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

// AcceptDataChunkFile
func AcceptDataChunkFile(
	mi *MinIO,
	s3address *atlas.S3Address,
	src atlas.DataChunkTransport,
	options *Options,
) (int64, *atlas.Metadata, error) {
	log.Infof("AcceptDataChunkFile() - start")
	defer log.Infof("AcceptDataChunkFile() - end")

	r, err := atlas.OpenDataChunkFileReader(src, options.GetDecompress())
	if err != nil {
		log.Errorf("got error: %v", err)
		return 0, nil, err
	}
	defer r.Close()

	written, err := mi.Put(s3address.Bucket, s3address.Object, r)
	if err != nil {
		log.Errorf("AcceptDataChunkFile() got error: %v", err.Error())
	}
	r.DataChunkFile.PayloadMetadata.Log()

	return written, r.DataChunkFile.PayloadMetadata, err
}

// FetchDataChunkFile
func FetchDataChunkFile(
	dst atlas.DataChunkTransport,
	mi *MinIO,
	s3address *atlas.S3Address,
	options *Options,
) (int64, error) {
	log.Infof("FetchDataChunkFile() - start")
	defer log.Infof("FetchDataChunkFile() - end")

	r, err := mi.Get(s3address.Bucket, s3address.Object)
	if err != nil {
		log.Errorf("got error from MinIO: %v", err)
		return 0, err
	}

	metadata := new(atlas.Metadata)
	metadata.SetFilename(s3address.Object)
	return atlas.SendDataChunkFile(dst, metadata, r, options.GetCompress())
}
