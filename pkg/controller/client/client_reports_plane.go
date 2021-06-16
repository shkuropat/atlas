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

// Status requests status(es) of an object
func Status(ReportsPlaneClient atlas.ReportsPlaneClient, meta *atlas.Metadata) *DataExchangeResult {
	log.Infof("Status() - start")
	defer log.Infof("Status() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := atlas.NewObjectsRequest().SetDomain(atlas.DomainStatus).Append(atlas.NewObjectRequest().SetHeader(meta))
	result := NewDataExchangeResult()
	list, err := ReportsPlaneClient.ObjectsReport(ctx, request)
	if len(list.GetStatuses()) > 0 {
		result.Recv.Status = list.GetStatuses()[0]
	}
	result.Error = err

	return result
}

// Task
func Task(ReportsPlaneClient atlas.ReportsPlaneClient, meta *atlas.Metadata) *DataExchangeResult {
	log.Infof("Task() - start")
	defer log.Infof("Task() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := atlas.NewObjectsRequest().SetDomain(atlas.DomainTask).Append(atlas.NewObjectRequest().SetHeader(meta))
	result := NewDataExchangeResult()
	result.Recv.ObjectsList, result.Error = ReportsPlaneClient.ObjectsReport(ctx, request)

	return result
}

// Report
func Report(ReportsPlaneClient atlas.ReportsPlaneClient, meta *atlas.Metadata) *DataExchangeResult {
	log.Infof("Report() - start")
	defer log.Infof("Report() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := atlas.NewObjectsRequest().SetDomain(atlas.DomainReport).Append(atlas.NewObjectRequest().SetHeader(meta))
	result := NewDataExchangeResult()
	result.Recv.ObjectsList, result.Error = ReportsPlaneClient.ObjectsReport(ctx, request)

	return result
}
