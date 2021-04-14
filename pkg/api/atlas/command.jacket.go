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

// GetType gets command type
func (m *Command) GetType() int32 {
	return m.GetHeader().GetType()
}

// SetType
func (m *Command) SetType(_type int32) *Command {
	m.GetHeader().SetType(_type)
	return m
}

// GetName gets command name
func (m *Command) GetName() string {
	return m.GetHeader().GetName()
}

// SetName
func (m *Command) SetName(name string) *Command {
	m.GetHeader().SetName(name)
	return m
}

// GetUUID
func (m *Command) GetUUID() *UUID {
	return m.GetHeader().GetAddresses().First(DomainThis, DomainUUID).GetUuid()
}

// GetIDAsString
func (m *Command) GetUUIDAsString() string {
	return m.GetUUID().String()
}

// SetUUID
func (m *Command) SetUUID(address *Address) *Command {
	m.GetHeader().EnsureAddresses().Set(DomainThis, DomainUUID, address)
	return m
}

// SetUUIDFromString
func (m *Command) SetUUIDFromString(id string) *Command {
	m.SetUUID(NewAddressUUIDFromString(id, DomainUUID))
	return m
}

// CreateUUID creates new random UUID
func (m *Command) CreateUUID() *Command {
	return m.SetUUID(NewAddressUUIDRandom(DomainUUID))
}

// GetReferenceUUID
func (m *Command) GetReferenceUUID() *UUID {
	return m.GetHeader().GetAddresses().First(DomainReference, DomainUUID).GetUuid()
}

// GetReferenceIDAsString
func (m *Command) GetReferenceUUIDAsString() string {
	return m.GetReferenceUUID().String()
}

// SetReferenceUUID
func (m *Command) SetReferenceUUID(address *Address) *Command {
	m.GetHeader().EnsureAddresses().Set(DomainReference, DomainUUID, address)
	return m
}

// SetReferenceIDFromString
func (m *Command) SetReferenceUUIDFromString(id string) *Command {
	m.SetReferenceUUID(NewAddressUUIDFromString(id, DomainUUID))
	return m
}

// GetDescription
func (m *Command) GetDescription() string {
	return m.GetHeader().GetDescription()
}

// SetDescription
func (m *Command) SetDescription(description string) *Command {
	m.GetHeader().SetDescription(description)
	return m
}
