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

// NewAddressMap creates new AddressMap
func NewAddressMap() *AddressMap {
	return new(AddressMap)
}

// GetList gets specified AddressList of specified domain
func (m *AddressMap) GetList(domain *Domain) *AddressList {
	if mp := m.GetMap(); mp != nil {
		if list, ok := mp[domain.GetName()]; ok {
			return list
		}
	}
	return nil
}

// GetLists gets all AddressList from the AddressMap
func (m *AddressMap) GetLists() []*AddressList {
	if mp := m.GetMap(); mp != nil {
		var res []*AddressList
		for _, list := range mp {
			if res == nil {
				res = make([]*AddressList, m.Len())
			}
			res = append(res, list)
		}
		return res
	}
	return nil
}

// EnsureList makes sure AddressList of specified domain exists.
// It uses already existing domain AddressList or creates new if none found
func (m *AddressMap) EnsureList(domain *Domain) *AddressList {
	if m == nil {
		return nil
	}
	if m.Has(domain) {
		return m.GetList(domain)
	}
	return m.NewList(domain)
}

// NewList creates new AddressList of specified domain. Existing one will be overwritten.
func (m *AddressMap) NewList(domain *Domain) *AddressList {
	if m == nil {
		return nil
	}
	return m.SetList(domain, NewAddressList()).GetList(domain)
}

// SetList sets AddressList of specified domain. Existing one will be overwritten.
func (m *AddressMap) SetList(domain *Domain, list *AddressList) *AddressMap {
	if m == nil {
		return nil
	}
	m.ensureMap()
	m.Map[domain.GetName()] = list
	return m
}

// ensureMap makes sure map is created
func (m *AddressMap) ensureMap() {
	if m == nil {
		// Unable to ensure map inside nil struct
		return
	}
	if m.GetMap() == nil {
		// Ensure map exists
		m.Map = make(map[string]*AddressList)
	}
}

// String
func (m *AddressMap) String() string {
	return "no be implemented"
}
