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

	"github.com/binarly-io/atlas/pkg/minio"
	log "github.com/sirupsen/logrus"
)

// SendDataChunkFile
func SendDataChunkFile(
	dst DataChunkTransport,
	metadata *Metadata,
	src io.Reader,
	compress bool,
) (int64, error) {
	log.Infof("SendDataChunkFile() - start")
	defer log.Infof("SendDataChunkFile() - end")

	w, err := OpenDataChunkFileWriter(
		dst,
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
		compress,
	)
	if err != nil {
		log.Warnf("got err: %v", err)
		return 0, err
	}
	defer w.Close()

	return io.Copy(w, src)
}

// RecvDataChunkFile
func RecvDataChunkFile(src DataChunkTransport, dst io.Writer, decompress bool) (int64, *Metadata, error) {
	log.Infof("RecvDataChunkFile() - start")
	defer log.Infof("RecvDataChunkFile() - end")

	r, err := OpenDataChunkFileReader(src, decompress)
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

// RecvDataChunkFileIntoBuf
func RecvDataChunkFileIntoBuf(src DataChunkTransport) (int64, *bytes.Buffer, *Metadata, error) {
	log.Infof("RecvDataChunkFileIntoBuf() - start")
	defer log.Infof("RecvDataChunkFileIntoBuf() - end")

	var buf = &bytes.Buffer{}
	written, metadata, err := RecvDataChunkFile(src, buf, true)
	if err != nil {
		log.Errorf("RecvDataChunkFileIntoBuf() got error: %v", err.Error())
	}

	// Debug
	log.Infof("metadata: %s", metadata.String())
	log.Infof("data: %s", buf.String())

	return written, buf, metadata, err
}

// RelayDataChunkFileIntoMinIO
func RelayDataChunkFileIntoMinIO(
	src DataChunkTransport,
	mi *minio.MinIO,
	bucketName string,
	objectName string,
) (int64, *Metadata, error) {
	log.Infof("RelayDataChunkFileIntoMinIO() - start")
	defer log.Infof("RelayDataChunkFileIntoMinIO() - end")

	f, err := OpenDataChunkFile(src)
	if err != nil {
		log.Errorf("got error: %v", err)
		return 0, nil, err
	}
	defer f.Close()

	written, err := mi.Put(bucketName, objectName, f)
	if err != nil {
		log.Errorf("RelayDataChunkFileIntoMinIO() got error: %v", err.Error())
	}
	f.PayloadMetadata.Log()

	return written, f.PayloadMetadata, err
}

// RelayDataChunkFileFromMinIO
func RelayDataChunkFileFromMinIO(
	dst DataChunkTransport,
	mi *minio.MinIO,
	bucketName string,
	objectName string,
) (int64, error) {
	log.Infof("RelayDataChunkFileFromMinIO() - start")
	defer log.Infof("RelayDataChunkFileFromMinIO() - end")

	r, err := mi.Get(bucketName, objectName)
	if err != nil {
		log.Errorf("got error from MinIO: %v", err)
		return 0, err
	}

	metadata := new(Metadata)
	metadata.SetFilename(objectName)
	return SendDataChunkFile(dst, metadata, r, true)
}
