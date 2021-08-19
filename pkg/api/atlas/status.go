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

// NewStatus
func NewStatus(code ...int32) *Status {
	d := new(Status)
	if len(code) > 0 {
		d.SetCode(code[0])
	}
	return d
}

// Ensure returns new or existing Status
func (m *Status) Ensure() *Status {
	if m == nil {
		return NewStatus()
	}
	return m
}

// SetCode sets status
func (m *Status) SetCode(code int32) *Status {
	if m == nil {
		return nil
	}
	m.Code = code
	return m
}

// Equals checks whether two objects are equal internally
func (m *Status) Equals(status *Status) bool {
	if m == nil {
		return false
	}
	if status == nil {
		return false
	}
	return m.GetCode() == status.GetCode()
}

// String
func (m *Status) String() string {
	return "to be implemented"
}
