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
	"fmt"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

type ReportsPlaneServer struct {
	atlas.UnimplementedReportsPlaneServer
}

func NewReportsPlaneServer() *ReportsPlaneServer {
	return &ReportsPlaneServer{}
}

// ReportsHandler is a user-provided handler for Reports call
var ReportsHandler func(atlas.ReportsPlane_ReportsServer, jwt.MapClaims) error

// Reports gRPC call
func (s *ReportsPlaneServer) Reports(ReportsServer atlas.ReportsPlane_ReportsServer) error {
	log.Info("Reports() - start")
	defer log.Info("Reports() - end")

	if ReportsHandler == nil {
		return fmt.Errorf("no ReportsHandler provided")
	}

	metadata := fetchMetadata(ReportsServer.Context())
	return ReportsHandler(ReportsServer, metadata)
}
