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

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/softwareid"
)

type ConsumerGroup struct {
	endpoint *Endpoint
	address  *atlas.KafkaAddress
	groupID  string
}

func NewConsumerGroup(endpoint *Endpoint, address *atlas.KafkaAddress, groupID string) *ConsumerGroup {
	return &ConsumerGroup{
		endpoint: endpoint,
		address:  address,
		groupID:  groupID,
	}
}

// ConsumeLoop runs an endless loop of kafka consumer
func (c *ConsumerGroup) ConsumeLoop(consumeNewest bool, ack bool) {

	// new configuration instance with sane defaults.
	config := sarama.NewConfig()
	config.ClientID = softwareid.Name
	if consumeNewest {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	group, err := sarama.NewConsumerGroup(c.endpoint.Brokers, c.groupID, config)
	if err != nil {
		panic(err)
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
		handler := NewConsumerGroupHandler(ack)

		// Consume joins a cluster of consumers for a given list of topics
		//
		// `Consume` should be called inside an infinite loop.
		// When a server-side rebalance happens, the consumer session will need to be recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
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
type ConsumerGroupHandler struct {
	ack bool
}

// NewConsumerGroupHandler
func NewConsumerGroupHandler(ack bool) ConsumerGroupHandler {
	return ConsumerGroupHandler{
		ack: ack,
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	log.Infof("ConsumerGroupHandler.Setup() - start")
	defer log.Infof("ConsumerGroupHandler.Setup() - end")

	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Infof("ConsumerGroupHandler.Cleanup() - start")
	defer log.Infof("ConsumerGroupHandler.Cleanup() - end")

	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish
// its processing loop and exit.
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Claim is a claimed Partition, so Claim refers to Partition
	log.Infof("ConsumerGroupHandler.ConsumeClaim() - start")
	defer log.Infof("ConsumerGroupHandler.ConsumeClaim() - end")

	for msg := range claim.Messages() {
		// msg.Headers
		log.Printf("Got message topic:%q partition:%d offset:%d data:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		if h.ack {
			sess.MarkMessage(msg, "")
			log.Infof("Ack message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		}
	}
	return nil
}
