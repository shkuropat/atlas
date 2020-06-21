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
	"io"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

type DataChunkTransport struct {
	Transport
}

// NewDataChunkTransport
func NewDataChunkTransport(producer *Producer, consumer *Consumer, close bool) *DataChunkTransport {
	log.Infof("NewDataChunkTransport() - start")
	defer log.Infof("NewDataChunkTransport() - end")

	return &DataChunkTransport{
		Transport{
			producer: producer,
			consumer: consumer,
			close:    close,
		},
	}
}

// Send
func (t *DataChunkTransport) Send(dataChunk *atlas.DataChunk) error {
	if buf, err := proto.Marshal(dataChunk); err == nil {
		return t.producer.Send(buf)
	} else {
		return err
	}
}

// Recv
func (t *DataChunkTransport) Recv() (*atlas.DataChunk, error) {
	msg := t.consumer.Recv()
	if msg == nil {
		// TODO not sure
		return nil, io.EOF
	}
	dataChunk := &atlas.DataChunk{}
	return dataChunk, proto.Unmarshal(msg.Value, dataChunk)
}
