// Copyright 2020 The Atlas Authors. All rights reserved.
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

import "github.com/binarly-io/atlas/pkg/config/parts"

// ClickHouseEndpointConfig
type ClickHouseEndpointConfig interface {
	GetClickHouseEndpoint() string
}

// Interface compatibility
var _ ClickHouseEndpointConfig = ConfigClickHouse{}

// ConfigClickHouse
type ConfigClickHouse struct {
	ClickHouse *parts.ConfigClickHouse `mapstructure:"clickhouse"`
}

// ConfigClickHouseNormalize
func (c ConfigClickHouse) ConfigClickHouseNormalize() {
	if c.ClickHouse == nil {
		c.ClickHouse = parts.NewConfigClickHouse()
	}
}

// GetClickHouseEndpoint
func (c ConfigClickHouse) GetClickHouseEndpoint() string {
	return c.ClickHouse.DSN
}
