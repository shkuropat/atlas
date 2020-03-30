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
	"context"
	log "github.com/golang/glog"
	"github.com/sunsingerus/mservice/pkg/transiever"
	"io"
	"os"

	pb "github.com/sunsingerus/mservice/pkg/api/mservice"
)

func Init() {
	transiever.Init()
}

func GetOutgoingQueue() chan *pb.Command {
	return transiever.GetOutgoingQueue()
}

func GetIncomingQueue() chan *pb.Command {
	return transiever.GetIncomingQueue()
}

func RunMServiceControlPlaneClient(client pb.MServiceControlPlaneClient) {
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

func StreamDataChunks(client pb.MServiceControlPlaneClient, dataSource io.Reader) (n int64, err error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Infof("rpcData()")
	rpcData, err := client.Data(ctx)
	if err != nil {
		log.Fatalf("client.Data() failed %v", err)
		os.Exit(1)
	}
	defer rpcData.CloseSend()
	dataChunkStream, err := pb.OpenDataChunkStream(
		rpcData,
		uint32(pb.DataChunkType_DATA_CHUNK_DATA),
		"",
		0,
		"123",
		"desc",
	)
	n, err = io.Copy(dataChunkStream, dataSource)
	err = dataChunkStream.Close()
	return n, err
}
