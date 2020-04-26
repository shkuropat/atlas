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

package transiever_client

import (
	"bytes"
	"context"
	"github.com/binarly-io/binarly-atlas/pkg/transiever"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	pb "github.com/binarly-io/binarly-atlas/pkg/api/mservice"
)

func CommandsExchange(client pb.MServiceControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//this code sends token per each RPC call:
	//	md := metadata.Pairs("authorization", "my-secret-token")
	//	ctx = metadata.NewOutgoingContext(ctx, md)

	rpcCommands, err := client.Commands(ctx)
	if err != nil {
		log.Fatalf("client.Control() failed %v", err)
		os.Exit(1)
	}
	defer rpcCommands.CloseSend()

	log.Infof("Commands() called")
	transiever.CommandsExchangeEndlessLoop(rpcCommands)
}

func DataExchange(
	ControlPlaneClient pb.MServiceControlPlaneClient,
	metadata *pb.Metadata,
	src io.Reader,
	recv bool,
) (int64, int64, *bytes.Buffer, error) {
	var (
		sent, received int64
		buf            *bytes.Buffer
	)

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Infof("DataChunks()")
	DataChunksClient, err := ControlPlaneClient.DataChunks(ctx)
	if err != nil {
		log.Fatalf("ControlPlaneClient.DataChunks() failed %v", err)
		os.Exit(1)
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

	if src != nil {
		sent, err = pb.SendDataChunkFile(DataChunksClient, metadata, src)
		if err != nil {
			return sent, 0, nil, err
		}
	}

	if recv {
		received, buf, err = pb.RecvDataChunkFile(DataChunksClient)
	}

	return sent, received, buf, err
}
