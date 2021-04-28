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

// NewDirname
func NewDirname(dirname ...string) *Dirname {
	f := new(Dirname)
	if len(dirname) > 0 {
		f.Set(dirname[0])
	}
	return f
}

// Set
func (m *Dirname) Set(dirname string) *Dirname {
	if m == nil {
		return nil
	}
	m.Dirname = dirname
	return m
}

// Equals
func (m *Dirname) Equals(dirname *Dirname) bool {
	if m == nil {
		return false
	}
	if dirname == nil {
		return false
	}
	return m.GetDirname() == dirname.GetDirname()
}

// String
func (m *Dirname) String() string {
	if m == nil {
		return ""
	}
	return m.GetDirname()
}
