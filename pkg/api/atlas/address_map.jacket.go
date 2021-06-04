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

// Len returns len of specified domain
func (m *AddressMap) Len(domains ...*Domain) int {
	switch len(domains) {
	case 0:
		// Len of the map
		return len(m.GetMap())
	case 1:
		// Len of particular domain
		return m.GetList(domains[0]).Len()
	case 2:
		// Len of particular domain within particular domain
		return m.GetList(domains[0]).Len(domains[1])
	}
	return 0
}

// Has checks whether specified domain exists
func (m *AddressMap) Has(domains ...*Domain) bool {
	switch len(domains) {
	case 1:
		return m.GetList(domains[0]) != nil
	case 2:
		return m.GetList(domains[0]).Has(domains[1])
	}
	return false
}

// Append wraps AddressList.Append
// Call example:
// Append([domain,] address0, [address1,...])
func (m *AddressMap) Append(entities ...interface{}) *AddressMap {
	domains, addresses := m.normalizeParams(true, entities...)
	if len(domains) > 0 {
		// Domain extracted, append provided addresses to specified domain
		m.EnsureList(domains[0]).Append(addresses...)
	}
	return m
}

// Set wraps Replace
func (m *AddressMap) Set(entities ...interface{}) *AddressMap {
	return m.Replace(entities...)
}

// normalizeParams builds list of domains and list of addresses out of provided entities
func (m *AddressMap) normalizeParams(insertDefaultDomain bool, entities ...interface{}) ([]*Domain, []*Address) {
	var domains []*Domain
	var addresses []*Address
	for _, entity := range entities {
		switch typed := entity.(type) {
		case *Domain:
			domains = append(domains, typed)
		case *Address:
			addresses = append(addresses, typed)
		}
	}

	// Introduce default domain if needed
	if insertDefaultDomain && (len(domains) == 0) {
		domains = append(domains, DomainThis)
	}

	return domains, addresses
}

// Replace wraps ReplaceList and ReplaceAddresses
// Call example:
// Replace([domain0, {domain1, domain2,... nested domains to be replaced with provided addresses}] address0[, address1,...])
func (m *AddressMap) Replace(entities ...interface{}) *AddressMap {
	domains, addresses := m.normalizeParams(true, entities...)
	switch len(domains) {
	case 0:
		// No domains specified, don't know what to do
		return m
	case 1:
		// Replace whole AddressList with specified addresses
		return m.ReplaceList(domains[0], addresses...)
	}

	// Replace some nested domains with specified addresses
	return m.ReplaceAddresses(domains[0], domains[1:], addresses...)
}

// ReplaceList replaces whole AddressList with specified addresses
func (m *AddressMap) ReplaceList(domain *Domain, addresses ...*Address) *AddressMap {
	m.NewList(domain).Append(addresses...)
	return m
}

// ReplaceAddresses replaces specified deleteDomains within domain with provided addresses
func (m *AddressMap) ReplaceAddresses(domain *Domain, deleteDomains []*Domain, addresses ...*Address) *AddressMap {
	return m.ReplaceList(
		domain,
		m.EnsureList(domain). // get current AddressList
					Exclude(deleteDomains...). // create new AddressList w/o deleteDomains. New AddressList can be nil
					Ensure().                  // but we definitely need to have this list, thus ensure it exists
					Append(addresses...).      // append to new AddressList w/o deleteDomains provided addresses
					All()...,                  // get AddressList as slice
	)
}

// First wraps AddressList.First
func (m *AddressMap) First(domains ...*Domain) *Address {
	switch len(domains) {
	case 0:
		// First of any available addresses
		if lists := m.GetLists(); lists != nil {
			for _, list := range lists {
				if first := list.First(); first != nil {
					return first
				}
			}
		}
	case 1:
		return m.GetList(domains[0]).First()
	case 2:
		return m.GetList(domains[0]).First(domains[1:]...)
	}
	return nil
}

// All wraps AddressList.All
func (m *AddressMap) All(domains ...*Domain) []*Address {
	switch len(domains) {
	case 1:
		return m.GetList(domains[0]).All()
	case 2:
		return m.GetList(domains[0]).All(domains[1:]...)
	}
	return nil
}
