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
type OAuthConfigurator interface {
	GetEnabled() bool
	GetClientID() string
	GetClientSecret() string
	GetTokenURL() string
	GetRegisterURL() string
	GetInitialAccessToken() string
}

// Interface compatibility
var _ OAuthConfigurator = OAuthConfig{}

// OAuthConfig
type OAuthConfig struct {
	OAuth *parts.OAuthConfig `mapstructure:"oauth"`
}

// OAuthConfigNormalize
func (c OAuthConfig) OAuthConfigNormalize() {
	if c.OAuth == nil {
		c.OAuth = parts.NewOAuthConfig()
	}
}

// GetEnabled
func (c OAuthConfig) GetEnabled() bool {
	return c.OAuth.Enabled
}

// GetClientID
func (c OAuthConfig) GetClientID() string {
	return c.OAuth.ClientID
}

// GetClientSecret
func (c OAuthConfig) GetClientSecret() string {
	return c.OAuth.ClientSecret
}

// GetTokenURL
func (c OAuthConfig) GetTokenURL() string {
	return c.OAuth.TokenURL
}

// GetRegisterURL
func (c OAuthConfig) GetRegisterURL() string {
	return c.OAuth.RegisterURL
}

// GetInitialAccessToken
func (c OAuthConfig) GetInitialAccessToken() string {
	return c.OAuth.InitialAccessToken
}
