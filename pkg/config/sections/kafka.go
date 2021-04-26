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

package sections

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config/items"
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
var _ KafkaConfigurator = Kafka{}

// Kafka
type Kafka struct {
	Kafka *items.Kafka `mapstructure:"kafka"`
}

// KafkaNormalize
func (c Kafka) KafkaNormalize() {
	if c.Kafka == nil {
		c.Kafka = items.NewKafka()
	}
}

// GetKafkaEndpoint
func (c Kafka) GetKafkaEndpoint() *atlas.KafkaEndpoint {
	return atlas.NewKafkaEndpoint(c.Kafka.GetBrokers())
}

// GetKafkaAddress
func (c Kafka) GetKafkaAddress() *atlas.KafkaAddress {
	return atlas.NewKafkaAddress(c.Kafka.GetTopic(), 0)
}

// GetKafkaTopic
func (c Kafka) GetKafkaTopic() string {
	return c.Kafka.GetTopic()
}

// GetKafkaGroupID
func (c Kafka) GetKafkaGroupID() string {
	return c.Kafka.GetGroupID()
}

// GetKafkaReadNewest
func (c Kafka) GetKafkaReadNewest() bool {
	return c.Kafka.GetReadNewest()
}

// GetKafkaAck
func (c Kafka) GetKafkaAck() bool {
	return c.Kafka.GetAck()
}
