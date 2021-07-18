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

// NewFilename
func NewFilename(filename ...string) *Filename {
	f := new(Filename)
	if len(filename) > 0 {
		f.Set(filename[0])
	}
	return f
}

// Ensure
func (m *Filename) Ensure() *Filename {
	if m != nil {
		return m
	}
	return new(Filename)
}

// Set
func (m *Filename) Set(filename string) *Filename {
	if m == nil {
		return nil
	}
	m.Filename = filename
	return m
}

// Equals
func (m *Filename) Equals(filename *Filename) bool {
	if m == nil {
		return false
	}
	if filename == nil {
		return false
	}
	return m.GetFilename() == filename.GetFilename()
}

// String
func (m *Filename) String() string {
	if m == nil {
		return ""
	}
	return m.GetFilename()
}
