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
	"strings"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type Command struct {
	Workdir string   `mapstructure:"workdir"`
	Env     []string `mapstructure:"env"`
	Command []string `mapstructure:"command"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewCommand
func NewCommand() *Command {
	return new(Command)
}

// GetWorkdir
func (c *Command) GetWorkdir() string {
	if c == nil {
		return ""
	}
	return c.Workdir
}

// GetEnv
func (c *Command) GetEnv() []string {
	if c == nil {
		return nil
	}
	return c.Env
}

// GetCommand
func (c *Command) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

// String
func (c *Command) String() string {
	if c == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, strings.Join(c.Command, " "))

	return b.String()
}
