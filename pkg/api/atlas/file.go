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

import "fmt"

// NewFile
func NewFile() *File {
	return new(File)
}

// Len
func (m *File) Len() int {
	if m != nil {
		return len(m.Data)
	}
	return 0
}

// SetFilename
func (m *File) SetFilename(filename string) *File {
	if m == nil {
		return nil
	}
	m.Filename = m.Filename.Ensure().Set(filename)
	return m
}

// SetData
func (m *File) SetData(data []byte) *File {
	if m == nil {
		return nil
	}
	m.Data = data
	return m
}

// String
func (m *File) String() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("%s[%d]", m.Filename.String(), m.Len())
}
