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

// NewDataChunkFileOptions
func NewDataChunkFileOptions() *DataChunkFileOptions {
	return new(DataChunkFileOptions)
}

// SetHeader
func (opts *DataChunkFileOptions) SetHeader(header *Metadata) *DataChunkFileOptions {
	if opts == nil {
		return nil
	}
	opts.Header = header
	return opts
}

// GetHeader
func (opts *DataChunkFileOptions) GetHeader() *Metadata {
	if opts == nil {
		return nil
	}
	return opts.Header
}

// SetMetadata
func (opts *DataChunkFileOptions) SetMetadata(meta *Metadata) *DataChunkFileOptions {
	if opts == nil {
		return nil
	}
	opts.Metadata = meta
	return opts
}

// GetMetadata
func (opts *DataChunkFileOptions) GetMetadata() *Metadata {
	if opts == nil {
		return nil
	}
	return opts.Metadata
}

// SetCompress
func (opts *DataChunkFileOptions) SetCompress(compress bool) *DataChunkFileOptions {
	if opts == nil {
		return nil
	}
	opts.Compress = compress
	return opts
}

// GetCompress
func (opts *DataChunkFileOptions) GetCompress() bool {
	if opts == nil {
		return false
	}
	return opts.Compress
}

// SetDecompress
func (opts *DataChunkFileOptions) SetDecompress(decompress bool) *DataChunkFileOptions {
	if opts == nil {
		return nil
	}
	opts.Decompress = decompress
	return opts
}

// GetDecompress
func (opts *DataChunkFileOptions) GetDecompress() bool {
	if opts == nil {
		return false
	}
	return opts.Decompress
}
