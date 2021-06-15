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

// EnsureDomain
func (m *ObjectsRequest) EnsureDomain() *Domain {
	if m == nil {
		return nil
	}
	if m.DomainOptional == nil {
		m.DomainOptional = new(ObjectsRequest_Domain)
	}
	if m.DomainOptional.(*ObjectsRequest_Domain).Domain == nil {
		m.DomainOptional.(*ObjectsRequest_Domain).Domain = NewDomain()
	}
	return m.GetDomain()
}

// SetDomain
func (m *ObjectsRequest) SetDomain(domain *Domain) *ObjectsRequest {
	if m == nil {
		return nil
	}
	if m.DomainOptional == nil {
		m.DomainOptional = new(ObjectsRequest_Domain)
	}
	m.DomainOptional.(*ObjectsRequest_Domain).Domain = domain
	return m
}

// EnsureMode
func (m *ObjectsRequest) EnsureMode() RequestMode {
	if m == nil {
		return RequestMode_RESERVED
	}
	if m.RequestModeOptional == nil {
		m.RequestModeOptional = new(ObjectsRequest_RequestMode)
	}
	return m.GetRequestMode()
}

// SetMode
func (m *ObjectsRequest) SetMode(mode RequestMode) *ObjectsRequest {
	if m == nil {
		return nil
	}
	m.EnsureMode()
	m.RequestModeOptional.(*ObjectsRequest_RequestMode).RequestMode = mode
	return m
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
