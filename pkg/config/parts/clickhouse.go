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
type ConfigClickHouse struct {
	Enabled bool `mapstructure:"enabled"`
	// DSN in the form: http://clickhouse_operator:clickhouse_operator_password@clickhouse-journal.clickhouse:8123/atlas
	DSN string `mapstructure:"dsn"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewConfigClickHouse
func NewConfigClickHouse() *ConfigClickHouse {
	return new(ConfigClickHouse)
}

// String
func (c *ConfigClickHouse) String() string {
	if c == nil {
		return ""
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", c.Enabled)
	_, _ = fmt.Fprintf(b, "DSN: %v\n", c.DSN)

	return b.String()
}
