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

package controller

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// CommandSenderReceiver defines transport level interface (for both client and server),
// which serves Command streams bi-directionally.
type CommandSenderReceiver interface {
	Send(*atlas.Command) error
	Recv() (*atlas.Command, error)
}

func CommandsExchangeEndlessLoop(CommandSenderReceiver CommandSenderReceiver) {
	waitIncoming := make(chan bool)
	waitOutgoing := make(chan bool)

	// Recv() loop
	go func() {
		for {
			msg, err := CommandSenderReceiver.Recv()
			if msg != nil {
				log.Infof("CommandsExchangeEndlessLoop.Recv() got msg")
				GetIncoming() <- msg
			}
			if err == nil {
				// All went well, ready to receive more data
			} else if err == io.EOF {
				// Correct EOF
				log.Infof("CommandsExchangeEndlessLoop.Recv() got EOF")

				close(waitIncoming)
				return
			} else {
				// Stream broken
				log.Infof("CommandsExchangeEndlessLoop.Recv() got err: %v", err)

				close(waitIncoming)
				return
			}
		}
	}()

	// Send() loop
	go func() {
		for {
			select {
			case <-waitIncoming:
				// Incoming stream from this client is closed/broken, no need to wait commands for it
				close(waitOutgoing)
				return
			case command := <-GetOutgoing():
				log.Infof("got command to send")
				err := CommandSenderReceiver.Send(command)
				if err == nil {
					// All went well
					log.Infof("CommandsExchangeEndlessLoop.Send() OK")
				} else if err == io.EOF {
					log.Infof("CommandsExchangeEndlessLoop.Send() got EOF")

					close(waitOutgoing)
					return
				} else {
					log.Fatalf("CommandsExchangeEndlessLoop.Send() got err: %v", err)

					close(waitOutgoing)
					return
				}
			}
		}
	}()

	<-waitIncoming
	<-waitOutgoing
}

var (
	incomingBacklog int32 = 100
	incoming        chan *atlas.Command
	outgoingBacklog int32 = 100
	outgoing        chan *atlas.Command
)

func Init() {
	incoming = make(chan *atlas.Command, incomingBacklog)
	outgoing = make(chan *atlas.Command, outgoingBacklog)
}

func GetOutgoing() chan *atlas.Command {
	return outgoing
}

func GetIncoming() chan *atlas.Command {
	return incoming
}
