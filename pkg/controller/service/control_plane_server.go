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

package controller_service

import (
	"bytes"
	"strings"

	log "github.com/golang/glog"

	pb "github.com/binarly-io/binarly-atlas/pkg/api/mservice"
	"github.com/binarly-io/binarly-atlas/pkg/controller"
)

func GetOutgoingQueue() chan *pb.Command {
	return controller.GetOutgoing()
}

func GetIncomingQueue() chan *pb.Command {
	return controller.GetIncoming()
}

type MServiceControlPlaneServer struct {
	pb.UnimplementedMServiceControlPlaneServer
}

func (s *MServiceControlPlaneServer) Commands(server pb.MServiceControlPlane_CommandsServer) error {
	log.Info("Commands() called")
	defer log.Info("Commands() exited")

	controller.CommandsExchangeEndlessLoop(server)
	return nil
}

func (s *MServiceControlPlaneServer) Data(DataChunksServer pb.MServiceControlPlane_DataChunksServer) error {
	log.Info("Data() called")
	defer log.Info("Data() exited")

	_, buf, err := pb.RecvDataChunkFile(DataChunksServer)

	// Send back
	var buf2 = &bytes.Buffer{}
	buf2.WriteString(strings.ToUpper(buf.String()))

	_, err = pb.SendDataChunkFile(DataChunksServer, pb.NewMetadata("returnback.file"), buf2)

	return err
}
