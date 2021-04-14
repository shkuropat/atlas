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

// Compressor is a compression descriptor
type Compressor struct {
	ReadCompression  *Compression
	LZMAReader       *lzma.Reader
	WriteCompression *Compression
	LZMAWriter       *lzma.Writer
}

// NewCompressor
func NewCompressor(
	read *Compression,
	reader io.Reader,
	write *Compression,
	writer io.Writer,
) (*Compressor, error) {
	res := &Compressor{}

	switch read.GetType() {
	case CompressionTypeLZMA:
		res.ReadCompression = read
		lzmaReader, err := lzma.NewReader(reader)
		if err != nil {
			log.Warnf("FAILED to create lzma reader. err: %v", err)
			return nil, err
		}
		res.LZMAReader = lzmaReader
	default:
		res.ReadCompression = nil

	}

	switch write.GetType() {
	case CompressionTypeLZMA:
		lzmaWriter, err := lzma.NewWriter(writer)
		if err != nil {
			log.Warnf("FAILED to create lzma writer. err: %v", err)
			return nil, err
		}
		res.WriteCompression = write
		res.LZMAWriter = lzmaWriter
	default:
		res.WriteCompression = nil
	}

	return res, nil
}

// Close
func (c *Compressor) Close() error {
	if c == nil {
		return nil
	}
	if c.LZMAWriter != nil {
		return c.LZMAWriter.Close()
	}

	return nil
}

// WriteEnabled
func (c *Compressor) WriteEnabled() bool {
	if c == nil {
		return false
	}
	return c.WriteCompression != nil
}

// ReadEnabled
func (c *Compressor) ReadEnabled() bool {
	if c == nil {
		return false
	}
	return c.ReadCompression != nil
}

// Write
func (c *Compressor) Write(p []byte) (n int, err error) {
	if c == nil {
		return 0, fmt.Errorf("can't write to empty")
	}

	return c.LZMAWriter.Write(p)
}

// Read
func (c *Compressor) Read(p []byte) (n int, err error) {
	if c == nil {
		return 0, fmt.Errorf("can't read from empty")
	}

	return c.LZMAReader.Read(p)
}
