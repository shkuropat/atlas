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

package base

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
)

const (
	EndpointUnknown    int32 = 0
	EndpointDataChunks int32 = 1
	EndpointReports    int32 = 2
	EndpointFileStatus int32 = 3
)

var (
	EndpointTypeEnum = atlas.NewEnum()
)

func init() {
	EndpointTypeEnum.MustRegister("EndpointUnknown", EndpointUnknown)
	EndpointTypeEnum.MustRegister("EndpointDataChunks", EndpointDataChunks)
	EndpointTypeEnum.MustRegister("EndpointReports", EndpointReports)
	EndpointTypeEnum.MustRegister("EndpointRFileStatus", EndpointFileStatus)
}
