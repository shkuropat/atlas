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

import "github.com/google/uuid"

// NewUUID
func NewUUID() *UUID {
	return &UUID{}
}

// CreateUUID
func CreateUUID() *UUID {
	return NewUUID().SetString(uuid.New().String())
}

// SetBytes
func (m *UUID) SetBytes(bytes []byte) *UUID {
	m.Data = bytes
	return m
}

// GetBytes
func (m *UUID) GetBytes() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// SetString
func (m *UUID) SetString(str string) *UUID {
	m.Data = []byte(str)
	return m
}

// GetString
func (m *UUID) GetString() string {
	if m != nil {
		return string(m.Data)
	}
	return ""
}
