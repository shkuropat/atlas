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

package controller_service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// ReportsPlaneServer
type ReportsPlaneServer struct {
	atlas.UnimplementedReportsPlaneServer
}

// NewReportsPlaneServer
func NewReportsPlaneServer() *ReportsPlaneServer {
	return &ReportsPlaneServer{}
}

// ObjectsReportHandler is a user-provided handler for ObjectsReport call
var ObjectsReportHandler = func(context.Context, *atlas.ObjectsRequest, jwt.Claims) (*atlas.ObjectsList, error) {
	return nil, HandlerUnavailable
}

// ObjectsReport gRPC call
func (s *ReportsPlaneServer) ObjectsReport(ctx context.Context, req *atlas.ObjectsRequest) (*atlas.ObjectsList, error) {
	log.Info("ObjectsReport() - start")
	defer log.Info("ObjectsReport() - end")

	if ObjectsReportHandler == nil {
		return nil, HandlerUnavailable
	}

	var claims jwt.Claims
	if ClaimsExtractor != nil {
		claims = ClaimsExtractor(ctx)
	}
	return ObjectsReportHandler(ctx, req, claims)
}
