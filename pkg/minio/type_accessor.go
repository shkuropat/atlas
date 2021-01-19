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
	log "github.com/sirupsen/logrus"
)

type Accessor struct {
	mi        *MinIO
	s3address *atlas.S3Address

	// Chunks names targeted as a new destination object. Used for incoming
	chunks []string
}

const (
	maxChunks = 1000
)

// NewAccessor
func NewAccessor(mi *MinIO, s3address *atlas.S3Address) (*Accessor, error) {
	log.Tracef("minio.NewAccessor() - start")
	defer log.Tracef("minio.NewAccessor() - end")

	// Sanity check
	if (mi == nil) || (s3address == nil) {
		return nil, fmt.Errorf("minio.OpenFile() requires full object address to be specfied")
	}

	// TODO ping MinIO?

	// All seems to be good, create file
	return &Accessor{
		mi:        mi,
		s3address: s3address,
	}, nil
}

// Close
func (a *Accessor) Close() error {
	log.Tracef("minio.Accessor.Close() - start")
	defer log.Tracef("minio.Accessor.Close() - end")

	return a.composeObject()
}

// writeChunk writes chunk of data as separate object
func (a *Accessor) writeChunk(data []byte) (int, error) {
	log.Tracef("minio.Accessor.writeChunk() - start")
	defer log.Tracef("minio.Accessor.writeChunk() - end")

	chunkObjectName := a.makeUniqueChunkName()
	n, err := a.mi.Put(a.s3address.Bucket, chunkObjectName, bytes.NewBuffer(data))
	if err != nil {
		log.Errorf("unable to put chunk. err:%v", err)
		return int(n), err
	}

	a.chunks = append(a.chunks, chunkObjectName)
	if len(a.chunks) > maxChunks {
		// TODO compact
		// there is limitation in 10K chunks in
		// err = f.mi.client.ComposeObject(dst, sources)
		// need to implement compaction
	}
	log.Infof("put chunk %s/%s size %d", a.s3address.Bucket, chunkObjectName, n)

	return int(n), err
}

// makeUniqueChunkName makes unique object name for this DataChunk
func (a *Accessor) makeUniqueChunkName() string {
	var (
		uuidPart   string
		indexPart  string
		objectPart string
	)

	if _uuid, err := uuid.NewUUID(); err == nil {
		uuidPart = _uuid.String()
	}
	indexPart = fmt.Sprintf("%d", len(a.chunks))
	objectPart = a.s3address.Object
	return fmt.Sprintf("%s_%s_%s", objectPart, indexPart, uuidPart)
}

// composeObject composes new destination object out of chunks
func (a *Accessor) composeObject() error {
	// Compose single object out of slice of chunks targeted to be the object
	log.Tracef("minio.Accessor.composeObject() - start")
	defer log.Tracef("minio.Accessor.composeObject() - end")

	// We need to have at least 1 chunk to compose object from
	if len(a.chunks) < 1 {
		return nil
	}

	log.Infof("compose object out of %d chunks", len(a.chunks))

	// Slice of sources.
	sources := make([]minio.SourceInfo, 0)
	for _, chunk := range a.chunks {
		sources = append(sources, minio.NewSourceInfo(a.s3address.Bucket, chunk, nil))
	}

	// Create destination info
	dst, err := minio.NewDestinationInfo(a.s3address.Bucket, a.s3address.Object, nil, nil)
	if err != nil {
		log.Errorf("unable to make DestinationInfo() err:%v", err)
		return err
	}

	// Compose object by concatenating multiple source files.
	err = a.mi.client.ComposeObject(dst, sources)
	if err != nil {
		log.Errorf("unable to ComposeObject() err:%v", err)
		return err
	}
	log.Infof("composed object %s/%s", a.s3address.Bucket, a.s3address.Object)

	for _, chunk := range a.chunks {
		if err = a.mi.client.RemoveObject(a.s3address.Bucket, chunk); err != nil {
			log.Errorf("unable to RemoveObject() %s/%s err:%v", a.s3address.Bucket, chunk, err)
		}
	}

	// Chunks are empty from this moment, since we've just created object out of these chunks
	a.chunks = nil

	return nil
}
