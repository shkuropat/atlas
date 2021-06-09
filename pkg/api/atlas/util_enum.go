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
	"sort"
	"strings"
)

// Enum
type Enum struct {
	value2Name map[int32]string
	name2Value map[string]int32
}

// NewEnum
func NewEnum() *Enum {
	return &Enum{
		value2Name: make(map[int32]string),
		name2Value: make(map[string]int32),
	}
}

// Has
func (e *Enum) Has(what interface{}) bool {
	switch typed := what.(type) {
	case int32:
		_, ok := e.value2Name[typed]
		return ok
	case string:
		_, ok := e.name2Value[typed]
		return ok
	}
	return false
}

// HasLower
func (e *Enum) HasLower(name string) bool {
	return e.Has(strings.ToLower(name))
}

// HasUpper
func (e *Enum) HasUpper(name string) bool {
	return e.Has(strings.ToUpper(name))
}

// GetValue
func (e *Enum) GetValue(what interface{}) (int32, bool) {
	switch typed := what.(type) {
	case int32:
		// Verify value is known
		_, ok := e.value2Name[typed]
		return typed, ok
	case string:
		// Find value by the name
		v, ok := e.name2Value[typed]
		return v, ok
	}
	return 0, false
}

// Value
func (e *Enum) Value(what interface{}) int32 {
	if value, found := e.GetValue(what); found {
		return value
	} else {
		return 0
	}
}

// MustGetValue
func (e *Enum) MustGetValue(what interface{}) int32 {
	if v, ok := e.GetValue(what); ok {
		return v
	}
	panic("MustGetValue")
}

// GetName
func (e *Enum) GetName(what interface{}) (string, bool) {
	switch typed := what.(type) {
	case int32:
		// Find name by the value
		name, ok := e.value2Name[typed]
		return name, ok
	case string:
		// Verify name is known
		_, ok := e.name2Value[typed]
		return typed, ok
	}
	return "", false
}

// name
func (e *Enum) Name(what interface{}) string {
	if name, found := e.GetName(what); found {
		return name
	} else {
		return "unknown enum entry"
	}
}

// MustGetName
func (e *Enum) MustGetName(what interface{}) string {
	if v, ok := e.GetName(what); ok {
		return v
	}
	panic("MustGetName")
}

// GetLower
func (e *Enum) GetLower(name string) (int32, bool) {
	return e.GetValue(strings.ToLower(name))
}

// MustGetLower
func (e *Enum) MustGetLower(name string) int32 {
	return e.MustGetValue(strings.ToLower(name))
}

// GetUpper
func (e *Enum) GetUpper(name string) (int32, bool) {
	return e.GetValue(strings.ToUpper(name))
}

// MustGetUpper
func (e *Enum) MustGetUpper(name string) int32 {
	return e.MustGetValue(strings.ToUpper(name))
}

// Register
func (e *Enum) Register(name string, value int32) bool {
	if e.Has(name) || e.Has(value) {
		return false
	}
	e.value2Name[value] = name
	e.name2Value[name] = value
	return true
}

// MustRegister
func (e *Enum) MustRegister(name string, value int32) bool {
	if e.Has(name) || e.Has(value) {
		panic("MustRegister")
	}
	return e.Register(name, value)
}

// CastRegister
func (e *Enum) CastRegister(name string, value interface{}) bool {
	if e.Has(name) || e.Has(value) {
		return false
	}
	cast := value.(int32)
	e.value2Name[cast] = name
	e.name2Value[name] = cast
	return true
}

// MustCastRegister
func (e *Enum) MustCastRegister(name string, value interface{}) bool {
	if e.Has(name) || e.Has(value) {
		panic("MustCastRegister")
	}
	return e.CastRegister(name, value)
}

// RegisterLower
func (e *Enum) RegisterLower(name string, value int32) bool {
	return e.Register(strings.ToLower(name), value)
}

// RegisterUpper
func (e *Enum) RegisterUpper(name string, value int32) bool {
	return e.Register(strings.ToUpper(name), value)
}

// Names
func (e *Enum) Names() []string {
	var names []string
	for name := range e.name2Value {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Values
func (e *Enum) Values() []int32 {
	var values []int
	for value := range e.value2Name {
		values = append(values, int(value))
	}
	sort.Ints(values)
	var values32 []int32
	for _, i := range values {
		values32 = append(values32, int32(i))
	}
	return values32
}
