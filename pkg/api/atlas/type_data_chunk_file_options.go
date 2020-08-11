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

// DataChunkFileOptions
type DataChunkFileOptions struct {
	Header   *Metadata
	Metadata *Metadata

	// Compress outgoing data
	Compress bool

	// Decompress incoming data
	Decompress bool
}

// GetHeader
func (opts *DataChunkFileOptions) GetHeader() *Metadata {
	if opts == nil {
		return nil
	}
	return opts.Header
}

// GetMetadata
func (opts *DataChunkFileOptions) GetMetadata() *Metadata {
	if opts == nil {
		return nil
	}
	return opts.Metadata
}

// GetCompress
func (opts *DataChunkFileOptions) GetCompress() bool {
	if opts == nil {
		return false
	}
	return opts.Compress
}

// GetDecompress
func (opts *DataChunkFileOptions) GetDecompress() bool {
	if opts == nil {
		return false
	}
	return opts.Decompress
}
