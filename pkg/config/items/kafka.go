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
type Kafka struct {
	Enabled bool `mapstructure:"enabled"`
	// Brokers specifies list of Kafka brokers to connect to. Used by server and client
	Brokers []string `mapstructure:"brokers"`
	// Topic specifies topic to read from or write into in Kafka. Used by server and client
	Topic string `mapstructure:"topic"`
	// GroupID specifies consumer group id. Used by client only
	GroupID string `mapstructure:"groupID"`
	// ReadNewest specifies whether to read newest messages. Used by client only
	ReadNewest bool `mapstructure:"readNewest"`
	// Ack specifies whether to ack messages. Used by client only
	Ack bool `mapstructure:"ack"`
	//
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewKafka
func NewKafka() *Kafka {
	return new(Kafka)
}

// GetBrokers
func (k *Kafka) GetBrokers() []string {
	if k == nil {
		return nil
	}
	return k.Brokers
}

// GetTopic
func (k *Kafka) GetTopic() string {
	if k == nil {
		return ""
	}
	return k.Topic
}

// GetGroupID
func (k *Kafka) GetGroupID() string {
	if k == nil {
		return ""
	}
	return k.GroupID
}

// GetReadNewest
func (k *Kafka) GetReadNewest() bool {
	if k == nil {
		return false
	}
	return k.ReadNewest
}

// GetAck
func (k *Kafka) GetAck() bool {
	if k == nil {
		return false
	}
	return k.Ack
}

// String
func (k *Kafka) String() string {
	if k == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", k.Enabled)
	_, _ = fmt.Fprintf(b, "Brokers: %v\n", k.Brokers)
	_, _ = fmt.Fprintf(b, "Topic: %v\n", k.Topic)
	_, _ = fmt.Fprintf(b, "GroupID: %v\n", k.GroupID)
	_, _ = fmt.Fprintf(b, "ReadNewest: %v\n", k.ReadNewest)
	_, _ = fmt.Fprintf(b, "Ack: %v\n", k.Ack)

	return b.String()
}
