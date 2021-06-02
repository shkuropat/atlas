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

package items

import (
	"bytes"
	"fmt"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type OAuth struct {
	Enabled bool `mapstructure:"enabled"`
	// OAuth Login
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
	TokenURL     string `mapstructure:"token-url"`
	// OAuth Register
	RegisterURL        string `mapstructure:"register-url"`
	InitialAccessToken string `mapstructure:"initial-access-token"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewOAuth
func NewOAuth() *OAuth {
	return new(OAuth)
}

// GetEnabled
func (o *OAuth) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

// GetClientID
func (o *OAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

// GetClientSecret
func (o *OAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

// GetTokenURL
func (o *OAuth) GetTokenURL() string {
	if o == nil {
		return ""
	}
	return o.TokenURL
}

// GetRegisterURL
func (o *OAuth) GetRegisterURL() string {
	if o == nil {
		return ""
	}
	return o.RegisterURL
}

// GetInitialAccessToken
func (o *OAuth) GetInitialAccessToken() string {
	if o == nil {
		return ""
	}
	return o.InitialAccessToken
}

// String
func (o *OAuth) String() string {
	if o == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", o.Enabled)
	_, _ = fmt.Fprintf(b, "ClientID: %v\n", o.ClientID)
	_, _ = fmt.Fprintf(b, "ClientSecret: %v\n", o.ClientSecret)
	_, _ = fmt.Fprintf(b, "TokenURL: %v\n", o.TokenURL)
	_, _ = fmt.Fprintf(b, "RegisterURL: %v\n", o.RegisterURL)
	_, _ = fmt.Fprintf(b, "InitialAccessToken: %v\n", o.InitialAccessToken)

	return b.String()
}
