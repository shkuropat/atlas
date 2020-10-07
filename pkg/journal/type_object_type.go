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

type ObjectType uint8

const (
	ObjectType1     ObjectType = 1
	ObjectType1Name string     = "ObjectType1"
)

var (
	objectTypeName = map[ObjectType]string{
		ObjectType1: ObjectType1Name,
	}
	objectTypeValue = map[string]ObjectType{
		ObjectType1Name: ObjectType1,
	}
)

var (
	ErrObjectTypeBusy = fmt.Errorf("busy")
)

// RegisterObjectType
func RegisterObjectType(_type ObjectType, name string) error {
	// Check availability
	if n, ok := objectTypeName[_type]; ok {
		if n != name {
			return ErrObjectTypeBusy
		}
	}
	if v, ok := objectTypeValue[name]; ok {
		if v != _type {
			return ErrEndpointIDBusy
		}
	}

	// Register
	objectTypeName[_type] = name
	objectTypeValue[name] = _type

	return nil
}
