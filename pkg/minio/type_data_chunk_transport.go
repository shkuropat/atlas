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
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// DataChunkTransport defines transport level interface
// Has the following functions:
//   Send(*DataChunk) error
//   Recv() (*DataChunk, error)

type DataChunkTransport struct {
	Accessor
}

// NewDataChunkTransport
func NewDataChunkTransport(mi *MinIO, s3address *atlas.S3Address) *DataChunkTransport {
	log.Tracef("minio.NewDataChunkTransport() - start")
	defer log.Tracef("minio.NewDataChunkTransport() - end")

	if accessor, err := NewAccessor(mi, s3address); err != nil {
		return nil
	} else {
		return &DataChunkTransport{
			Accessor: *accessor,
		}
	}
}

// Send puts each data chunk into own unique-UUID-named object in bucket and appends object to slice of chunks
func (t *DataChunkTransport) Send(dataChunk *atlas.DataChunk) error {
	log.Tracef("minio.DataChunkTransport.Send() - start")
	defer log.Tracef("minio.DataChunkTransport.Send() - end")

	_, err := t.writeChunk(dataChunk.GetData())
	return err
}

// Recv
func (t *DataChunkTransport) Recv() (*atlas.DataChunk, error) {
	log.Tracef("minio.DataChunkTransport.Recv() - start")
	defer log.Tracef("minio.DataChunkTransport.Recv() - end")

	return nil, fmt.Errorf("unimplemented minio.DataChunkTransport.Recv()")
}
