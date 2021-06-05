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

package journal

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
)

const (
	// Unknown endpoint
	EndpointUnknown int32 = 0

	// Control Plane
	ControlPlane          int32 = 100
	EndpointTasks         int32 = ControlPlane + 1
	EndpointDataChunks    int32 = ControlPlane + 2
	EndpointUploadObject  int32 = ControlPlane + 3
	EndpointUploadObjects int32 = ControlPlane + 4
	EndpointMetrics       int32 = ControlPlane + 5
	EndpointStatusObject  int32 = ControlPlane + 6
	EndpointStatusObjects int32 = ControlPlane + 7

	// Reports Plane
	ReportsPlane    int32 = 200
	EndpointReport  int32 = ReportsPlane + 1
	EndpointReports int32 = ReportsPlane + 2
)

var (
	EndpointTypeEnum = atlas.NewEnum()
)

func init() {
	EndpointTypeEnum.MustRegister("EndpointUnknown", EndpointUnknown)
	// Control Plane
	EndpointTypeEnum.MustRegister("EndpointTasks", EndpointTasks)
	EndpointTypeEnum.MustRegister("EndpointDataChunks", EndpointDataChunks)
	EndpointTypeEnum.MustRegister("EndpointUploadObject", EndpointUploadObject)
	EndpointTypeEnum.MustRegister("EndpointUploadObjects", EndpointUploadObjects)
	EndpointTypeEnum.MustRegister("EndpointMetrics", EndpointMetrics)
	EndpointTypeEnum.MustRegister("EndpointStatusObject", EndpointStatusObject)
	EndpointTypeEnum.MustRegister("EndpointStatusObjects", EndpointStatusObjects)
	// Reports Plane
	EndpointTypeEnum.MustRegister("EndpointReport", EndpointReport)
	EndpointTypeEnum.MustRegister("EndpointReports", EndpointReports)
}
