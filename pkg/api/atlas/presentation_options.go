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

// NewPresentationOptions
func NewPresentationOptions() *PresentationOptions {
	return new(PresentationOptions)
}

// HasEncoding
func (m *PresentationOptions) HasEncoding() bool {
	if m == nil {
		return false
	}
	return m.EncodingOptional != nil
}

// SetEncoding
func (m *PresentationOptions) SetEncoding(encoding *Encoding) *PresentationOptions {
	if m == nil {
		return nil
	}
	if m.EncodingOptional == nil {
		m.EncodingOptional = new(PresentationOptions_Encoding)
	}
	m.EncodingOptional.(*PresentationOptions_Encoding).Encoding = encoding
	return m
}

// HasCompression
func (m *PresentationOptions) HasCompression() bool {
	if m == nil {
		return false
	}
	return m.CompressionOptional != nil
}

// SetCompression
func (m *PresentationOptions) SetCompression(compression *Compression) *PresentationOptions {
	if m == nil {
		return nil
	}
	if m.CompressionOptional == nil {
		m.CompressionOptional = new(PresentationOptions_Compression)
	}
	m.CompressionOptional.(*PresentationOptions_Compression).Compression = compression
	return m
}

// HasDigest
func (m *PresentationOptions) HasDigest() bool {
	if m == nil {
		return false
	}
	return m.DigestOptional != nil
}

// SetDigest
func (m *PresentationOptions) SetDigest(digest *Digest) *PresentationOptions {
	if m == nil {
		return nil
	}
	if m.DigestOptional == nil {
		m.DigestOptional = new(PresentationOptions_Digest)
	}
	m.DigestOptional.(*PresentationOptions_Digest).Digest = digest
	return m
}

// String
func (m *PresentationOptions) String() string {
	return "to be implemented"
}
