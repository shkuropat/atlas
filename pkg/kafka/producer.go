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
	"github.com/binarly-io/binarly-atlas/pkg/softwareid"
	log "github.com/sirupsen/logrus"
)

// Producer
type Producer struct {
	Endpoint

	config   *sarama.Config
	producer sarama.SyncProducer
}

// NewProducer
func NewProducer(endpoint Endpoint) *Producer {
	var err error

	p := &Producer{}
	p.Endpoint = endpoint
	p.config = sarama.NewConfig()
	p.config.ClientID = softwareid.Name
	p.producer, err = sarama.NewSyncProducer(p.Brokers, p.config)
	if err != nil {
		log.Error(err)
		p.Close()
		return nil
	}

	return p
}

// Close
func (p *Producer) Close() {
	if p.producer != nil {
		_ = p.producer.Close()
		p.producer = nil
	}
}

// Send
func (p *Producer) Send(data []byte) error {

	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.ByteEncoder(data),
		// Key
		// Headers - relayed to consumer
		// Metadata - relayed to the Successes and Errors channels
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Errorf("FAILED to send message: %s\n", err)
	} else {
		log.Infof("> message sent to partition %d at offset %d\n", partition, offset)
	}

	return err
}
