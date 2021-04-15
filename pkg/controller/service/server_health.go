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

	atlas_health "github.com/binarly-io/atlas/pkg/api/health"
)

type HealthServer struct {
	atlas_health.UnimplementedHealthServer
}

func NewHealthServer() *HealthServer {
	return &HealthServer{}
}

func (h *HealthServer) Check(ctx context.Context, args *atlas_health.HealthCheckRequest) (*atlas_health.HealthCheckResponse, error) {
	return &atlas_health.HealthCheckResponse{
		Status: atlas_health.ServingStatus_SERVING,
	}, nil
}

func (h *HealthServer) Watch(*atlas_health.HealthCheckRequest, atlas_health.Health_WatchServer) error {
	return nil
}
