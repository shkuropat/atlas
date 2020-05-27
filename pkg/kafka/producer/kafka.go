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
)

type Producer struct {
	brokers []string
	topic   string
}

// NewProducer
func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		brokers: brokers,
		topic:   topic,
	}
}

// Send
func (p *Producer) Send(data []byte) error {
	producer, err := sarama.NewSyncProducer(p.brokers, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Error(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
		// The partitioning key for this message.
		// Key: 1
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Errorf("FAILED to send message: %s\n", err)
	} else {
		log.Infof("> message sent to partition %d at offset %d\n", partition, offset)
	}

	return err
}
