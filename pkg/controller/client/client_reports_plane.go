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

func Report(
	ReportsPlaneClient atlas.ReportsPlaneClient,
	meta *atlas.Metadata,
) *DataExchangeResult {
	log.Infof("Report() - start")
	defer log.Infof("Report() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := atlas.NewReportRequest()
	request.Header = meta
	result := NewDataExchangeResult()
	result.Recv.Report.ReportMulti, result.Error = ReportsPlaneClient.Report(ctx, request)

	return result
}
