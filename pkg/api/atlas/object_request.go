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

// SetRequestDomain
func (m *ObjectRequest) SetRequestDomain(domain *Domain) *ObjectRequest {
	if m == nil {
		return nil
	}
	if m.RequestDomainOptional == nil {
		m.RequestDomainOptional = new(ObjectRequest_RequestDomain)
	}
	m.RequestDomainOptional.(*ObjectRequest_RequestDomain).RequestDomain = domain
	return m
}

// SetResultDomain
func (m *ObjectRequest) SetResultDomain(domain *Domain) *ObjectRequest {
	if m.ResultDomainOptional == nil {
		m.ResultDomainOptional = new(ObjectRequest_ResultDomain)
	}
	m.ResultDomainOptional.(*ObjectRequest_ResultDomain).ResultDomain = domain
	return m
}

// SetAddress
func (m *ObjectRequest) SetAddress(address *Address) *ObjectRequest {
	if m == nil {
		return nil
	}
	m.Address = address
	return m
}

// String
func (m *ObjectRequest) String() string {
	return "to be implemented"
}
