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
	"fmt"
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

type ControlPlaneServer struct {
	atlas.UnimplementedControlPlaneServer
}

func NewControlPlaneServer() *ControlPlaneServer {
	return &ControlPlaneServer{}
}

// TasksHandler is a user-provided handler for Commands call
var TasksHandler func(atlas.ControlPlane_TasksServer, jwt.MapClaims) error

// Tasks gRPC call
func (s *ControlPlaneServer) Tasks(TasksServer atlas.ControlPlane_TasksServer) error {
	log.Info("Tasks() - start")
	defer log.Info("Tasks() - end")

	if TasksHandler == nil {
		return fmt.Errorf("no TasksHandler provided")
	}

	metadata := fetchMetadata(TasksServer.Context())
	return TasksHandler(TasksServer, metadata)

	// controller.CommandsExchangeEndlessLoop(CommandsServer)
	// return nil
}

// DataChunksHandler is a user-provided handler for DataChunks call
var DataChunksHandler func(atlas.ControlPlane_DataChunksServer, jwt.MapClaims) error

// DataChunks gRPC call
func (s *ControlPlaneServer) DataChunks(DataChunksServer atlas.ControlPlane_DataChunksServer) error {
	log.Info("DataChunks() - start")
	defer log.Info("DataChunks() - end")

	if DataChunksHandler == nil {
		return fmt.Errorf("no DataChunks provided")
	}

	metadata := fetchMetadata(DataChunksServer.Context())
	return DataChunksHandler(DataChunksServer, metadata)
}

// UploadObject is a user-provided handler for UploadObject call
var UploadObjectHandler func(atlas.ControlPlane_UploadObjectServer, jwt.MapClaims) error

// UploadObject gRPC call
func (s *ControlPlaneServer) UploadObject(UploadObjectServer atlas.ControlPlane_UploadObjectServer) error {
	log.Info("UploadObject() - start")
	defer log.Info("UploadObject() - end")

	if UploadObjectHandler == nil {
		return fmt.Errorf("no UploadObject provided")
	}

	metadata := fetchMetadata(UploadObjectServer.Context())
	return UploadObjectHandler(UploadObjectServer, metadata)
}

// StatusObjectHandler is a user-provided handler for StatusObject call
var StatusObjectHandler func(*atlas.StatusRequest, jwt.MapClaims) (*atlas.Status, error)

// StatusObject gRPC call
func (s *ControlPlaneServer) StatusObject(ctx context.Context, req *atlas.StatusRequest) (*atlas.Status, error) {
	log.Info("StatusObject() - start")
	defer log.Info("StatusObject() - end")

	if StatusObjectHandler == nil {
		return nil, fmt.Errorf("no StatusObjectHandler provided")
	}

	metadata := fetchMetadata(ctx)
	return StatusObjectHandler(req, metadata)
}

// StatusObjectsHandler is a user-provided handler for StatusObjects call
var StatusObjectsHandler func(*atlas.StatusRequestMulti, jwt.MapClaims) (*atlas.StatusMulti, error)

// StatusObjects gRPC call
func (s *ControlPlaneServer) StatusObjects(ctx context.Context, req *atlas.StatusRequestMulti) (*atlas.StatusMulti, error) {
	log.Info("StatusObjects() - start")
	defer log.Info("StatusObjects() - end")

	if StatusObjectsHandler == nil {
		return nil, fmt.Errorf("no StatusObjectsHandler provided")
	}

	metadata := fetchMetadata(ctx)
	return StatusObjectsHandler(req, metadata)
}
