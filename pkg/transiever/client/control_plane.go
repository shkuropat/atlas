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
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	log "github.com/golang/glog"

	pb "github.com/sunsingerus/mservice/pkg/api/mservice"
)

func RunMServiceControlPlaneClient(client pb.MServiceControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpcControl, err := client.Commands(ctx)
	if err != nil {
		log.Fatalf("client.Control() failed %v", err)
		os.Exit(1)
	}
	defer rpcControl.CloseSend()

	log.Infof("rpcControl()")

	waitc := make(chan struct{})
	go func() {
		for {
			log.Infof("before Recv()")
			msg, err := rpcControl.Recv()
			log.Infof("after Recv()")
			if err == io.EOF {
				log.Infof("Recv() received EOF, return from func")
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to rpcControl.Recv() %v", err)
			}
			log.Infof("Control() - got message %v", msg)

			log.Infof("before Recv() sleep")
			time.Sleep(time.Second)
			log.Infof("after Recv() sleep")
		}
	}()

	log.Infof("sleep")
	time.Sleep(5 * time.Second)
	log.Infof("continue")

	for i := 0; i < 5; i++ {
		command := pb.NewCommand(
			pb.CommandType_COMMAND_ECHO_REPLY,
			"",
			0,
			"21-43-65-"+strconv.Itoa(i),
			"",
			0,
			0,
			"desc",
		)
		log.Infof("before Send()")

		err := rpcControl.Send(command)
		if err == io.EOF {
			log.Infof("Send() received EOF, return from func")
			return
		}
		if err != nil {
			log.Fatalf("failed to Send() %v", err)
			os.Exit(1)
		}
		log.Infof("after Send()")

		log.Infof("before Send() sleep")
		time.Sleep(3 * time.Second)
		log.Infof("after Send() sleep")
	}

	<-waitc
}

func SendFile(client pb.MServiceControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpcData, err := client.Data(ctx)
	if err != nil {
		log.Fatalf("client.Data() failed %v", err)
		os.Exit(1)
	}
	defer rpcData.CloseSend()

	log.Infof("rpcData()")
	log.Infof("sleep")
	time.Sleep(5 * time.Second)
	log.Infof("continue")

	for i := 0; i < 5; i++ {
		command := pb.NewDataChunk(
			[]byte(fmt.Sprintf("data piece %d", i)),
		)
		log.Infof("before Send()")

		err := rpcData.Send(command)
		if err == io.EOF {
			log.Infof("Send() received EOF, return from func")
			return
		}
		if err != nil {
			log.Fatalf("failed to Send() %v", err)
			os.Exit(1)
		}
		log.Infof("after Send()")

		log.Infof("before Send() sleep")
		time.Sleep(3 * time.Second)
		log.Infof("after Send() sleep")
	}
}
