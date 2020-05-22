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

package controller_client

import (
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/binarly-atlas/pkg/api/atlas"
)

// SendDataChunkFile sends file from client to service and receives response back
func SendFile(client atlas.ControlPlaneClient, filename string) (int64, error) {
	if _, err := os.Stat(filename); err != nil {
		return 0, err
	}

	log.Infof("Has file %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Infof("ERROR open file %s", filename)
		return 0, err
	}

	log.Infof("START send file %s", filename)
	metadata := atlas.NewMetadata()
	metadata.SetFilename(filepath.Base(filename))
	n, _, _, err := DataExchange(client, metadata, f, true)
	log.Infof("DONE send file %s size %d err %v", filename, n, err)

	return n, err
}

// SendStdin sends STDIN from client to service and receives response back
func SendStdin(client atlas.ControlPlaneClient) (int64, error) {
	n, _, _, err := DataExchange(client, nil, os.Stdin, true)
	log.Infof("DONE send %s size %d err %v", os.Stdin.Name(), n, err)
	return n, err
}

func SendEchoRequest(outgoingQueue chan *atlas.Command) {
	for i := 0; i < 5; i++ {
		command := atlas.NewCommand(
			atlas.CommandType_COMMAND_ECHO_REQUEST,
			"",
			0,
			atlas.CreateNewUUID(),
			"",
			0,
			0,
			"desc",
		)
		outgoingQueue <- command

		log.Infof("Wait before send new Echo Request")
		time.Sleep(3 * time.Second)
	}
}

func IncomingCommandsHandler(incomingQueue, outgoingQueue chan *atlas.Command) {
	log.Infof("Start IncomingCommandsHandler()")
	defer log.Infof("Exit IncomingCommandsHandler()")

	for {
		cmd := <-incomingQueue
		log.Infof("Got cmd %v", cmd)
		if cmd.GetType() == atlas.CommandType_COMMAND_ECHO_REQUEST {
			command := atlas.NewCommand(
				atlas.CommandType_COMMAND_ECHO_REPLY,
				"",
				0,
				atlas.CreateNewUUID(),
				"reference: "+cmd.GetHeader().GetUuid().StringValue,
				0,
				0,
				"desc",
			)
			outgoingQueue <- command
		}
	}
}
