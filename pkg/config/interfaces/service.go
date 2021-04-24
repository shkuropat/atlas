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
)

// OAuthConfigurator
type ServiceConfigurator interface {
	GetServiceAddress() string
}

// Interface compatibility
var _ ServiceConfigurator = ServiceConfig{}

// ServiceConfig
type ServiceConfig struct {
	Service *parts.ServiceConfig `mapstructure:"service"`
}

// ServiceConfigNormalize
func (c ServiceConfig) ServiceConfigNormalize() {
	if c.Service == nil {
		c.Service = parts.NewServiceConfig()
	}
}

// GetServiceAddress
func (c ServiceConfig) GetServiceAddress() string {
	return c.Service.ServiceAddress
}
