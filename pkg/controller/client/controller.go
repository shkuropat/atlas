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
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// SendFile sends file from client to service and receives response back (if any)
func SendFile(client atlas.ControlPlaneClient, filename string, options *DataExchangeOptions) (int64, error) {
	log.Info("SendFile() - start")
	defer log.Info("SendFile() - end")

	if _, err := os.Stat(filename); err != nil {
		log.Warnf("no file %s available err: %v", filename, err)
		return 0, err
	}

	log.Infof("Has file %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Warnf("ERROR open file %s err: %v", filename, err)
		return 0, err
	}

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(filepath.Base(filename))
	return SendReader(client, f, options)
}

// SendStdin sends STDIN from client to service and receives response back (if any)
func SendStdin(client atlas.ControlPlaneClient, options *DataExchangeOptions) (int64, error) {
	log.Info("SendStdin() - start")
	defer log.Info("SendStdin() - end")

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(os.Stdin.Name())
	return SendReader(client, os.Stdin, options)
}

// SendReader
func SendReader(client atlas.ControlPlaneClient, r io.Reader, options *DataExchangeOptions) (int64, error) {
	log.Info("SendReader() - start")
	defer log.Info("SendReader() - end")

	result := DataExchange(client, r, options)
	if result.Err == nil {
		log.Infof("DONE send %s size %d", "io.Reader", result.Send.Len)
	} else {
		log.Warnf("FAILED send %s size %d err %v", "io.Reader", result.Send.Len, result.Err)
	}

	return result.Send.Len, result.Err
}

// SendBytes
func SendBytes(client atlas.ControlPlaneClient, data []byte, options *DataExchangeOptions) (int64, error) {
	log.Info("SendBytes() - start")
	defer log.Info("SendBytes() - end")

	return SendReader(client, bytes.NewReader(data), options)
}

// SendEchoRequest
func SendEchoRequest(outgoingQueue chan *atlas.Command) {
	for i := 0; i < 5; i++ {
		command := atlas.NewCommand().SetType(atlas.CommandType_COMMAND_ECHO_REQUEST).CreateID().SetDescription("desc")
		outgoingQueue <- command

		log.Infof("Wait before send new Echo Request")
		time.Sleep(3 * time.Second)
	}
}

// IncomingCommandsHandler
func IncomingCommandsHandler(incomingQueue, outgoingQueue chan *atlas.Command) {
	log.Infof("IncomingCommandsHandler() - start")
	defer log.Infof("IncomingCommandsHandler() - end")

	for {
		cmd := <-incomingQueue
		log.Infof("Got cmd %v", cmd)
		if cmd.GetType() == atlas.CommandType_COMMAND_ECHO_REQUEST {
			command := atlas.NewCommand().
				SetType(atlas.CommandType_COMMAND_ECHO_REPLY).
				CreateID().
				SetReferenceIDFromString("reference: " + cmd.GetHeader().GetId().GetString()).
				SetDescription("desc")
			outgoingQueue <- command
		}
	}
}
