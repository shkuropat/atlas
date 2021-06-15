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

// NewObjectRequest
func NewObjectRequest() *ObjectRequest {
	return new(ObjectRequest)
}

// EnsureHeader
func (m *ObjectRequest) EnsureHeader() *Metadata {
	if m == nil {
		return nil
	}
	if m.Header == nil {
		m.Header = NewMetadata()
	}
	return m.Header
}

// SetHeader
func (m *ObjectRequest) SetHeader(header *Metadata) *ObjectRequest {
	if m == nil {
		return nil
	}
	m.Header = header
	return m
}

// String
func (m *ObjectRequest) String() string {
	return "to be implemented"
}
