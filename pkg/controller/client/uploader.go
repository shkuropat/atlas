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

package controller_client

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// UploadFile sends file from client to service and receives response back (if any)
func UploadFile(client atlas.ControlPlaneClient, filename string, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadFile() - start")
	defer log.Info("UploadFile() - end")

	if _, err := os.Stat(filename); err != nil {
		log.Warnf("no file %s available err: %v", filename, err)
		return NewDataExchangeResultError(err)
	}

	log.Infof("Has file %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Warnf("ERROR open file %s err: %v", filename, err)
		return NewDataExchangeResultError(err)
	}

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(filepath.Base(filename))
	return UploadReader(client, f, options)
}

// UploadStdin sends STDIN from client to service and receives response back (if any)
func UploadStdin(client atlas.ControlPlaneClient, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadStdin() - start")
	defer log.Info("UploadStdin() - end")

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(os.Stdin.Name())
	return UploadReader(client, os.Stdin, options)
}

// UploadBytes
func UploadBytes(client atlas.ControlPlaneClient, data []byte, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadBytes() - start")
	defer log.Info("UploadBytes() - end")

	return UploadReader(client, bytes.NewReader(data), options)
}

// UploadReader
func UploadReader(client atlas.ControlPlaneClient, r io.Reader, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadReader() - start")
	defer log.Info("UploadReader() - end")

	result := Upload(client, r, options)
	if result.Err == nil {
		log.Infof("DONE send %s size %d", "io.Reader", result.Send.Data.Len)
	} else {
		log.Warnf("FAILED send %s size %d err %v", "io.Reader", result.Send.Data.Len, result.Err)
	}

	return result
}
