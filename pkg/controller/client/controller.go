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
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// IncomingTasksHandler
func IncomingTasksHandler(incomingQueue, outgoingQueue chan *atlas.Task) {
	log.Infof("IncomingTasksHandler() - start")
	defer log.Infof("IncomingTasksHandler() - end")

	for {
		task := <-incomingQueue
		log.Infof("Got task %s", task)
		if task.GetType() == atlas.TaskEchoRequest {
			task := atlas.NewTask().
				SetType(atlas.TaskEchoReply).
				CreateUUID().
				SetReferenceUUIDFromString("reference: " + task.GetUUID().String()).
				SetDescription("desc")
			outgoingQueue <- task
		}
	}
}
