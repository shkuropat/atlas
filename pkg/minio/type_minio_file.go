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
	"github.com/binarly-io/atlas/pkg/api/atlas"
	log "github.com/sirupsen/logrus"
	"io"
)

// File
// Inspired by os.File handler and is expected to be used in the same context.
type File struct {
	Accessor
}

// OpenFile
func OpenFile(mi *MinIO, s3address *atlas.S3Address) (*File, error) {
	log.Tracef("minio.OpenFile() - start")
	defer log.Tracef("minio.OpenFile() - end")

	if accessor, err := NewAccessor(mi, s3address); err != nil {
		return nil, err
	} else {
		return &File{
			Accessor: *accessor,
		}, nil
	}
}

// Read
func (f *File) Read(p []byte) (int, error) {
	log.Tracef("minio.File.Read() - start")
	defer log.Tracef("minio.File.Read() - end")

	log.Errorf("unimplemented method minio.File.Read()")
	return 0, fmt.Errorf("unimplemented method minio.File.Read()")
}

// ReadFrom reads data from src
func (f *File) ReadFrom(src io.Reader) (int64, error) {
	log.Tracef("minio.File.ReadFrom() - start")
	defer log.Tracef("minio.File.ReadFrom() - end")

	return f.mi.Put(f.s3address.Bucket, f.s3address.Object, src)
}

// Write
func (f *File) Write(p []byte) (int, error) {
	log.Tracef("minio.File.Write() - start")
	defer log.Tracef("minio.File.Write() - end")

	return f.writeChunk(p)
}

// WriteTo writes data to dst
func (f *File) WriteTo(dst io.Writer) (int64, error) {
	log.Tracef("minio.File.WriteTo() - start")
	defer log.Tracef("minio.File.WriteTo() - end")

	r, err := f.mi.Get(f.s3address.Bucket, f.s3address.Object)
	if err != nil {
		log.Errorf("got error from MinIO: %v", err)
		return 0, err
	}

	return io.Copy(dst, r)
}
