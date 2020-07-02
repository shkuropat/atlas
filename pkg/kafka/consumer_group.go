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
	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/softwareid"
)

// ConsumerGroup
type ConsumerGroup struct {
	endpoint *atlas.KafkaEndpoint
	address  *atlas.KafkaAddress
	groupID  string

	consumerGroupHandler sarama.ConsumerGroupHandler
	messageProcessor     func(*sarama.ConsumerMessage) bool
}

// NewConsumerGroup creates new consumer group
func NewConsumerGroup(endpoint *atlas.KafkaEndpoint, address *atlas.KafkaAddress, groupID string) *ConsumerGroup {
	return &ConsumerGroup{
		endpoint: endpoint,
		address:  address,
		groupID:  groupID,
	}
}

// NewConsumerGroupConfig
func NewConsumerGroupConfig(cfg config.KafkaEndpointConfig, groupID string) *ConsumerGroup {
	return NewConsumerGroup(cfg.GetKafkaEndpoint(), nil, groupID)
}

// SetAddress
func (c *ConsumerGroup) SetAddress(address *atlas.KafkaAddress) *ConsumerGroup {
	c.address = address
	return c
}

// SetTopic
func (c *ConsumerGroup) SetTopic(topic string) *ConsumerGroup {
	c.address = atlas.NewKafkaAddress(topic, 0)
	return c
}

// SetConsumerGroupHandler
func (c *ConsumerGroup) SetConsumerGroupHandler(handler sarama.ConsumerGroupHandler) {
	c.consumerGroupHandler = handler
}

// SetMessageProcessor
func (c *ConsumerGroup) SetMessageProcessor(processor func(*sarama.ConsumerMessage) bool) {
	c.messageProcessor = processor
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
		log.Fatalf("unable to create NewConsumerGroup for %v %v with err: %v", c.endpoint.Brokers, c.groupID, err)
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
		topics := []string{c.address.Topic}

		// Handler can be either explicitly specified, or a default one
		// Default handler can still use external c.messageProcessor
		handler := c.consumerGroupHandler
		if handler == nil {
			handler = newDefaultConsumerGroupHandler(ack, c.messageProcessor)
		}

		// Consume joins a cluster of consumers for a given list of topics
		//
		// `Consume` should be called inside an infinite loop.
		// When a server-side rebalance happens, the consumer session will need to be recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			log.Fatalf("unable to Consume topics %v with err: %v", topics, err)
		}
	}
}

// ConsumerGroupHandler instances are used to handle individual topic/partition claims.
// It also provides hooks for your consumer group session life-cycle and allow you to
// trigger logic before or after the consume loop(s).
//
// PLEASE NOTE that handlers are likely be called from several goroutines concurrently,
// ensure that all state is safely protected against race conditions.
//
// Implements sarama.ConsumerGroupHandler interface
type DefaultConsumerGroupHandler struct {
	ack       bool
	processor func(*sarama.ConsumerMessage) bool
}

// newDefaultConsumerGroupHandler
func newDefaultConsumerGroupHandler(ack bool, processor func(*sarama.ConsumerMessage) bool) *DefaultConsumerGroupHandler {
	return &DefaultConsumerGroupHandler{
		ack:       ack,
		processor: processor,
	}
}

// Implement sarama.ConsumerGroupHandler interface

// Setup is run at the beginning of a new session, before ConsumeClaim.
// Part of sarama.ConsumerGroupHandler interface
func (*DefaultConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	log.Infof("ConsumerGroupHandler.Setup() - start")
	defer log.Infof("ConsumerGroupHandler.Setup() - end")

	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
// Part of sarama.ConsumerGroupHandler interface
func (*DefaultConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Infof("ConsumerGroupHandler.Cleanup() - start")
	defer log.Infof("ConsumerGroupHandler.Cleanup() - end")

	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish
// its processing loop and exit.
// Part of sarama.ConsumerGroupHandler interface
func (h *DefaultConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Claim is a claimed Partition, so Claim refers to Partition
	log.Infof("ConsumerGroupHandler.ConsumeClaim() - start")
	defer log.Infof("ConsumerGroupHandler.ConsumeClaim() - end")

	for msg := range claim.Messages() {
		// msg.Headers
		log.Printf("Got message topic:%q partition:%d offset:%d data:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		// Call message processor
		ack := h.ack
		if h.processor == nil {
			log.Warn("no message processor specified with DefaultConsumerGroupHandler")
		} else {
			ack = h.processor(msg)
		}

		if ack {
			sess.MarkMessage(msg, "")
			log.Infof("Ack message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		}
	}
	return nil
}
