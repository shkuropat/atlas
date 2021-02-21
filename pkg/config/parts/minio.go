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
type ConfigMinIO struct {
	Enabled bool `mapstructure:"enabled"`
	// MinIO connection
	Endpoint           string `mapstructure:"endpoint"`
	AccessKeyID        string `mapstructure:"accessKeyID"`
	SecretAccessKey    string `mapstructure:"secretAccessKey"`
	Secure             bool   `mapstructure:"secure"`
	InsecureSkipVerify bool   `mapstructure:"insecureSkipVerify"`
	// MinIO internals
	Bucket string `mapstructure:"bucket"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

func NewConfigMinIO() *ConfigMinIO {
	return new(ConfigMinIO)
}

func (c *ConfigMinIO) String() string {
	if c == nil {
		return ""
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", c.Enabled)
	_, _ = fmt.Fprintf(b, "Endpoint: %v\n", c.Endpoint)
	_, _ = fmt.Fprintf(b, "AccessKeyID: %v\n", c.AccessKeyID)
	_, _ = fmt.Fprintf(b, "SecretAccessKey: %v\n", c.SecretAccessKey)
	_, _ = fmt.Fprintf(b, "Secure: %v\n", c.Secure)
	_, _ = fmt.Fprintf(b, "InsecureSkipVerify: %v\n", c.InsecureSkipVerify)
	_, _ = fmt.Fprintf(b, "Bucket: %v\n", c.Bucket)

	return b.String()
}
