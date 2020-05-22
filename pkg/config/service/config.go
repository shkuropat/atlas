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

package config_service

import (
	"bytes"
	"fmt"

	conf "github.com/spf13/viper"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type ConfigService struct {
	Verbose bool `mapstructure:"verbose"`

	// Service
	ServiceAddress string `mapstructure:"service-address"`

	// TLS
	TLS         bool   `mapstructure:"tls"`
	TLSCertFile string `mapstructure:"tls-cert-file"`
	TLSKeyFile  string `mapstructure:"tls-key-file"`

	// OAuth
	OAuth            bool   `mapstructure:"oauth"`
	JWTPublicKeyFile string `mapstructure:"jwt-public-key-file"`

	// Kafka
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`

	// MinIO
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	Secure          bool   `mapstructure:"secure"`

	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

var Config ConfigService

func ReadIn() {
	_ = conf.Unmarshal(&Config)
}

func (c *ConfigService) String() string {
	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Verbose: %v\n", c.Verbose)

	_, _ = fmt.Fprintf(b, "ServiceAddress: %v\n", c.ServiceAddress)

	_, _ = fmt.Fprintf(b, "TLS: %v\n", c.TLS)
	_, _ = fmt.Fprintf(b, "TLSCertFile: %v\n", c.TLSCertFile)
	_, _ = fmt.Fprintf(b, "TLSKeyFile: %v\n", c.TLSKeyFile)

	_, _ = fmt.Fprintf(b, "OAuth: %v\n", c.OAuth)
	_, _ = fmt.Fprintf(b, "JWTPublicKeyFile: %v\n", c.JWTPublicKeyFile)

	_, _ = fmt.Fprintf(b, "Brokers: %v\n", c.Brokers)
	_, _ = fmt.Fprintf(b, "Topic: %v\n", c.Topic)

	_, _ = fmt.Fprintf(b, "Endpoint: %v\n", c.Endpoint)
	_, _ = fmt.Fprintf(b, "AccessKeyID: %v\n", c.AccessKeyID)
	_, _ = fmt.Fprintf(b, "SecretAccessKey: %v\n", c.SecretAccessKey)
	_, _ = fmt.Fprintf(b, "Secure: %v\n", c.Secure)

	return b.String()
}
