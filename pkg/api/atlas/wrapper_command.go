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

var CommandTypeEnum = NewEnum()

func init() {
	CommandTypeEnum.MustRegister("COMMAND_RESERVED", int32(CommandType_COMMAND_RESERVED))
	CommandTypeEnum.MustRegister("COMMAND_UNSPECIFIED", int32(CommandType_COMMAND_UNSPECIFIED))
	CommandTypeEnum.MustRegister("COMMAND_ECHO_REQUEST", int32(CommandType_COMMAND_ECHO_REQUEST))
	CommandTypeEnum.MustRegister("COMMAND_ECHO_REPLY", int32(CommandType_COMMAND_ECHO_REPLY))
	CommandTypeEnum.MustRegister("COMMAND_CONFIG_REQUEST", int32(CommandType_COMMAND_CONFIG_REQUEST))
	CommandTypeEnum.MustRegister("COMMAND_CONFIG", int32(CommandType_COMMAND_CONFIG))
	CommandTypeEnum.MustRegister("COMMAND_METRICS_SCHEDULE", int32(CommandType_COMMAND_METRICS_SCHEDULE))
	CommandTypeEnum.MustRegister("COMMAND_METRICS_REQUEST", int32(CommandType_COMMAND_METRICS_REQUEST))
	CommandTypeEnum.MustRegister("COMMAND_METRICS", int32(CommandType_COMMAND_METRICS))
	CommandTypeEnum.MustRegister("COMMAND_DATA_SCHEDULE", int32(CommandType_COMMAND_DATA_SCHEDULE))
	CommandTypeEnum.MustRegister("COMMAND_DATA_REQUEST", int32(CommandType_COMMAND_DATA_REQUEST))
	CommandTypeEnum.MustRegister("COMMAND_DATA", int32(CommandType_COMMAND_DATA))
	CommandTypeEnum.MustRegister("COMMAND_ADDRESS", int32(CommandType_COMMAND_ADDRESS))
	CommandTypeEnum.MustRegister("COMMAND_EXTRACT", int32(CommandType_COMMAND_EXTRACT))
	CommandTypeEnum.MustRegister("COMMAND_EXTRACT_EXECUTABLES", int32(CommandType_COMMAND_EXTRACT_EXECUTABLES))
}

// NewCommand creates new Command with pre-allocated header
func NewCommand() *Command {
	return &Command{
		Header: NewMetadata(),
	}
}

// NewCommandUnmarshalFrom creates new Command from bytes
func NewCommandUnmarshalFrom(bytes []byte) (*Command, error) {
	cmd := &Command{}
	if err := cmd.UnmarshalFrom(bytes); err != nil {
		return nil, err
	}
	return cmd, nil
}

// GetType gets command type
func (m *Command) GetType() CommandType {
	return CommandType(m.GetHeader().GetType())
}

// SetType
func (m *Command) SetType(_type CommandType) *Command {
	m.GetHeader().SetType(int32(_type))
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

// GetID
func (m *Command) GetID() *UUID {
	return m.GetHeader().GetId()
}

// GetIDAsString
func (m *Command) GetIDAsString() string {
	if id := m.GetHeader().GetId(); id == nil {
		return ""
	} else {
		return id.GetString()
	}
}

// SetID
func (m *Command) SetID(id *UUID) *Command {
	m.GetHeader().SetID(id)
	return m
}

// SetIDFromString
func (m *Command) SetIDFromString(id string) *Command {
	m.GetHeader().SetID(NewUUID().SetString(id))
	return m
}

// CreateID
func (m *Command) CreateID() *Command {
	return m.SetID(CreateUUID())
}

// GetReferenceID
func (m *Command) GetReferenceID() *UUID {
	return m.GetHeader().GetReferenceId()
}

// GetReferenceIDAsString
func (m *Command) GetReferenceIDAsString() string {
	if id := m.GetReferenceID(); id == nil {
		return ""
	} else {
		return id.GetString()
	}
}

// SetReferenceID
func (m *Command) SetReferenceID(id *UUID) *Command {
	m.GetHeader().SetReferenceID(id)
	return m
}

// SetReferenceIDFromString
func (m *Command) SetReferenceIDFromString(id string) *Command {
	m.GetHeader().SetReferenceIDFromString(id)
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

// GetPayload unmarshalls command's payload
func (m *Command) GetPayload(msg proto.Message) error {
	return proto.Unmarshal(m.GetBytes(), msg)
}

// AddSubject
func (m *Command) AddSubject(subject *Metadata) *Command {
	m.Subjects = append(m.Subjects, subject)
	return m
}

// AddSubjects
func (m *Command) AddSubjects(subjects ...*Metadata) *Command {
	m.Subjects = append(m.Subjects, subjects...)
	return m
}

// AddCommand
func (m *Command) AddCommand(command *Command) *Command {
	m.Commands = append(m.Commands, command)
	return m
}

// AddCommands
func (m *Command) AddCommands(commands ...*Command) *Command {
	m.Commands = append(m.Commands, commands...)
	return m
}

// ShiftCommands fetches first (0-indexed) command from available commands
func (m *Command) ShiftCommands() *Command {
	var cmd *Command = nil
	if len(m.Commands) > 0 {
		cmd = m.Commands[0]
		m.Commands = m.Commands[1:]
	}
	return cmd
}

// Shift fetches first (0-indexed) command from available commands and attaches all the rest commands (if any) as
// commands of the fetched one, so commands can be Shift-ed from this returned command
func (m *Command) Shift() *Command {
	root := m.ShiftCommands()
	if root == nil {
		return nil
	}

	root.Commands = m.Commands
	return root
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
	if len(m.GetSubjects()) > 0 {
		parts = append(parts, "subjects:")
		for _, subj := range m.GetSubjects() {
			parts = append(parts, subj.String())
		}
	}

	return strings.Join(parts, " ")
}
