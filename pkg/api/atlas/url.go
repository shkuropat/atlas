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

// NewURL
func NewURL(url ...string) *URL {
	f := new(URL)
	if len(url) > 0 {
		f.Set(url[0])
	}
	return f
}

// Set
func (m *URL) Set(url string) *URL {
	if m == nil {
		return nil
	}
	m.Url = url
	return m
}

// Equals
func (m *URL) Equals(url *URL) bool {
	if m == nil {
		return false
	}
	if url == nil {
		return false
	}
	return m.GetUrl() == url.GetUrl()
}

// String
func (m *URL) String() string {
	if m == nil {
		return ""
	}
	return m.GetUrl()
}
