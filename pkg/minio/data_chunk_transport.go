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
	"bytes"
	"fmt"
	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
)

type MinIODataChunkTransport struct {
	// MinIO handler
	mi *MinIO
	// Not used yet
	close bool
	// Name of the bucket where object located
	bucket string
	// Name of the object which contains data chunks
	object string
	// Slice of data chunks targeted as a new object
	chunks []string
}

// NewMinIODataChunkTransport
func NewMinIODataChunkTransport(mi *MinIO, bucket, object string, close bool) *MinIODataChunkTransport {
	return &MinIODataChunkTransport{
		mi:     mi,
		bucket: bucket,
		object: object,
		close:  close,
	}
}

// Close
func (t *MinIODataChunkTransport) Close() error {
	return t.compose()
}

// compose object out of chunks, if any
func (t *MinIODataChunkTransport) compose() error {
	// Compose single object out of slice of chunks targeted to be the object

	// We need to have at least 1 chunk to compose object from
	if len(t.chunks) < 1 {
		return nil
	}

	// Slice of sources.
	sources := make([]minio.SourceInfo, 0)
	for _, chunk := range t.chunks {
		sources = append(sources, minio.NewSourceInfo(t.bucket, chunk, nil))
	}

	// Create destination info
	dst, err := minio.NewDestinationInfo(t.bucket, t.object, nil, nil)
	if err != nil {
		return err
	}

	// Compose object by concatenating multiple source files.
	err = t.mi.client.ComposeObject(dst, sources)
	if err != nil {
		return err
	}

	return nil
}

// Send puts each data chunk into own uniq-UUID-named object in bucket and appends object to slice of chunks
func (t *MinIODataChunkTransport) Send(dataChunk *atlas.DataChunk) error {

	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	object := uuid.String()
	_, err = t.mi.Put(t.bucket, object, bytes.NewBuffer(dataChunk.GetBytes()))
	if err != nil {
		return err
	}
	t.chunks = append(t.chunks, object)

	return nil
}

// Recv
func (t *MinIODataChunkTransport) Recv() (*atlas.DataChunk, error) {
	return nil, fmt.Errorf("unimplemented MinIODataChunkTransport.Recv()")
}
