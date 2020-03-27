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
	"strconv"
	"time"

	log "github.com/golang/glog"

	pb "github.com/sunsingerus/mservice/pkg/mservice"
)

func DispatchEchoRequest(outgoingQueue chan *pb.Command) {
	for i := 0; i < 5; i++ {
		command := pb.NewCommand(
			pb.CommandType_COMMAND_ECHO_REQUEST,
			"",
			0,
			"12-34-56-"+strconv.Itoa(i),
			"",
			0,
			0,
			"desc",
		)
		log.Infof("before Transmit")
		outgoingQueue <- command
		log.Infof("after Transmit")

		log.Infof("before Transmit sleep")
		time.Sleep(3 * time.Second)
		log.Infof("after Transmit sleep")
	}
}

func HandleIncomingCommands(incomingQueue chan *pb.Command) {
	for {
		cmd := <-incomingQueue
		log.Infof("Got cmd %v", cmd)
	}
}
