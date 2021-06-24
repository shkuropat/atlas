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

package sections

import (
	"fmt"
	"strings"

	"github.com/binarly-io/atlas/pkg/config/items"
	"github.com/binarly-io/atlas/pkg/macros"
)

// NamedCommandsConfigurator
type NamedCommandsConfigurator interface {
	GetCommandLines(name string) []string
	GetCommand(name string) string
	ParseCommandLines(name string, macro *macros.Expander) []string
	ParseCommand(name string, macro *macros.Expander) string
}

// Interface compatibility
var _ NamedCommandsConfigurator = NamedCommands{}

// NamedCommands
type NamedCommands struct {
	Commands *items.NamedCommands `mapstructure:"commands"`
}

// NamedCommandsNormalize
func (c NamedCommands) NamedCommandsNormalize() NamedCommands {
	if c.Commands == nil {
		c.Commands = items.NewNamedCommands()
	}
	return c
}

// GetCommandLines
func (c NamedCommands) GetCommandLines(name string) []string {
	return c.Commands.GetCommand(name).GetCommand()
}

// GetCommand
func (c NamedCommands) GetCommand(name string) string {
	return strings.Join(c.GetCommandLines(name), " ")
}

// ParseCommandLines
func (c NamedCommands) ParseCommandLines(name string, macro *macros.Expander) []string {
	return macro.ExpandAll(c.GetCommandLines(name)...)
}

// ParseCommand
func (c NamedCommands) ParseCommand(name string, macro *macros.Expander) string {
	return strings.Join(c.ParseCommandLines(name, macro), " ")
}

// String
func (c NamedCommands) String() string {
	return fmt.Sprintf("Commands=%s", c.Commands)
}
