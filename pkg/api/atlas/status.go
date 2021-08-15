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

// StatusType represents all types of statuses
const (
	// Due to first enum value has to be zero in proto3
	StatusReserved = 0
	// Unspecified means we do not know its type
	StatusUnspecified = 100
	// Object found
	StatusOK = 200
	// Object created
	StatusCreated = 201
	// Object accepted
	StatusAccepted = 202
	// Not all parts/objects requested were found
	StatusPartial = 206
	// All objects found
	StatusFoundAll = 220
	// Object is in progress of something
	StatusInProgress = 230
	// Object moved to other location
	StatusMovedPermanently = 301
	// Object not found
	StatusNotFound = 404
	// Object not ready
	StatusNotReady = 405
	// Object has failed due to internal error
	StatusInternalError = 500
	// Object failed somehow
	StatusFailed = 550
)

var StatusTypeEnum = NewEnum()

func init() {
	StatusTypeEnum.MustRegister("StatusReserved", StatusReserved)
	StatusTypeEnum.MustRegister("StatusUnspecified", StatusUnspecified)
	StatusTypeEnum.MustRegister("StatusOK", StatusOK)
	StatusTypeEnum.MustRegister("StatusCreated", StatusCreated)
	StatusTypeEnum.MustRegister("StatusAccepted", StatusAccepted)
	StatusTypeEnum.MustRegister("StatusPartial", StatusPartial)
	StatusTypeEnum.MustRegister("StatusFoundAll", StatusFoundAll)
	StatusTypeEnum.MustRegister("StatusInProgress", StatusInProgress)
	StatusTypeEnum.MustRegister("StatusMovedPermanently", StatusMovedPermanently)
	StatusTypeEnum.MustRegister("StatusNotFound", StatusNotFound)
	StatusTypeEnum.MustRegister("StatusNotReady", StatusNotReady)
	StatusTypeEnum.MustRegister("StatusInternalError", StatusInternalError)
	StatusTypeEnum.MustRegister("StatusFailed", StatusFailed)
}

// NewStatus
func NewStatus(status ...int32) *Status {
	d := new(Status)
	if len(status) > 0 {
		d.SetStatus(status[0])
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

// SetStatus sets status
func (m *Status) SetStatus(status int32) *Status {
	if m == nil {
		return nil
	}
	m.Status = status
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
	return m.GetStatus() == status.GetStatus()
}

// String
func (m *Status) String() string {
	return "to be implemented"
}
