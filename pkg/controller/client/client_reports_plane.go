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
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// GetTaskStatus requests status(es) of the task
func GetTaskStatus(ReportsPlaneClient atlas.ReportsPlaneClient, taskUUID *atlas.UUID) *DataExchangeResult {
	log.Infof("Status() - start")
	defer log.Infof("Status() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// One object request
	taskAddress := atlas.NewAddress(taskUUID)
	objectRequest := atlas.NewObjectRequest()
	objectRequest.SetAddress(taskAddress)
	// Multi-object request
	request := atlas.NewObjectsRequest()
	request.SetRequestDomain(atlas.DomainTask)
	request.SetResultDomain(atlas.DomainStatus)
	request.Append(objectRequest)
	list, err := ReportsPlaneClient.ObjectsReport(ctx, request)
	// Unify call result
	result := NewDataExchangeResult()
	if len(list.GetObjectStatuses()) > 0 {
		result.Recv.ObjectStatus = atlas.NewObjectStatus().SetStatus(list.GetObjectStatuses()[0].GetStatus())
	}
	result.Error = err

	return result
}

// GetTask requests task
func GetTask(ReportsPlaneClient atlas.ReportsPlaneClient, taskUUID *atlas.UUID) *DataExchangeResult {
	log.Infof("Task() - start")
	defer log.Infof("Task() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// One object request
	objectRequest := atlas.NewObjectRequest()
	objectRequest.SetAddress(atlas.NewAddress(taskUUID))
	// Multi-object request
	request := atlas.NewObjectsRequest()
	request.SetRequestDomain(atlas.DomainTask)
	request.SetResultDomain(atlas.DomainTask)
	request.Append(objectRequest)
	// Unify call result
	result := NewDataExchangeResult()
	result.Recv.ObjectsList, result.Error = ReportsPlaneClient.ObjectsReport(ctx, request)

	return result
}

// GetTaskReport requests report(es) of the task
func GetTaskReport(ReportsPlaneClient atlas.ReportsPlaneClient, taskUUID *atlas.UUID) *DataExchangeResult {
	log.Infof("Report() - start")
	defer log.Infof("Report() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// One object request
	objectRequest := atlas.NewObjectRequest()
	objectRequest.SetAddress(atlas.NewAddress(taskUUID))
	// Multi-object request
	request := atlas.NewObjectsRequest()
	request.SetRequestDomain(atlas.DomainTask)
	request.SetResultDomain(atlas.DomainReport)
	request.Append(objectRequest)
	// Unify call result
	result := NewDataExchangeResult()
	result.Recv.ObjectsList, result.Error = ReportsPlaneClient.ObjectsReport(ctx, request)

	return result
}

// GetTaskFiles requests file(es) of the task
func GetTaskFiles(ReportsPlaneClient atlas.ReportsPlaneClient, taskUUID *atlas.UUID) *DataExchangeResult {
	log.Infof("Files() - start")
	defer log.Infof("Files() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// One object request
	objectRequest := atlas.NewObjectRequest()
	objectRequest.SetAddress(atlas.NewAddress(taskUUID))
	// Multi-object request
	request := atlas.NewObjectsRequest()
	request.SetRequestDomain(atlas.DomainTask)
	request.SetResultDomain(atlas.DomainFile)
	request.Append(objectRequest)
	// Unify call result
	result := NewDataExchangeResult()
	result.Recv.ObjectsList, result.Error = ReportsPlaneClient.ObjectsReport(ctx, request)

	return result
}
