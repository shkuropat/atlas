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

import "github.com/binarly-io/atlas/pkg/api/atlas"

// DataExchangeOptions
type DataExchangeOptions struct {
	// compress specifies whether to compress data on send
	compress bool
	// decompress specifies whether to decompress data on receive
	decompress bool
	// waitReply specifies whether to wait for answer/reply
	waitReply bool
	// metadata describes data stream
	metadata *atlas.Metadata
}

// NewDataExchangeOptions
func NewDataExchangeOptions() *DataExchangeOptions {
	return new(DataExchangeOptions)
}

// SetCompress
func (opts *DataExchangeOptions) SetCompress(compress bool) *DataExchangeOptions {
	if opts == nil {
		return nil
	}
	opts.compress = compress
	return opts
}

// GetCompress
func (opts *DataExchangeOptions) GetCompress() bool {
	if opts == nil {
		return false
	}
	return opts.compress
}

// SetDecompress
func (opts *DataExchangeOptions) SetDecompress(decompress bool) *DataExchangeOptions {
	if opts == nil {
		return nil
	}
	opts.decompress = decompress
	return opts
}

// GetDecompress
func (opts *DataExchangeOptions) GetDecompress() bool {
	if opts == nil {
		return false
	}
	return opts.decompress
}

// SetWaitReply
func (opts *DataExchangeOptions) SetWaitReply(waitReply bool) *DataExchangeOptions {
	if opts == nil {
		return nil
	}
	opts.waitReply = waitReply
	return opts
}

// GetWaitReply
func (opts *DataExchangeOptions) GetWaitReply() bool {
	if opts == nil {
		return false
	}
	return opts.waitReply
}

// SetMetadata
func (opts *DataExchangeOptions) SetMetadata(meta *atlas.Metadata) *DataExchangeOptions {
	if opts == nil {
		return nil
	}
	opts.metadata = meta
	return opts
}

// GetMetadata
func (opts *DataExchangeOptions) GetMetadata() *atlas.Metadata {
	if opts == nil {
		return nil
	}
	return opts.metadata
}

// Ensure
func (opts *DataExchangeOptions) Ensure() *DataExchangeOptions {
	if opts == nil {
		return &DataExchangeOptions{}
	}
	return opts
}

// EnsureMetadata
func (opts *DataExchangeOptions) EnsureMetadata() *atlas.Metadata {
	if opts == nil {
		return nil
	}
	if opts.metadata == nil {
		opts.metadata = new(atlas.Metadata)
	}
	return opts.metadata
}
