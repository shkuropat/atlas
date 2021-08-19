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

var (
	StatusReserved         = NewStatus(StatusCodeReserved)
	StatusUnspecified      = NewStatus(StatusCodeUnspecified)
	StatusOK               = NewStatus(StatusCodeOK)
	StatusCreated          = NewStatus(StatusCodeCreated)
	StatusAccepted         = NewStatus(StatusCodeAccepted)
	StatusPartial          = NewStatus(StatusCodePartial)
	StatusFoundAll         = NewStatus(StatusCodeFoundAll)
	StatusInProgress       = NewStatus(StatusCodeInProgress)
	StatusMovedPermanently = NewStatus(StatusCodeMovedPermanently)
	StatusNotFound         = NewStatus(StatusCodeNotFound)
	StatusNotReady         = NewStatus(StatusCodeNotReady)
	StatusInternalError    = NewStatus(StatusCodeInternalError)
	StatusFailed           = NewStatus(StatusCodeFailed)

	Statuses = []*Status{
		StatusReserved,
		StatusUnspecified,
		StatusOK,
		StatusCreated,
		StatusAccepted,
		StatusPartial,
		StatusFoundAll,
		StatusInProgress,
		StatusMovedPermanently,
		StatusNotFound,
		StatusNotReady,
		StatusInternalError,
		StatusFailed,
	}
)

// RegisterStatus tries to register specified Status.
// New entity must be non-equal to all registered entities.
// Returns nil in case new entity can not be registered, say it is equal to previously registered entity
func RegisterStatus(status *Status) *Status {
	if FindStatus(status) != nil {
		// Such a domain already exists
		return nil
	}
	Statuses = append(Statuses, status)
	return status
}

// MustRegisterStatus the same as RegisterStatus but with panic
func MustRegisterStatus(status *Status) {
	if RegisterStatus(status) == nil {
		panic("unable to register status")
	}
}

// FindStatus returns registered status with the same string value as provided
func FindStatus(status *Status) *Status {
	return StatusFromCode(status.Code)
}

// NormalizeStatus returns either registered status with the same string value as provided status or provided status.
func NormalizeStatus(status *Status) *Status {
	if f := FindStatus(status); f != nil {
		// Return registered
		return f
	}
	// Unable to find registered, return provided
	return status
}

// StatusFromCode tries to find registered status with specified code
func StatusFromCode(code int32) *Status {
	s := NewStatus(code)
	for _, status := range Statuses {
		if status.Equals(s) {
			return status
		}
	}
	return nil
}
