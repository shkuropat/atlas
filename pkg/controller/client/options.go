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
	// Compress specifies whether to compress data on send
	Compress bool
	// Decompress specifies whether to decompress data on receive
	Decompress bool
	// WaitReply specifies whether to wait for answer/reply
	WaitReply bool
	// Metadata describes data stream
	Metadata *atlas.Metadata
}

// GetCompress
func (opts *DataExchangeOptions) GetCompress() bool {
	if opts == nil {
		return false
	}

	return opts.Compress
}

// GetDecompress
func (opts *DataExchangeOptions) GetDecompress() bool {
	if opts == nil {
		return false
	}

	return opts.Decompress
}

// GetWaitReply
func (opts *DataExchangeOptions) GetWaitReply() bool {
	if opts == nil {
		return false
	}

	return opts.WaitReply
}

// GetMetadata
func (opts *DataExchangeOptions) GetMetadata() *atlas.Metadata {
	if opts == nil {
		return nil
	}

	return opts.Metadata
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
	if opts.Metadata == nil {
		opts.Metadata = new(atlas.Metadata)
	}

	return opts.Metadata
}
