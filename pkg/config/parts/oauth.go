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
type ConfigOAuth struct {
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

func NewConfigOAuth() *ConfigOAuth {
	return new(ConfigOAuth)
}

func (c *ConfigOAuth) String() string {
	if c == nil {
		return ""
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", c.Enabled)
	_, _ = fmt.Fprintf(b, "ClientID: %v\n", c.ClientID)
	_, _ = fmt.Fprintf(b, "ClientSecret: %v\n", c.ClientSecret)
	_, _ = fmt.Fprintf(b, "TokenURL: %v\n", c.TokenURL)
	_, _ = fmt.Fprintf(b, "RegisterURL: %v\n", c.RegisterURL)
	_, _ = fmt.Fprintf(b, "InitialAccessToken: %v\n", c.InitialAccessToken)

	return b.String()
}
