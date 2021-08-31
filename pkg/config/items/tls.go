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
	Enabled            bool   `mapstructure:"enabled"`
	CAFile             string `mapstructure:"ca-file"`
	KeyFile            string `mapstructure:"key-file"`
	ServerHostOverride string `mapstructure:"server-host-override"`
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

// GetKeyFile
func (t *TLS) GetKeyFile() string {
	if t == nil {
		return ""
	}
	return t.KeyFile
}

// GetServerHostOverride
func (t *TLS) GetServerHostOverride() string {
	if t == nil {
		return ""
	}
	return t.ServerHostOverride
}

// String
func (t *TLS) String() string {
	if t == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", t.Enabled)
	_, _ = fmt.Fprintf(b, "CAFile: %v\n", t.CAFile)
	_, _ = fmt.Fprintf(b, "KeyFile: %v\n", t.KeyFile)
	_, _ = fmt.Fprintf(b, "ServerHostOverride: %v\n", t.ServerHostOverride)

	return b.String()
}
