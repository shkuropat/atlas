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

// KafkaEndpointConfig
type KafkaEndpointConfig interface {
	GetKafkaEndpoint() *atlas.KafkaEndpoint
	GetKafkaAddress() *atlas.KafkaAddress
	GetKafkaTopic() string
	GetKafkaGroupID() string
	GetKafkaReadNewest() bool
	GetKafkaAck() bool
}

// Interface compatibility
var _ KafkaEndpointConfig = ConfigKafka{}

// ConfigKafka
type ConfigKafka struct {
	Kafka *parts.ConfigKafka `mapstructure:"kafka"`
}

// ConfigKafkaNormalize
func (c ConfigKafka) ConfigKafkaNormalize() {
	if c.Kafka == nil {
		c.Kafka = parts.NewConfigKafka()
	}
}

// GetKafkaEndpoint
func (c ConfigKafka) GetKafkaEndpoint() *atlas.KafkaEndpoint {
	return atlas.NewKafkaEndpoint(c.Kafka.Brokers)
}

// GetKafkaAddress
func (c ConfigKafka) GetKafkaAddress() *atlas.KafkaAddress {
	return atlas.NewKafkaAddress(c.Kafka.Topic, 0)
}

// GetKafkaTopic
func (c ConfigKafka) GetKafkaTopic() string {
	return c.Kafka.Topic
}

// GetKafkaGroupID
func (c ConfigKafka) GetKafkaGroupID() string {
	return c.Kafka.GroupID
}

// GetKafkaReadNewest
func (c ConfigKafka) GetKafkaReadNewest() bool {
	return c.Kafka.ReadNewest
}

// GetKafkaAck
func (c ConfigKafka) GetKafkaAck() bool {
	return c.Kafka.Ack
}
