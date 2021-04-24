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

package interfaces

import "github.com/binarly-io/atlas/pkg/config/parts"

// MinIOConfigurator
type MinIOConfigurator interface {
	GetMinIOEndpoint() string
	GetMinIOAccessKeyID() string
	GetMinIOSecretAccessKey() string
	GetMinIOSecure() bool
	GetMinIOInsecureSkipVerify() bool
	GetMinIOBucket() string
}

// Interface compatibility
var _ MinIOConfigurator = MinIOConfig{}

// MinIOConfig
type MinIOConfig struct {
	MinIO *parts.MinIOConfig `mapstructure:"minio"`
}

// MinIOConfigNormalize
func (c MinIOConfig) MinIOConfigNormalize() {
	if c.MinIO == nil {
		c.MinIO = parts.NewMinIOConfig()
	}
}

// GetMinIOEndpoint
func (c MinIOConfig) GetMinIOEndpoint() string {
	return c.MinIO.Endpoint
}

// GetMinIOAccessKeyID
func (c MinIOConfig) GetMinIOAccessKeyID() string {
	return c.MinIO.AccessKeyID
}

// GetMinIOSecretAccessKey
func (c MinIOConfig) GetMinIOSecretAccessKey() string {
	return c.MinIO.SecretAccessKey
}

// GetMinIOSecure
func (c MinIOConfig) GetMinIOSecure() bool {
	return c.MinIO.Secure
}

// GetMinIOInsecureSkipVerify
func (c MinIOConfig) GetMinIOInsecureSkipVerify() bool {
	return c.MinIO.InsecureSkipVerify
}

// GetMinIOBucket
func (c MinIOConfig) GetMinIOBucket() string {
	return c.MinIO.Bucket
}
