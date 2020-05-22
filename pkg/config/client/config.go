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

package config_client

import (
	"bytes"
	"fmt"

	conf "github.com/spf13/viper"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type ConfigClient struct {
	Verbose bool `mapstructure:"verbose"`

	// Service
	ServiceAddress string `mapstructure:"service-address"`

	// TLS
	TLS                bool   `mapstructure:"tls"`
	CAFile             string `mapstructure:"ca-file"`
	ServerHostOverride string `mapstructure:"server-host-override"`

	// OAuth
	OAuth bool `mapstructure:"oauth"`
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

var Config ConfigClient

func ReadIn() {
	_ = conf.Unmarshal(&Config)
}

func (c *ConfigClient) String() string {
	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Verbose: %v\n", c.Verbose)

	_, _ = fmt.Fprintf(b, "ServiceAddress: %v\n", c.ServiceAddress)

	// TLS
	_, _ = fmt.Fprintf(b, "TLS: %v\n", c.TLS)
	_, _ = fmt.Fprintf(b, "CAFile: %v\n", c.CAFile)
	_, _ = fmt.Fprintf(b, "ServerHostOverride: %v\n", c.ServerHostOverride)

	// OAuth
	_, _ = fmt.Fprintf(b, "OAuth: %v\n", c.OAuth)
	// OAuth Login
	_, _ = fmt.Fprintf(b, "ClientID: %v\n", c.ClientID)
	_, _ = fmt.Fprintf(b, "ClientSecret: %v\n", c.ClientSecret)
	_, _ = fmt.Fprintf(b, "TokenURL: %v\n", c.TokenURL)
	// OAuth Register
	_, _ = fmt.Fprintf(b, "RegisterURL: %v\n", c.RegisterURL)
	_, _ = fmt.Fprintf(b, "InitialAccessToken: %v\n", c.InitialAccessToken)

	return b.String()
}
