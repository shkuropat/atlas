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

// NewObjectsRequest
func NewObjectsRequest() *ObjectsRequest {
	return new(ObjectsRequest)
}

// EnsureHeader
func (m *ObjectsRequest) EnsureHeader() *Metadata {
	if m.HeaderOptional == nil {
		m.HeaderOptional = new(ObjectsRequest_Header)
	}
	if m.HeaderOptional.(*ObjectsRequest_Header).Header == nil {
		m.HeaderOptional.(*ObjectsRequest_Header).Header = new(Metadata)
	}
	return m.HeaderOptional.(*ObjectsRequest_Header).Header
}

// GetRequestsNum
func (m *ObjectsRequest) GetRequestsNum() int {
	return len(m.GetRequests())
}

// Append
func (m *ObjectsRequest) Append(request *ObjectRequest) *ObjectsRequest {
	if m == nil {
		return nil
	}
	m.Requests = append(m.Requests, request)
	return m
}

// String
func (m *ObjectsRequest) String() string {
	return "to be implemented"
}
