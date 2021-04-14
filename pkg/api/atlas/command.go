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

const (
	// Due to first enum value has to be zero in proto3
	CommandReserved int32 = 0
	// Unspecified
	CommandUnspecified int32 = 100
	// Echo request expects echo reply as an answer
	CommandEchoRequest int32 = 200
	// Echo reply is an answer to echo request
	CommandEchoReply int32 = 300
	// Request for configuration from the other party
	CommandConfigRequest int32 = 400
	// Configuration
	CommandConfig int32 = 500
	// Metrics schedule sends schedule by which metrics should be sent.
	CommandMetricsSchedule int32 = 600
	// Metrics request is an explicit request for metrics to be sent
	CommandMetricsRequest int32 = 700
	// One-time metrics
	CommandMetrics int32 = 800
	// Schedule to send data
	CommandDataSchedule int32 = 900
	// Explicit data request
	CommandDataRequest int32 = 1000
	// Data are coming
	CommandData int32 = 1100
	// Address is coming
	CommandAddress            int32 = 1200
	CommandExtract            int32 = 1300
	CommandExtractExecutables int32 = 1400
)

var CommandTypeEnum = NewEnum()

func init() {
	CommandTypeEnum.MustRegister("CommandReserved", CommandReserved)
	CommandTypeEnum.MustRegister("CommandUnspecified", CommandUnspecified)
	CommandTypeEnum.MustRegister("CommandEchoRequest", CommandEchoRequest)
	CommandTypeEnum.MustRegister("CommandEchoReply", CommandEchoReply)
	CommandTypeEnum.MustRegister("CommandConfigRequest", CommandConfigRequest)
	CommandTypeEnum.MustRegister("CommandConfig", CommandConfig)
	CommandTypeEnum.MustRegister("CommandMetricsSchedule", CommandMetricsSchedule)
	CommandTypeEnum.MustRegister("CommandMetricsRequest", CommandMetricsRequest)
	CommandTypeEnum.MustRegister("CommandMetrics", CommandMetrics)
	CommandTypeEnum.MustRegister("CommandDataSchedule", CommandDataSchedule)
	CommandTypeEnum.MustRegister("CommandDataRequest", CommandDataRequest)
	CommandTypeEnum.MustRegister("CommandData", CommandData)
	CommandTypeEnum.MustRegister("CommandAddress", CommandAddress)
	CommandTypeEnum.MustRegister("CommandExtract", CommandExtract)
	CommandTypeEnum.MustRegister("CommandExtractExecutables", CommandExtractExecutables)
}

// NewCommand creates new Command with pre-allocated header
func NewCommand() *Command {
	return &Command{
		Header: NewMetadata(),
	}
}

// NewCommandUnmarshalFrom creates new Command from bytes
func NewCommandUnmarshalFrom(bytes []byte) (*Command, error) {
	cmd := new(Command)
	if err := cmd.UnmarshalFrom(bytes); err != nil {
		return nil, err
	}
	return cmd, nil
}

// UnmarshalFrom unmarshal commands from bytes
func (m *Command) UnmarshalFrom(bytes []byte) error {
	return proto.Unmarshal(bytes, m)
}

// SetBytes sets payload bytes
func (m *Command) SetBytes(bytes []byte) *Command {
	m.Bytes = bytes
	return m
}

// SetPayload marshals msg as command's data
func (m *Command) SetPayload(msg proto.Message) error {
	if bytes, err := proto.Marshal(msg); err == nil {
		m.SetBytes(bytes)
		return nil
	} else {
		return err
	}
}

// GetPayload unmarshalls command's data into msg
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

// String
func (m *Command) String() string {
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
