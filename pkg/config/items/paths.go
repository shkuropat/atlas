// Copyright 2021 The Atlas Authors. All rights reserved.
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

package items

import (
	"bytes"
	"fmt"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type Paths struct {
	Map map[string][]string `mapstructure:"map"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewPaths
func NewPaths() *Paths {
	return &Paths{
		Map: make(map[string][]string),
	}
}

// GetMap
func (f *Paths) GetMap() map[string][]string {
	if f == nil {
		return nil
	}
	return f.Map
}

// GetPaths
func (f *Paths) GetPaths(name string) []string {
	if f == nil {
		return nil
	}
	if paths, ok := f.Map[name]; ok {
		return paths
	}
	return nil
}

// GetOne
func (f *Paths) GetOne(name string) string {
	if paths := f.GetPaths(name); len(paths) > 0 {
		return paths[0]
	}
	return ""
}

// String
func (f *Paths) String() string {
	if f == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Map: %v\n", f.Map)

	return b.String()
}
