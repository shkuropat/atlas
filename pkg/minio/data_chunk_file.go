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

// RelayDataChunkFileIntoMinIO
func RelayDataChunkFileIntoMinIO(
	src atlas.DataChunkTransport,
	mi *MinIO,
	s3address *atlas.S3Address,
) (int64, *atlas.Metadata, error) {
	log.Infof("RelayDataChunkFileIntoMinIO() - start")
	defer log.Infof("RelayDataChunkFileIntoMinIO() - end")

	f, err := atlas.OpenDataChunkFile(src)
	if err != nil {
		log.Errorf("got error: %v", err)
		return 0, nil, err
	}
	defer f.Close()

	written, err := mi.Put(s3address.Bucket, s3address.Object, f)
	if err != nil {
		log.Errorf("RelayDataChunkFileIntoMinIO() got error: %v", err.Error())
	}
	f.PayloadMetadata.Log()

	return written, f.PayloadMetadata, err
}

// RelayDataChunkFileFromMinIO
func RelayDataChunkFileFromMinIO(
	dst atlas.DataChunkTransport,
	mi *MinIO,
	s3address *atlas.S3Address,
) (int64, error) {
	log.Infof("RelayDataChunkFileFromMinIO() - start")
	defer log.Infof("RelayDataChunkFileFromMinIO() - end")

	r, err := mi.Get(s3address.Bucket, s3address.Object)
	if err != nil {
		log.Errorf("got error from MinIO: %v", err)
		return 0, err
	}

	metadata := new(atlas.Metadata)
	metadata.SetFilename(s3address.Object)
	return atlas.SendDataChunkFile(dst, metadata, r, true)
}
