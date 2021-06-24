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
type NamedCommands struct {
	Commands map[string]*Command `mapstructure:"commands"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewNamedCommands
func NewNamedCommands() *NamedCommands {
	return new(NamedCommands)
}

// GetLines
func (c *NamedCommands) GetCommand(name string) *Command {
	if c == nil {
		return nil
	}
	return c.Commands[name]
}

// String
func (c *NamedCommands) String() string {
	if c == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	for name, command := range c.Commands {
		_, _ = fmt.Fprintf(b, "%s:%s\n", name, command)
	}

	return b.String()
}
