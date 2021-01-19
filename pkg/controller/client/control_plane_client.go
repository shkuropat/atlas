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
	"context"
	"io"
	"os"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/controller"
	log "github.com/sirupsen/logrus"
)

// CommandsExchange exchanges commands
func CommandsExchange(ControlPlaneClient atlas.ControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//this code sends token per each RPC call:
	//	md := metadata.Pairs("authorization", "my-secret-token")
	//	ctx = metadata.NewOutgoingContext(ctx, md)

	rpcCommands, err := ControlPlaneClient.Commands(ctx)
	if err != nil {
		log.Fatalf("ControlPlaneClient.Control() failed %v", err)
		os.Exit(1)
	}
	defer rpcCommands.CloseSend()

	log.Infof("Commands() called")
	controller.CommandsExchangeEndlessLoop(rpcCommands)
}

// DataExchange send data to server and receives back reply (if needed)
func DataExchange(
	ControlPlaneClient atlas.ControlPlaneClient,
	src io.Reader,
	options *DataExchangeOptions,
) *DataExchangeResult {
	log.Infof("DataExchange() - start")
	defer log.Infof("DataExchange() - end")

	result := NewDataExchangeResult()

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var DataChunksClient atlas.ControlPlane_DataChunksClient

	DataChunksClient, result.Err = ControlPlaneClient.DataChunks(ctx)
	if result.Err != nil {
		log.Errorf("ControlPlaneClient.DataChunks() failed %v", result.Err)
		return result
	}

	defer func() {
		// This is hand-made flush() replacement for gRPC
		// It is required in order to flush all outstanding data before
		// context's cancel() is called, which simply discards all outstanding data.
		// On receiving end, when cancel() is the first in the race, f receives 'cancel' and (sometimes) no data
		// instead of complete set of data and EOF
		// See https://github.com/grpc/grpc-go/issues/1714 for more details
		DataChunksClient.CloseSend()
		DataChunksClient.Recv()
	}()

	f, err := atlas.OpenDataChunkFileWOptions(
		DataChunksClient,
		&atlas.DataChunkFileOptions{
			Metadata:   options.GetMetadata(),
			Compress:   options.GetCompress(),
			Decompress: options.GetDecompress(),
		})
	if err != nil {
		log.Errorf("ControlPlaneClient.DataChunks() failed %v", result.Err)
		result.Err = err
		return result
	}

	if src != nil {
		// We have something to send
		result.Send.Len, result.Err = f.ReadFrom(src)
		if result.Err != nil {
			log.Warnf("SendDataChunkFile() failed with err %v", result.Err)
			return result
		}
	}

	if options.GetWaitReply() {
		// We should wait for reply
		result.Receive.Len, result.Receive.Data, result.Err = f.WriteToBuf()
		if result.Err != nil {
			log.Warnf("RecvDataChunkFileIntoBuf() failed with err %v", result.Err)
			return result
		}
	}
	result.Receive.Metadata = f.DataChunkFile.PayloadMetadata

	return result
}
