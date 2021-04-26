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
	"context"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config/sections"
	"github.com/binarly-io/atlas/pkg/softwareid"
)

// ConsumeMessageFunction specifies message consumeMessageFunction function
type ConsumeMessageFunction func(context.Context, *sarama.ConsumerMessage) bool

// ConsumerGroup
type ConsumerGroup struct {
	// endpoint specifies Kafka endpoint which to connect to
	endpoint *atlas.KafkaEndpoint
	// address specifies Kafka address which to connect to
	address *atlas.KafkaAddress
	// groupID specifies id of the consumer group
	groupID string

	consumerGroupHandler   sarama.ConsumerGroupHandler
	ctx                    context.Context
	consumeMessageFunction ConsumeMessageFunction
}

// NewConsumerGroup creates new consumer group
func NewConsumerGroup(endpoint *atlas.KafkaEndpoint, address *atlas.KafkaAddress, groupID string) *ConsumerGroup {
	return &ConsumerGroup{
		endpoint: endpoint,
		address:  address,
		groupID:  groupID,
	}
}

// NewConsumerGroupFromEndpoint.
// IMPORTANT - you have to specify topic to read from either with
//	1. SetAddress
//	2. SetTopic
func NewConsumerGroupFromEndpoint(cfg sections.KafkaConfigurator, groupID string) *ConsumerGroup {
	return NewConsumerGroup(cfg.GetKafkaEndpoint(), nil, groupID)
}

// SetAddress - sets the full address - Topic and Partition
func (c *ConsumerGroup) SetAddress(address *atlas.KafkaAddress) *ConsumerGroup {
	c.address = address
	return c
}

// SetTopic - sets address in simplified form - specified Topic and Partition 0
func (c *ConsumerGroup) SetTopic(topic string) *ConsumerGroup {
	c.address = atlas.NewKafkaAddress(topic, 0)
	return c
}

// SetContext - sets context to be used by MessageProcessor
func (c *ConsumerGroup) SetContext(ctx context.Context) *ConsumerGroup {
	c.ctx = ctx
	return c
}

// SetConsumerGroupHandler sets handler which performs setup, cleanup and message processing activities
func (c *ConsumerGroup) SetConsumerGroupHandler(handler sarama.ConsumerGroupHandler) *ConsumerGroup {
	c.consumerGroupHandler = handler
	return c
}

// SetConsumeMessageFunction sets function which will be called for each message received from Kafka
func (c *ConsumerGroup) SetConsumeMessageFunction(consumeMessageFunction ConsumeMessageFunction) *ConsumerGroup {
	c.consumeMessageFunction = consumeMessageFunction
	return c
}

// ConsumeLoop runs an endless loop of kafka consumer
func (c *ConsumerGroup) ConsumeLoop(consumeNewest bool, ack bool) {
	log.Info("ConsumerGroup.ConsumeLoop() - start")
	defer log.Info("ConsumerGroup.ConsumeLoop() - end")

	// New configuration instance with sane defaults.
	config := sarama.NewConfig()
	// Consumer groups require Version to be >= V0_10_2_0
	config.Version = sarama.V2_0_0_0
	config.ClientID = softwareid.Name
	if consumeNewest {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	group, err := sarama.NewConsumerGroup(c.endpoint.Brokers, c.groupID, config)
	if err != nil {
		log.Fatalf("unable to create NewConsumerGroup for %v %v err: %v", c.endpoint.Brokers, c.groupID, err)
	}
	defer func() {
		_ = group.Close()
	}()

	// Track errors
	//go func() {
	//	for err := range group.Errors() {
	//		fmt.Println("ERROR", err)
	//	}
	//}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		// Handler can be either explicitly specified, or a default one
		// Default handler can still use external c.consumeMessageFunction
		handler := c.consumerGroupHandler
		if handler == nil {
			handler = newConsumerGroupHandler(c.ctx, c.consumeMessageFunction, ack)
		}

		// Consume joins a cluster of consumers for a given list of topics
		//
		// `Consume` should be called inside an infinite loop.
		// When a server-side rebalance happens, the consumer session will need to be recreated to get the new claims
		err := group.Consume(ctx, c.address.GetTopics(), handler)
		if err != nil {
			log.Fatalf("unable to Consume topics %v err: %v", c.address.GetTopics(), err)
		}
	}
}
