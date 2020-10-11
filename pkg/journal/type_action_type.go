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

package journal

import "fmt"

type ActionType uint8

const (
	ActionUnknown              ActionType = 0
	ActionUnknownName          string     = "UNKNOWN"
	ActionRequestStart         ActionType = 1
	ActionRequestStartName     string     = "ActionRequestStart"
	ActionSaveData             ActionType = 2
	ActionSaveDataName         string     = "ActionSaveData"
	ActionSaveDataError        ActionType = 3
	ActionSaveDataErrorName    string     = "ActionSaveDataError"
	ActionProcessData          ActionType = 4
	ActionProcessDataName      string     = "ActionProcessData"
	ActionProcessDataError     ActionType = 5
	ActionProcessDataErrorName string     = "ActionProcessDataError"
	ActionRequestCompleted     ActionType = 6
	ActionRequestCompletedName string     = "ActionRequestCompleted"
	ActionRequestError         ActionType = 7
	ActionRequestErrorName     string     = "ActionRequestError"
)

var (
	actionName = map[ActionType]string{
		ActionUnknown:          ActionUnknownName,
		ActionRequestStart:     ActionRequestStartName,
		ActionSaveData:         ActionSaveDataName,
		ActionSaveDataError:    ActionSaveDataErrorName,
		ActionProcessData:      ActionProcessDataName,
		ActionProcessDataError: ActionProcessDataErrorName,
		ActionRequestCompleted: ActionRequestCompletedName,
		ActionRequestError:     ActionRequestErrorName,
	}

	actionValue = map[string]ActionType{
		ActionUnknownName:          ActionUnknown,
		ActionRequestStartName:     ActionRequestStart,
		ActionSaveDataName:         ActionSaveData,
		ActionSaveDataErrorName:    ActionSaveDataError,
		ActionProcessDataName:      ActionProcessData,
		ActionProcessDataErrorName: ActionProcessDataError,
		ActionRequestCompletedName: ActionRequestCompleted,
		ActionRequestErrorName:     ActionRequestError,
	}
)

var (
	ErrActionTypeBusy = fmt.Errorf("busy")
)

// RegisterActionType
func RegisterActionType(_type ActionType, name string) error {
	// Check availability
	if n, ok := actionName[_type]; ok {
		if n != name {
			return ErrActionTypeBusy
		}
	}
	if v, ok := actionValue[name]; ok {
		if v != _type {
			return ErrActionTypeBusy
		}
	}

	// Register
	actionName[_type] = name
	actionValue[name] = _type

	return nil
}
