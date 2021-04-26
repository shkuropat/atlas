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
)

// ConsumerGroupHandler instances are used to handle individual topic/partition claims.
// It also provides hooks for your consumer group session life-cycle and allow you to
// trigger logic before or after the consume loop(s).
//
// PLEASE NOTE that handlers are likely be called from several goroutines concurrently,
// ensure that all state is safely protected against race conditions.
//
// Implements sarama.ConsumerGroupHandler interface
type ConsumerGroupHandler struct {
	// ctx acts as media to pass external context into consumeMessageFunction
	ctx context.Context
	// consumeMessageFunction specifies function which will be called for each message received from Kafka
	consumeMessageFunction ConsumeMessageFunction
	// ack specifies whether to mark message as consumed
	ack bool
}

// newConsumerGroupHandler
func newConsumerGroupHandler(ctx context.Context, consumeMessageFunction ConsumeMessageFunction, ack bool) *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		ctx:                    ctx,
		consumeMessageFunction: consumeMessageFunction,
		ack:                    ack,
	}
}

// Implement sarama.ConsumerGroupHandler interface

// Setup is run at the beginning of a new session, before ConsumeClaim.
// Part of sarama.ConsumerGroupHandler interface
func (*ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	log.Infof("ConsumerGroupHandler.Setup() - start")
	defer log.Infof("ConsumerGroupHandler.Setup() - end")

	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
// Part of sarama.ConsumerGroupHandler interface
func (*ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Infof("ConsumerGroupHandler.Cleanup() - start")
	defer log.Infof("ConsumerGroupHandler.Cleanup() - end")

	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish
// its processing loop and exit.
// Part of sarama.ConsumerGroupHandler interface
func (h *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Claim is a claimed Partition, so Claim refers to Partition
	log.Infof("ConsumerGroupHandler.ConsumeClaim() - start")
	defer log.Infof("ConsumerGroupHandler.ConsumeClaim() - end")

	for msg := range claim.Messages() {
		// msg.Headers
		log.Infof("Got message %s", MsgAddressPrintable(msg))

		// Call message consumeMessageFunction
		ack := h.ack
		if h.consumeMessageFunction == nil {
			log.Warnf("no message consumeMessageFunction specified with ConsumerGroupHandler")
		} else {
			ack = h.consumeMessageFunction(h.ctx, msg)
		}

		if ack {
			sess.MarkMessage(msg, "")
			log.Infof("Ack message %s", MsgAddressPrintable(msg))
		}
	}
	return nil
}
