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

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config"
	"github.com/binarly-io/atlas/pkg/softwareid"
)

// Producer
type Producer struct {
	endpoint *atlas.KafkaEndpoint
	address  *atlas.KafkaAddress

	config   *sarama.Config
	producer sarama.SyncProducer
}

// NewProducer
func NewProducer(endpoint *atlas.KafkaEndpoint, address *atlas.KafkaAddress) *Producer {
	var err error

	p := &Producer{}
	p.endpoint = endpoint
	p.address = address
	p.config = sarama.NewConfig()
	p.config.ClientID = softwareid.Name
	// If, this config is used to create a `SyncProducer`, both must be set
	// to true and you shall not read from the channels since the producer does
	// this internally.
	p.config.Producer.Return.Successes = true
	p.config.Producer.Return.Errors = true
	p.producer, err = sarama.NewSyncProducer(p.endpoint.GetBrokers(), p.config)
	if err != nil {
		log.Errorf("unable to create NewSyncProducer(brokers:%v). err: %v", p.endpoint.GetBrokers(), err)
		p.Close()
		return nil
	}

	return p
}

// NewProducerConfig
func NewProducerConfig(cfg config.KafkaEndpointConfig) *Producer {
	return NewProducer(
		cfg.GetKafkaEndpoint(),
		nil,
	)
}

// SetAddress
func (p *Producer) SetAddress(address *atlas.KafkaAddress) *Producer {
	p.address = address
	return p
}

// SetTopic
func (p *Producer) SetTopic(topic string) *Producer {
	p.address = atlas.NewKafkaAddress(topic, 0)
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
		Topic: p.address.Topic,
		Value: sarama.ByteEncoder(data),
		// Key
		// Headers - relayed to consumer
		// Metadata - relayed to the Successes and Errors channels
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Errorf("FAILED to send message: %s", err)
	} else {
		log.Infof("message sent to %s", MsgAddressPrintable(msg))
	}

	return err
}
