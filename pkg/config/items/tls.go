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
type TLS struct {
	Enabled bool `mapstructure:"enabled"`

	// Client-side
	CAFile             string `mapstructure:"ca-file"`
	ServerNameOverride string `mapstructure:"server-name-override"`

	// Server-side
	PrivateKeyFile string `mapstructure:"private-key-file"`
	PublicCertFile string `mapstructure:"public-cert-file"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewTLS
func NewTLS() *TLS {
	return new(TLS)
}

// GetEnabled
func (t *TLS) GetEnabled() bool {
	if t == nil {
		return false
	}
	return t.Enabled
}

// GetCAFile
func (t *TLS) GetCAFile() string {
	if t == nil {
		return ""
	}
	return t.CAFile
}

// GetServerNameOverride
func (t *TLS) GetServerNameOverride() string {
	if t == nil {
		return ""
	}
	return t.ServerNameOverride
}

// GetPrivateKeyFile
func (t *TLS) GetPrivateKeyFile() string {
	if t == nil {
		return ""
	}
	return t.PrivateKeyFile
}

// GetPublicCertFile
func (t *TLS) GetPublicCertFile() string {
	if t == nil {
		return ""
	}
	return t.PublicCertFile
}

// String
func (t *TLS) String() string {
	if t == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", t.Enabled)
	_, _ = fmt.Fprintf(b, "CAFile: %v\n", t.CAFile)
	_, _ = fmt.Fprintf(b, "ServerNameOverride: %v\n", t.ServerNameOverride)
	_, _ = fmt.Fprintf(b, "PrivateKeyFile: %v\n", t.PrivateKeyFile)
	_, _ = fmt.Fprintf(b, "PublicCertFile: %v\n", t.PublicCertFile)

	return b.String()
}
