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
	"fmt"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	brokers []string
	groupID string
	topic   string
}

func NewConsumer(brokers []string, groupID string, topic string) *Consumer {
	return &Consumer{
		brokers: brokers,
		groupID: groupID,
		topic:   topic,
	}
}

func (c *Consumer) Consume() {
	//config := sarama.NewConfig()
	//config.Version = sarama.V2_0_0_0 // specify appropriate version
	//config.Consumer.Return.Errors = true
	//group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-group", config)
	group, err := sarama.NewConsumerGroup(c.brokers, c.groupID, nil)
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
		topics := []string{c.topic}
		handler := ConsumerGroupHandler{}

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
type ConsumerGroupHandler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}
