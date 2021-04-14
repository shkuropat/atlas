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

// NewDomain
func NewDomain() *Domain {
	return new(Domain)
}

// Set
func (m *Domain) Set(name string) *Domain {
	if m == nil {
		return nil
	}
	m.Name = name
	return m
}

// Equals
func (m *Domain) Equals(domain *Domain) bool {
	if m == nil {
		return false
	}
	if domain == nil {
		return false
	}
	return m.GetName() == domain.GetName()
}

// String
func (m *Domain) String() string {
	if m == nil {
		return ""
	}
	return m.GetName()
}
