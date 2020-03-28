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
	"io"

	log "github.com/golang/glog"

	pb "github.com/sunsingerus/mservice/pkg/api/mservice"
)

var (
	waitTransieverStarted  chan bool
	maxIncomingOutstanding int32 = 100
	incomingQueue          chan *pb.Command
	maxOutgoingOutstanding int32 = 100
	outgoingQueue          chan *pb.Command
)

func Init() {
	incomingQueue = make(chan *pb.Command, maxIncomingOutstanding)
	outgoingQueue = make(chan *pb.Command, maxOutgoingOutstanding)
	waitTransieverStarted = make(chan bool)
}

func WaitTransieverStarted() {
	<-waitTransieverStarted
}

func GetOutgoingQueue() chan *pb.Command {
	return outgoingQueue
}

func GetIncomingQueue() chan *pb.Command {
	return incomingQueue
}

type MServiceControlPlaneEndpoint struct {
	pb.UnimplementedMServiceControlPlaneServer
}

func (s *MServiceControlPlaneEndpoint) Commands(stream pb.MServiceControlPlane_CommandsServer) error {
	log.Info("Commands() called")

	close(waitTransieverStarted)

	waitIncoming := make(chan bool)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == nil {
				// All went well
				log.Infof("Recv() got msg")
				incomingQueue <- msg
			} else if err == io.EOF {
				// Correct EOF
				log.Infof("Recv() get EOF")

				close(waitIncoming)
				return
			} else {
				// Stream broken
				log.Infof("Recv() got err: %v", err)

				close(waitIncoming)
				return
			}

		}
	}()

	waitOutgoing := make(chan bool)
	go func() {
		for {
			command := <-outgoingQueue
			log.Infof("got command to send")
			err := stream.Send(command)
			if err == nil {
				// All went well
				log.Infof("Send() ok")
			} else if err == io.EOF {
				log.Infof("Send() got EOF")

				close(waitOutgoing)
				return
			} else {
				log.Fatalf("Send() got err: %v", err)

				close(waitOutgoing)
				return
			}
		}
	}()

	<-waitIncoming
	<-waitOutgoing

	return nil
}

func (s *MServiceControlPlaneEndpoint) Data(stream pb.MServiceControlPlane_DataServer) error {
	log.Info("Data() called")
	defer log.Info("Data() exited")

	for {
		dataChunk, err := stream.Recv()
		if err == nil {
			// All went well
			log.Infof("Recv() got msg")
			fmt.Printf("%s\n", string(dataChunk.Bytes))
		} else if err == io.EOF {
			// Correct EOF
			log.Infof("Recv() get EOF")

			return nil
		} else {
			// Stream broken
			log.Infof("Recv() got err: %v", err)

			return nil
		}
	}
}

func (s *MServiceControlPlaneEndpoint) Metrics(stream pb.MServiceControlPlane_MetricsServer) error {
	log.Info("Metrics() called")
	return nil
}
