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
	"github.com/binarly-io/atlas/pkg/config/items"
)

// LogConfigurator
type LogConfigurator interface {
	GetLogLevel() string
	GetLogFormatter() string
}

// Interface compatibility
var _ LogConfigurator = Log{}

// Log
type Log struct {
	Log *items.Log `mapstructure:"log"`
}

// LogNormalize
func (c Log) LogNormalize() {
	if c.Log == nil {
		c.Log = items.NewLog()
	}
}

// GetLogLevel
func (c Log) GetLogLevel() string {
	return c.Log.GetLevel()
}

// GetLogFormatter
func (c Log) GetLogFormatter() string {
	return c.Log.GetFormatter()
}

// String
func (c Log) String() string {
	return fmt.Sprintf("Log=%s", c.Log)
}
