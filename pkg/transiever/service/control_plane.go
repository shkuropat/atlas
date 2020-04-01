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

package transiever_service

import (
	"fmt"
	"github.com/sunsingerus/mservice/pkg/transiever"
	"io"

	log "github.com/golang/glog"

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

type MServiceControlPlaneEndpoint struct {
	pb.UnimplementedMServiceControlPlaneServer
}

func (s *MServiceControlPlaneEndpoint) Commands(stream pb.MServiceControlPlane_CommandsServer) error {
	log.Info("Commands() called")
	defer log.Info("Commands() exited")

	transiever.CommandsExchangeEndlessLoop(stream)
	return nil
}

func (s *MServiceControlPlaneEndpoint) Data(stream pb.MServiceControlPlane_DataServer) error {
	log.Info("Data() called")
	defer log.Info("Data() exited")

	for {
		dataChunk, err := stream.Recv()
		if dataChunk != nil {
			// We have data chunk received
			filename := "not specified"
			if md := dataChunk.GetMetadata(); md != nil {
				filename = md.GetFilename()
			}
			offset := "not specified"
			off, ok := dataChunk.GetOffsetWithTest()
			if ok {
				offset = fmt.Sprintf("%d", off)
			}
			log.Infof("Data.Recv() got msg filename %s, chunk len %d, chunk offset %s, last chunk %v",
				filename,
				len(dataChunk.GetBytes()),
				offset,
				dataChunk.GetLast(),
			)
			fmt.Printf("%s\n", string(dataChunk.GetBytes()))
		}

		if err == nil {
			// All went well, ready to receive more data
		} else if err == io.EOF {
			// Correct EOF
			log.Infof("Data.Recv() get EOF")

			return nil
		} else {
			// Stream broken
			log.Infof("Data.Recv() got err: %v", err)

			return nil
		}
	}
}

func (s *MServiceControlPlaneEndpoint) Metrics(stream pb.MServiceControlPlane_MetricsServer) error {
	log.Info("Metrics() called")
	defer log.Info("Metrics() exited")

	return nil
}
