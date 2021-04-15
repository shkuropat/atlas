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
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	// Due to first enum value has to be zero in proto3
	TaskReserved int32 = 0
	// Unspecified
	TaskUnspecified int32 = 100
	// Echo request expects echo reply as an answer
	TaskEchoRequest int32 = 200
	// Echo reply is an answer to echo request
	TaskEchoReply int32 = 300
	// Request for configuration from the other party
	TaskConfigRequest int32 = 400
	// Configuration
	TaskConfig int32 = 500
	// Metrics schedule sends schedule by which metrics should be sent.
	TaskMetricsSchedule int32 = 600
	// Metrics request is an explicit request for metrics to be sent
	TaskMetricsRequest int32 = 700
	// One-time metrics
	TaskMetrics int32 = 800
	// Schedule to send data
	TaskDataSchedule int32 = 900
	// Explicit data request
	TaskDataRequest int32 = 1000
	// Data are coming
	TaskData int32 = 1100
	// Address is coming
	TaskAddress            int32 = 1200
	TaskExtract            int32 = 1300
	TaskExtractExecutables int32 = 1400
)

var TaskTypeEnum = NewEnum()

func init() {
	TaskTypeEnum.MustRegister("TaskReserved", TaskReserved)
	TaskTypeEnum.MustRegister("TaskUnspecified", TaskUnspecified)
	TaskTypeEnum.MustRegister("TaskEchoRequest", TaskEchoRequest)
	TaskTypeEnum.MustRegister("TaskEchoReply", TaskEchoReply)
	TaskTypeEnum.MustRegister("TaskConfigRequest", TaskConfigRequest)
	TaskTypeEnum.MustRegister("TaskConfig", TaskConfig)
	TaskTypeEnum.MustRegister("TaskMetricsSchedule", TaskMetricsSchedule)
	TaskTypeEnum.MustRegister("TaskMetricsRequest", TaskMetricsRequest)
	TaskTypeEnum.MustRegister("TaskMetrics", TaskMetrics)
	TaskTypeEnum.MustRegister("TaskDataSchedule", TaskDataSchedule)
	TaskTypeEnum.MustRegister("TaskDataRequest", TaskDataRequest)
	TaskTypeEnum.MustRegister("TaskData", TaskData)
	TaskTypeEnum.MustRegister("TaskAddress", TaskAddress)
	TaskTypeEnum.MustRegister("TaskExtract", TaskExtract)
	TaskTypeEnum.MustRegister("TaskExtractExecutables", TaskExtractExecutables)
}

// NewTask creates new Command with pre-allocated header
func NewTask() *Task {
	return &Task{
		Header: NewMetadata(),
	}
}

// NewTaskUnmarshalFrom creates new Task from bytes
func NewTaskUnmarshalFrom(bytes []byte) (*Task, error) {
	task := new(Task)
	if err := task.UnmarshalFrom(bytes); err != nil {
		return nil, err
	}
	return task, nil
}

// UnmarshalFrom unmarshal commands from bytes
func (m *Task) UnmarshalFrom(bytes []byte) error {
	return proto.Unmarshal(bytes, m)
}

// SetBytes sets payload bytes
func (m *Task) SetBytes(bytes []byte) *Task {
	m.Bytes = bytes
	return m
}

// SetPayload marshals msg as command's data
func (m *Task) SetPayload(msg proto.Message) error {
	if bytes, err := proto.Marshal(msg); err == nil {
		m.SetBytes(bytes)
		return nil
	} else {
		return err
	}
}

// GetPayload unmarshalls command's data into msg
func (m *Task) GetPayload(msg proto.Message) error {
	return proto.Unmarshal(m.GetBytes(), msg)
}

// AddSubject
func (m *Task) AddSubject(subject *Metadata) *Task {
	m.Subjects = append(m.Subjects, subject)
	return m
}

// AddSubjects
func (m *Task) AddSubjects(subjects ...*Metadata) *Task {
	m.Subjects = append(m.Subjects, subjects...)
	return m
}

// AddTask
func (m *Task) AddTask(task *Task) *Task {
	m.Tasks = append(m.Tasks, task)
	return m
}

// AddTasks
func (m *Task) AddTasks(tasks ...*Task) *Task {
	m.Tasks = append(m.Tasks, tasks...)
	return m
}

// ShiftTasks fetches first (0-indexed) task from available tasks.
// fetched task is removed from list of tasks
func (m *Task) ShiftTasks() *Task {
	var task *Task = nil
	if len(m.Tasks) > 0 {
		task = m.Tasks[0]
		m.Tasks = m.Tasks[1:]
	}
	return task
}

// Derive produces derivative task from the task as:
//   1. fetches first (0-indexed) task from available tasks
//   2. and attaches all the rest tasks (if any) as tasks of the fetched one
// Original task is modified.
func (m *Task) Derive() *Task {
	root := m.ShiftTasks()
	if root == nil {
		return nil
	}

	root.Tasks = m.Tasks
	return root
}

// String
func (m *Task) String() string {
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
