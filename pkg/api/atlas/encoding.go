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

// NewEncoding
func NewEncoding(method ...string) *Encoding {
	f := new(Encoding)
	if len(method) > 0 {
		f.Set(method[0])
	}
	return f
}

// Set
func (m *Encoding) Set(method string) *Encoding {
	if m == nil {
		return nil
	}
	m.Method = method
	return m
}

// Equals
func (m *Encoding) Equals(encoding *Encoding) bool {
	if m == nil {
		return false
	}
	if encoding == nil {
		return false
	}
	return m.GetMethod() == encoding.GetMethod()
}

// String
func (m *Encoding) String() string {
	if m == nil {
		return ""
	}
	return m.GetMethod()
}
