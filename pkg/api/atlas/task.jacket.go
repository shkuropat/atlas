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

//
// Wrap metadata
//

// GetType gets task type
func (m *Task) GetType() int32 {
	return m.GetHeader().GetType()
}

// SetType
func (m *Task) SetType(_type int32) *Task {
	m.GetHeader().SetType(_type)
	return m
}

// GetName gets task name
func (m *Task) GetName() string {
	return m.GetHeader().GetName()
}

// SetName
func (m *Task) SetName(name string) *Task {
	m.GetHeader().SetName(name)
	return m
}

// GetUUID
func (m *Task) GetUUID() *UUID {
	return m.GetHeader().GetAddresses().First(DomainThis, DomainUUID).GetUUID()
}

// GetUUIDAsString
func (m *Task) GetUUIDAsString() string {
	return m.GetUUID().String()
}

// SetUUID
func (m *Task) SetUUID(address *Address) *Task {
	m.GetHeader().Set(DomainThis, DomainUUID, address)
	return m
}

// SetUUIDFromString
func (m *Task) SetUUIDFromString(id string) *Task {
	m.SetUUID(NewAddressUUIDFromString(id, DomainUUID))
	return m
}

// CreateUUID creates new random UUID
func (m *Task) CreateUUID() *Task {
	return m.SetUUID(NewAddressUUIDRandom(DomainUUID))
}

// GetReferenceUUID
func (m *Task) GetReferenceUUID() *UUID {
	return m.GetHeader().GetAddresses().First(DomainReference, DomainUUID).GetUUID()
}

// GetReferenceUUIDAsString
func (m *Task) GetReferenceUUIDAsString() string {
	return m.GetReferenceUUID().String()
}

// SetReferenceUUID
func (m *Task) SetReferenceUUID(uuid *UUID) *Task {
	m.GetHeader().EnsureAddresses().Set(DomainReference, DomainUUID, uuid)
	return m
}

// SetReferenceUUIDFromString
func (m *Task) SetReferenceUUIDFromString(id string) *Task {
	m.SetReferenceUUID(NewUUIDFromString(id))
	return m
}

// GetDescription
func (m *Task) GetDescription() string {
	return m.GetHeader().GetDescription()
}

// SetDescription
func (m *Task) SetDescription(description string) *Task {
	m.GetHeader().SetDescription(description)
	return m
}
