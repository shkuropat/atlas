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
	"io"
)

// File
// Inspired by os.File handler and is expected to be used in the same context.
type File struct {
	mi        *MinIO
	s3address *atlas.S3Address

	// Slice of data chunks targeted as a new object
	chunks []string
}

// OpenFile
func OpenFile(mi *MinIO, s3address *atlas.S3Address) (*File, error) {
	// Sanity check
	if (mi == nil) || (s3address == nil) {
		return nil, fmt.Errorf("minio.OpenFile() requires full object address to be specfied")
	}

	// TODO ping MinIO?

	// All seems to be good, create file
	return &File{
		mi:        mi,
		s3address: s3address,
	}, nil
}

// Close
func (f *File) Close() error {
	log.Infof("minio.File.Close() - start")
	defer log.Infof("minio.File.Close() - end")

	return f.compose()
}

// Read
func (f *File) Read(p []byte) (int, error) {
	log.Infof("minio.File.Read() - start")
	defer log.Infof("minio.File.Read() - end")

	log.Errorf("unimplemented method minio.File.Read()")
	return 0, fmt.Errorf("unimplemented method minio.File.Read()")
}

// ReadFrom reads data from src
func (f *File) ReadFrom(src io.Reader) (int64, error) {
	log.Infof("minio.File.ReadFrom() - start")
	defer log.Infof("minio.File.ReadFrom() - end")

	return f.mi.Put(f.s3address.Bucket, f.s3address.Object, src)
}

// Write
func (f *File) Write(p []byte) (int, error) {
	log.Infof("minio.File.Write() - start")
	defer log.Infof("minio.File.Write() - end")

	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("unable to put create UUID. err:%v", err)
		return 0, err
	}
	object := uuid.String()
	n, err := f.mi.Put(f.s3address.Bucket, f.s3address.Object, bytes.NewBuffer(p))
	if err != nil {
		log.Errorf("unable to put chunk. err:%v", err)
		return int(n), err
	}
	f.chunks = append(f.chunks, object)

	return int(n), err
}

// WriteTo writes data to dst
func (f *File) WriteTo(dst io.Writer) (int64, error) {
	log.Infof("minio.File.WriteTo() - start")
	defer log.Infof("minio.File.WriteTo() - end")

	r, err := f.mi.Get(f.s3address.Bucket, f.s3address.Object)
	if err != nil {
		log.Errorf("got error from MinIO: %v", err)
		return 0, err
	}

	return io.Copy(dst, r)
}

// compose object out of chunks, if any
func (f *File) compose() error {
	// Compose single object out of slice of chunks targeted to be the object
	log.Infof("minio.File.compose() - start")
	defer log.Infof("minio.File.compose() - end")

	// We need to have at least 1 chunk to compose object from
	if len(f.chunks) < 1 {
		return nil
	}

	log.Infof("compose object out of %d chunks", len(f.chunks))

	// Slice of sources.
	sources := make([]minio.SourceInfo, 0)
	for _, chunk := range f.chunks {
		sources = append(sources, minio.NewSourceInfo(f.s3address.Bucket, chunk, nil))
	}

	// Create destination info
	dst, err := minio.NewDestinationInfo(f.s3address.Bucket, f.s3address.Object, nil, nil)
	if err != nil {
		log.Errorf("unable to make DestinationInfo() err:%v", err)
		return err
	}

	// Chunks are obsoleted from this moment
	f.chunks = nil

	// Compose object by concatenating multiple source files.
	err = f.mi.client.ComposeObject(dst, sources)
	if err != nil {
		log.Errorf("unable to ComposeObject() err:%v", err)
		return err
	}

	return nil
}
