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

// TLSConfigurator
type TLSConfigurator interface {
	GetEnabled() bool
	GetCAFile() string
	GetServerHostOverride() string
}

// Interface compatibility
var _ TLSConfigurator = TLSConfig{}

// TLSConfig
type TLSConfig struct {
	TLS *parts.TLSConfig `mapstructure:"tls"`
}

// ConfigTLSNormalize
func (c TLSConfig) ConfigTLSNormalize() {
	if c.TLS == nil {
		c.TLS = parts.NewTLSConfig()
	}
}

// GetEnabled
func (c TLSConfig) GetEnabled() bool {
	return c.TLS.Enabled
}

// GetCAFile
func (c TLSConfig) GetCAFile() string {
	return c.TLS.CAFile
}

// GetServerHostOverride
func (c TLSConfig) GetServerHostOverride() string {
	return c.TLS.ServerHostOverride
}
