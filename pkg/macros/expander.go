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

package macros

import (
	"strings"
)

// Expander
type Expander struct {
	oldNew []string
	sep    string
}

// NewExpander
func NewExpander() *Expander {
	return new(Expander)
}

// SetSeparator
func (e *Expander) SetSeparator(sep string) *Expander {
	if e == nil {
		return nil
	}
	e.sep = sep
	return e
}

// Add
func (e *Expander) Add(key, value string) *Expander {
	if e == nil {
		return nil
	}
	e.oldNew = append(e.oldNew, key, value)
	return e
}

// Expand
func (e *Expander) Expand(lines ...string) string {
	if e == nil {
		return strings.Join(lines, " ")
	}
	return strings.Join(e.ExpandAll(lines...), e.sep)
}

// ExpandAll
func (e *Expander) ExpandAll(lines ...string) []string {
	if e == nil {
		return lines
	}
	var res []string
	for _, line := range lines {
		res = append(res, strings.NewReplacer(e.oldNew...).Replace(line))
	}
	return res
}
