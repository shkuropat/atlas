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

package kafka

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/softwareid"
)

// Consumer
type Consumer struct {
	Endpoint

	config            *sarama.Config
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
}

// NewConsumer
func NewConsumer(endpoint Endpoint) *Consumer {
	var err error

	c := &Consumer{}
	c.Endpoint = endpoint
	c.config = sarama.NewConfig()
	c.config.ClientID = softwareid.Name
	c.consumer, err = sarama.NewConsumer(c.Brokers, c.config)
	if err != nil {
		c.Close()
		return nil
	}
	c.partitionConsumer, err = c.consumer.ConsumePartition(c.Topic, c.Partition, sarama.OffsetNewest)
	if err != nil {
		c.Close()
		return nil
	}

	return c
}

// Close
func (c *Consumer) Close() {
	if c.partitionConsumer != nil {
		_ = c.partitionConsumer.Close()
		c.partitionConsumer = nil
	}

	if c.consumer != nil {
		_ = c.consumer.Close()
		c.consumer = nil
	}
}

// Recv
func (c *Consumer) Recv() *sarama.ConsumerMessage {
	msg := <-c.partitionConsumer.Messages()
	if msg != nil {
		log.Printf("Got message topic:%q partition:%d offset:%d data:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
	}
	return msg
}
