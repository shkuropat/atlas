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

// NewObjectStatus
func NewObjectStatus(status ...*Status) *ObjectStatus {
	d := new(ObjectStatus)
	if len(status) > 0 {
		d.SetStatus(status[0])
	}
	return d
}

// Ensure returns new or existing Status
func (m *ObjectStatus) Ensure() *ObjectStatus {
	if m == nil {
		return NewObjectStatus()
	}
	return m
}

// SetStatus sets status
func (m *ObjectStatus) SetStatus(status *Status) *ObjectStatus {
	if m == nil {
		return nil
	}
	m.Status = status
	return m
}

// SetDomain sets address
func (m *ObjectStatus) SetDomain(domain *Domain) *ObjectStatus {
	if m == nil {
		return nil
	}
	if m.DomainOptional == nil {
		m.DomainOptional = new(ObjectStatus_Domain)
	}
	m.DomainOptional.(*ObjectStatus_Domain).Domain = domain
	return m
}

// SetAddress sets address
func (m *ObjectStatus) SetAddress(address *Address) *ObjectStatus {
	if m == nil {
		return nil
	}
	if m.AddressOptional == nil {
		m.AddressOptional = new(ObjectStatus_Address)
	}
	m.AddressOptional.(*ObjectStatus_Address).Address = address
	return m
}

// String
func (m *ObjectStatus) String() string {
	return "to be implemented"
}
