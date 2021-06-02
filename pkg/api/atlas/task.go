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

// NewTaskUnmarshalFrom creates new Task from a slice of bytes
func NewTaskUnmarshalFrom(bytes []byte) (*Task, error) {
	task := new(Task)
	if err := task.UnmarshalFrom(bytes); err != nil {
		return nil, err
	}
	return task, nil
}

// UnmarshalFrom unmarshal commands from a slice of bytes
func (m *Task) UnmarshalFrom(bytes []byte) error {
	return proto.Unmarshal(bytes, m)
}

// SetBytes sets bytes as task's data. Provided bytes are not interpreted and used as-is.
func (m *Task) SetBytes(bytes []byte) *Task {
	m.Bytes = bytes
	return m
}

// SetPayload puts any protobuf message (type) into task's data.
// Message is marshalled into binary form and set as data bytes of the task.
func (m *Task) SetPayload(msg proto.Message) error {
	if bytes, err := proto.Marshal(msg); err == nil {
		m.SetBytes(bytes)
		return nil
	} else {
		return err
	}
}

// GetPayload extracts profobuf message (type) from task's data.
// Message is unmarshalled from task's data into provided message.
func (m *Task) GetPayload(msg proto.Message) error {
	return proto.Unmarshal(m.GetBytes(), msg)
}

// AddSubject ands one subject to the task
func (m *Task) AddSubject(subject *Metadata) *Task {
	m.Subjects = append(m.Subjects, subject)
	return m
}

// AddSubjects adds multiple subjects to the task
func (m *Task) AddSubjects(subjects ...*Metadata) *Task {
	m.Subjects = append(m.Subjects, subjects...)
	return m
}

// AddSubtask adds one subtask to the task
func (m *Task) AddSubtask(task *Task) *Task {
	m.Children = append(m.Children, task)
	return m
}

// AddSubtasks adds multiple subtasks to the task
func (m *Task) AddSubtasks(tasks ...*Task) *Task {
	m.Children = append(m.Children, tasks...)
	return m
}

// FirstSubtask fetches first (0-indexed) subtask. List of subtasks does not change.
func (m *Task) FirstSubtask() *Task {
	if m == nil {
		return nil
	}
	if len(m.Children) == 0 {
		return nil
	}
	return m.Children[0]
}

// LastSubtask fetches last subtask. List of subtasks does not change.
func (m *Task) LastSubtask() *Task {
	if m == nil {
		return nil
	}
	if len(m.Children) == 0 {
		return nil
	}
	return m.Children[len(m.Children)-1]
}

// ShiftSubtasks fetches first (0-indexed) task from available tasks.
// Fetched task is removed from the list of tasks.
// List of subtasks changes.
func (m *Task) ShiftSubtasks() *Task {
	var task *Task = nil
	if len(m.Children) > 0 {
		task = m.Children[0]
		m.Children = m.Children[1:]
	}
	return task
}

// Derive produces derivative task from the task as:
//   1. fetches first (0-indexed) subtask from available subtasks
//   2. and attaches all the rest subtasks (if any) as subtasks of the fetched one, which it the new top now.
// Original task is modified.
func (m *Task) Derive() *Task {
	// Assume new root task is the first subtask of current task
	root := m.FirstSubtask()
	if root == nil {
		return nil
	}

	// Parent of the new task is current task
	root.AddParent(m)
	// Subtasks of the new task are the same as of the current task except the rrot itself
	root.Children = m.Children
	root.ShiftSubtasks()
	return root
}

// AddParent adds one parent of the task
func (m *Task) AddParent(task *Task) *Task {
	m.Parents = append(m.Parents, task)
	return m
}

// AddParents adds multiple parents of the task
func (m *Task) AddParents(tasks ...*Task) *Task {
	m.Parents = append(m.Parents, tasks...)
	return m
}

// FirstParent fetches first (0-indexed) parent. List of list does not change.
func (m *Task) FirstParent() *Task {
	if m == nil {
		return nil
	}
	if len(m.Parents) == 0 {
		return nil
	}
	return m.Parents[0]
}

// LastParent fetches last parent. List of list does not change.
func (m *Task) LastParent() *Task {
	if m == nil {
		return nil
	}
	if len(m.Parents) == 0 {
		return nil
	}
	return m.Parents[len(m.Parents)-1]
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

	if len(m.Children) > 0 {
		parts = append(parts, fmt.Sprintf("%d children", len(m.Children)))
	}

	if len(m.Parents) > 0 {
		parts = append(parts, fmt.Sprintf("%d parents", len(m.Parents)))
	}

	return strings.Join(parts, " ")
}
