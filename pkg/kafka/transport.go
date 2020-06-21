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
	log "github.com/sirupsen/logrus"
)

type Transport struct {
	producer *Producer
	consumer *Consumer
	close    bool
}

// NewTransport
func NewTransport(producer *Producer, consumer *Consumer, close bool) *Transport {
	log.Infof("NewTransport() - start")
	defer log.Infof("NewTransport() - end")

	return &Transport{
		producer: producer,
		consumer: consumer,
		close:    close,
	}
}

// Close
func (t *Transport) Close() {
	log.Infof("Transport.Close() - start")
	defer log.Infof("Transport.Close() - end")

	if !t.close {
		return
	}

	if t.producer != nil {
		t.producer.Close()
		t.producer = nil
	}

	if t.consumer != nil {
		t.consumer.Close()
		t.consumer = nil
	}
}
