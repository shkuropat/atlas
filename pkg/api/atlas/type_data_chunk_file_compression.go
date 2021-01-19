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
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/ulikunitz/xz/lzma"
)

type CompressionType string

const (
	CompressionNone CompressionType = "none"
	CompressionLZMA CompressionType = "lzma"
)

// String
func (c CompressionType) String() string {
	return string(c)
}

// ParseCompressionType
func ParseCompressionType(_type string) CompressionType {
	switch _type {
	case "lzma":
		return CompressionLZMA
	default:
		return CompressionNone
	}
}

// DataChunkFileCompression is a compression descriptor
type DataChunkFileCompression struct {
	ReadType   CompressionType
	LZMAReader *lzma.Reader
	WriteType  CompressionType
	LZMAWriter *lzma.Writer
}

// NewDataChunkFileCompression
func NewDataChunkFileCompression(
	read CompressionType,
	reader io.Reader,
	write CompressionType,
	writer io.Writer,
) (*DataChunkFileCompression, error) {
	res := &DataChunkFileCompression{}

	switch read {
	case CompressionLZMA:
		res.ReadType = read
		lzmaReader, err := lzma.NewReader(reader)
		if err != nil {
			log.Warnf("FAILED to create lzma reader. err: %v", err)
			return nil, err
		}
		res.LZMAReader = lzmaReader
	default:
		res.ReadType = CompressionNone

	}

	switch write {
	case CompressionLZMA:
		lzmaWriter, err := lzma.NewWriter(writer)
		if err != nil {
			log.Warnf("FAILED to create lzma writer. err: %v", err)
			return nil, err
		}
		res.WriteType = CompressionLZMA
		res.LZMAWriter = lzmaWriter
	default:
		res.WriteType = CompressionNone
	}

	return res, nil
}

// Close
func (c *DataChunkFileCompression) Close() error {
	if c == nil {
		return nil
	}
	if c.LZMAWriter != nil {
		return c.LZMAWriter.Close()
	}

	return nil
}

// WriteEnabled
func (c *DataChunkFileCompression) WriteEnabled() bool {
	if c == nil {
		return false
	}
	return c.WriteType != CompressionNone
}

// ReadEnabled
func (c *DataChunkFileCompression) ReadEnabled() bool {
	if c == nil {
		return false
	}
	return c.ReadType != CompressionNone
}

// Write
func (c *DataChunkFileCompression) Write(p []byte) (n int, err error) {
	if c == nil {
		return 0, fmt.Errorf("can't write to empty")
	}

	return c.LZMAWriter.Write(p)
}

// Read
func (c *DataChunkFileCompression) Read(p []byte) (n int, err error) {
	if c == nil {
		return 0, fmt.Errorf("can't read from empty")
	}

	return c.LZMAReader.Read(p)
}
