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

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config/parts"
)

// KafkaConfigurator
type KafkaConfigurator interface {
	GetKafkaEndpoint() *atlas.KafkaEndpoint
	GetKafkaAddress() *atlas.KafkaAddress
	GetKafkaTopic() string
	GetKafkaGroupID() string
	GetKafkaReadNewest() bool
	GetKafkaAck() bool
}

// Interface compatibility
var _ KafkaConfigurator = KafkaConfig{}

// KafkaConfig
type KafkaConfig struct {
	Kafka *parts.KafkaConfig `mapstructure:"kafka"`
}

// KafkaConfigNormalize
func (c KafkaConfig) KafkaConfigNormalize() {
	if c.Kafka == nil {
		c.Kafka = parts.NewKafkaConfig()
	}
}

// GetKafkaEndpoint
func (c KafkaConfig) GetKafkaEndpoint() *atlas.KafkaEndpoint {
	return atlas.NewKafkaEndpoint(c.Kafka.Brokers)
}

// GetKafkaAddress
func (c KafkaConfig) GetKafkaAddress() *atlas.KafkaAddress {
	return atlas.NewKafkaAddress(c.Kafka.Topic, 0)
}

// GetKafkaTopic
func (c KafkaConfig) GetKafkaTopic() string {
	return c.Kafka.Topic
}

// GetKafkaGroupID
func (c KafkaConfig) GetKafkaGroupID() string {
	return c.Kafka.GroupID
}

// GetKafkaReadNewest
func (c KafkaConfig) GetKafkaReadNewest() bool {
	return c.Kafka.ReadNewest
}

// GetKafkaAck
func (c KafkaConfig) GetKafkaAck() bool {
	return c.Kafka.Ack
}
