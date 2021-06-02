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

// OAuthConfigurator
type OAuthConfigurator interface {
	GetOAuthEnabled() bool
	GetOAuthClientID() string
	GetOAuthClientSecret() string
	GetOAuthTokenURL() string
	GetOAuthRegisterURL() string
	GetOAuthInitialAccessToken() string
}

// Interface compatibility
var _ OAuthConfigurator = OAuth{}

// OAuth
type OAuth struct {
	OAuth *items.OAuth `mapstructure:"oauth"`
}

// OAuthNormalize
func (c OAuth) OAuthNormalize() OAuth {
	if c.OAuth == nil {
		c.OAuth = items.NewOAuth()
	}
	return c
}

// GetOAuthEnabled
func (c OAuth) GetOAuthEnabled() bool {
	return c.OAuth.GetEnabled()
}

// GetOAuthClientID
func (c OAuth) GetOAuthClientID() string {
	return c.OAuth.GetClientID()
}

// GetOAuthClientSecret
func (c OAuth) GetOAuthClientSecret() string {
	return c.OAuth.GetClientSecret()
}

// GetOAuthTokenURL
func (c OAuth) GetOAuthTokenURL() string {
	return c.OAuth.GetTokenURL()
}

// GetOAuthRegisterURL
func (c OAuth) GetOAuthRegisterURL() string {
	return c.OAuth.GetRegisterURL()
}

// GetOAuthInitialAccessToken
func (c OAuth) GetOAuthInitialAccessToken() string {
	return c.OAuth.GetInitialAccessToken()
}

// String
func (c OAuth) String() string {
	return fmt.Sprintf("OAuth=%s", c.OAuth)
}
