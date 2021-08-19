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

// NewObjectsList
func NewObjectsList() *ObjectsList {
	return new(ObjectsList)
}

// SetStatus
func (m *ObjectsList) SetStatus(status *Status) *ObjectsList {
	if m == nil {
		return nil
	}
	m.Status = status
	return m
}

// AddReport
func (m *ObjectsList) AddReport(reports ...*Report) *ObjectsList {
	if m == nil {
		return nil
	}
	m.Reports = append(m.Reports, reports...)
	return m
}

// AddTask
func (m *ObjectsList) AddTask(tasks ...*Task) *ObjectsList {
	if m == nil {
		return nil
	}
	m.Tasks = append(m.Tasks, tasks...)
	return m
}

// AddStatus
func (m *ObjectsList) AddStatus(statuses ...*Status) *ObjectsList {
	if m == nil {
		return nil
	}
	m.Statuses = append(m.Statuses, statuses...)
	return m
}

// AddObjectStatus
func (m *ObjectsList) AddObjectStatus(statuses ...*ObjectStatus) *ObjectsList {
	if m == nil {
		return nil
	}
	m.ObjectStatuses = append(m.ObjectStatuses, statuses...)
	return m
}

// AddFile
func (m *ObjectsList) AddFile(files ...*File) *ObjectsList {
	if m == nil {
		return nil
	}
	m.Files = append(m.Files, files...)
	return m
}

// LenReports
func (m *ObjectsList) LenReports() int {
	if m == nil {
		return 0
	}
	return len(m.Reports)
}

// LenTasks
func (m *ObjectsList) LenTasks() int {
	if m == nil {
		return 0
	}
	return len(m.Tasks)
}

// LenStatuses
func (m *ObjectsList) LenStatuses() int {
	if m == nil {
		return 0
	}
	return len(m.Statuses)
}

// LenObjectStatuses
func (m *ObjectsList) LenObjectStatuses() int {
	if m == nil {
		return 0
	}
	return len(m.ObjectStatuses)
}

// LenFiles
func (m *ObjectsList) LenFiles() int {
	if m == nil {
		return 0
	}
	return len(m.Files)
}

/*
// First
func (m *ReportMulti) First() *Report {
	if m.Len() > 0 {
		return m.Reports[0]
	}
	return nil
}

// Shift
func (m *ReportMulti) Shift() *Report {
	if m.Len() > 0 {
		r := m.Reports[0]
		m.Reports = m.Reports[1:]
		return r
	}
	return nil
}
*/

// String
func (m *ObjectsList) String() string {
	return "to be implemented"
}
