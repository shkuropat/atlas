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

	pb "github.com/binarly-io/binarly-atlas/pkg/api/health"
)

type HealthServer struct {
	pb.UnimplementedHealthServer
}

func (h *HealthServer) Check(ctx context.Context, args *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthServer) Watch(*pb.HealthCheckRequest, pb.Health_WatchServer) error {
	return nil
}