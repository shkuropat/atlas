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

package parts

import (
	"bytes"
	"fmt"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type ConfigLog struct {
	Level     string `mapstructure:"level"`
	Formatter string `mapstructure:"format"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

func NewConfigLog() *ConfigLog {
	return new(ConfigLog)
}

func (c *ConfigLog) String() string {
	if c == nil {
		return ""
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Level: %v\n", c.Level)
	_, _ = fmt.Fprintf(b, "Formatter: %v\n", c.Formatter)

	return b.String()
}
