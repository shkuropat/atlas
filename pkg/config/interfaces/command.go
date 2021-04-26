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

package interfaces

import (
	"github.com/binarly-io/atlas/pkg/config/parts"
	"github.com/binarly-io/atlas/pkg/macros"
	"strings"
)

// CommandConfigurator
type CommandConfigurator interface {
	GetArgs() []string
	GetCommand() string
	ParseArgs(*macros.Expander) []string
	ParseCommand(*macros.Expander) string
}

// Interface compatibility
var _ CommandConfigurator = CommandConfig{}

// CommandConfig
type CommandConfig struct {
	Command *parts.CommandConfig `mapstructure:"command"`
}

// CommandConfigNormalize
func (c CommandConfig) CommandConfigNormalize() {
	if c.Command == nil {
		c.Command = parts.NewCommandConfig()
	}
}

// GetArgs
func (c CommandConfig) GetArgs() []string {
	return c.Command.Args
}

// GetCommand
func (c CommandConfig) GetCommand() string {
	return strings.Join(c.GetArgs(), " ")
}

// ParseArgs
func (c CommandConfig) ParseArgs(macro *macros.Expander) []string {
	return macro.ExpandAll(c.Command.Args...)
}

// ParseCommand
func (c CommandConfig) ParseCommand(macro *macros.Expander) string {
	return strings.Join(c.ParseArgs(macro), " ")
}
