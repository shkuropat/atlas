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
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/controller"
)

func GetOutgoingQueue() chan *atlas.Task {
	return controller.GetOutgoing()
}

func GetIncomingQueue() chan *atlas.Task {
	return controller.GetIncoming()
}

// ControlPlaneServer
type ControlPlaneServer struct {
	atlas.UnimplementedControlPlaneServer
}

// NewControlPlaneServer
func NewControlPlaneServer() *ControlPlaneServer {
	return &ControlPlaneServer{}
}

// TasksHandler is a user-provided handler for Tasks call
var TasksHandler = func(atlas.ControlPlane_TasksServer, jwt.MapClaims) error {
	return HandlerUnavailable
}

// Tasks gRPC call
func (s *ControlPlaneServer) Tasks(TasksServer atlas.ControlPlane_TasksServer) error {
	log.Info("Tasks() - start")
	defer log.Info("Tasks() - end")

	if TasksHandler == nil {
		return HandlerUnavailable
	}

	metadata := fetchMetadata(TasksServer.Context())
	return TasksHandler(TasksServer, metadata)

	// controller.CommandsExchangeEndlessLoop(CommandsServer)
	// return nil
}

// DataChunksHandler is a user-provided handler for DataChunks call
var DataChunksHandler = func(atlas.ControlPlane_DataChunksServer, jwt.MapClaims) error {
	return HandlerUnavailable
}

// DataChunks gRPC call
func (s *ControlPlaneServer) DataChunks(DataChunksServer atlas.ControlPlane_DataChunksServer) error {
	log.Info("DataChunks() - start")
	defer log.Info("DataChunks() - end")

	if DataChunksHandler == nil {
		return HandlerUnavailable
	}

	metadata := fetchMetadata(DataChunksServer.Context())
	return DataChunksHandler(DataChunksServer, metadata)
}

// UploadObject is a user-provided handler for UploadObject call
var UploadObjectHandler = func(atlas.ControlPlane_UploadObjectServer, jwt.MapClaims) error {
	return HandlerUnavailable
}

// UploadObject gRPC call
func (s *ControlPlaneServer) UploadObject(UploadObjectServer atlas.ControlPlane_UploadObjectServer) error {
	log.Info("UploadObject() - start")
	defer log.Info("UploadObject() - end")

	if UploadObjectHandler == nil {
		return HandlerUnavailable
	}

	metadata := fetchMetadata(UploadObjectServer.Context())
	return UploadObjectHandler(UploadObjectServer, metadata)
}
