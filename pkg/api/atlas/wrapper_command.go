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

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
)

func NewCommand() *Command {
	return &Command{
		Header: NewMetadata(),
	}
}

// GetType gets command type
func (m *Command) GetType() CommandType {
	return CommandType(m.GetHeader().GetType())
}

// SetType
func (m *Command) SetType(_type CommandType) *Command {
	m.GetHeader().SetType(MetadataType(_type))
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
func (m *Command) GetUUID() string {
	if uuid := m.GetHeader().GetUuid(); uuid == nil {
		return ""
	} else {
		return uuid.GetStringValue()
	}
}

// SetUUID
func (m *Command) SetUUID(uuid string) *Command {
	m.GetHeader().SetUUID(uuid)
	return m
}

// CreateUUID
func (m *Command) CreateUUID() *Command {
	return m.SetUUID(CreateNewUUID())
}

// GetUUIDReference
func (m *Command) GetUUIDReference() string {
	if uuid := m.GetHeader().GetUuidReference(); uuid == nil {
		return ""
	} else {
		return uuid.GetStringValue()
	}
}

// SetUUIDReference
func (m *Command) SetUUIDReference(uuid string) *Command {
	m.GetHeader().SetUUIDReference(uuid)
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

// SetBytes sets payload bytes
func (m *Command) SetBytes(bytes []byte) *Command {
	m.Bytes = bytes
	return m
}

// UnmarshalFrom unmarshal commands from bytes
func (m *Command) UnmarshalFrom(bytes []byte) error {
	return proto.Unmarshal(bytes, m)
}

// SetPayload marshals data as command's payload
func (m *Command) SetPayload(msg proto.Message) error {
	if bytes, err := proto.Marshal(msg); err == nil {
		m.SetBytes(bytes)
		return nil
	} else {
		return err
	}
}

// GetPayload unmarshals command's payload
func (m *Command) GetPayload(msg proto.Message) error {
	return proto.Unmarshal(m.GetBytes(), msg)
}

// AddAddresses
func (m *Command) AddAddresses(addresses ...*S3Address) *Command {
	m.Addresses = append(m.Addresses, addresses...)

	return m
}

// Printable
func (m *Command) Printable() string {
	if m == nil {
		return "nil"
	}

	var parts []string
	if _type := m.GetType(); _type > 0 {
		parts = append(parts, "type:"+fmt.Sprintf("%d", _type))
	}
	if name := m.GetName(); name != "" {
		parts = append(parts, "name:"+name)
	}
	if len(m.GetAddresses()) > 0 {
		parts = append(parts, "addresses:")
		for _, address := range m.GetAddresses() {
			parts = append(parts, address.Printable())
		}
	}

	return strings.Join(parts, " ")
}
