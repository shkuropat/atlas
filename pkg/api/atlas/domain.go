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

// NewDomain creates new Domain
func NewDomain(name ...string) *Domain {
	d := new(Domain)
	if len(name) > 0 {
		d.Set(name[0])
	}
	return d
}

// Ensure returns new or existing Domain
func (m *Domain) Ensure() *Domain {
	if m == nil {
		return NewDomain()
	}
	return m
}

// Set sets Domain name
func (m *Domain) Set(name string) *Domain {
	if m == nil {
		return nil
	}
	m.Name = name
	return m
}

// Equals checks whether Domains are equal internally
func (m *Domain) Equals(domain *Domain) bool {
	if m == nil {
		return false
	}
	if domain == nil {
		return false
	}
	return m.GetName() == domain.GetName()
}

// String returns string representation of a Domain
func (m *Domain) String() string {
	if m == nil {
		return ""
	}
	return m.GetName()
}
