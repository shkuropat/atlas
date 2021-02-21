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
type ConfigTLS struct {
	Enabled            bool   `mapstructure:"enabled"`
	CAFile             string `mapstructure:"ca-file"`
	ServerHostOverride string `mapstructure:"server-host-override"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

func NewConfigTLS() *ConfigTLS {
	return new(ConfigTLS)
}

func (c *ConfigTLS) String() string {
	if c == nil {
		return ""
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", c.Enabled)
	_, _ = fmt.Fprintf(b, "CAFile: %v\n", c.CAFile)
	_, _ = fmt.Fprintf(b, "ServerHostOverride: %v\n", c.ServerHostOverride)

	return b.String()
}
