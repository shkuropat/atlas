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

import "fmt"

// EndpointIDType represents endpoint id for which journal is created. Be it data endpoint or commands endpoint.
type EndpointIDType uint32

const (
	EndpointUnknown        EndpointIDType = 0
	EndpointUnknownName    string         = "UNKNOWN"
	EndpointDataChunks     EndpointIDType = 1
	EndpointDataChunksName string         = "EndpointDataChunks"
	EndpointReports        EndpointIDType = 2
	EndpointReportsName    string         = "EndpointReports"
	EndpointFileStatus     EndpointIDType = 3
	EndpointFileStatusName string         = "EndpointRFileStatus"
)

// TODO extract dictionary as external entry

var (
	// map type to name
	endpointName = map[EndpointIDType]string{
		EndpointUnknown:    EndpointUnknownName,
		EndpointDataChunks: EndpointDataChunksName,
		EndpointReports:    EndpointReportsName,
		EndpointFileStatus: EndpointFileStatusName,
	}
	// map name to type
	endpointValue = map[string]EndpointIDType{
		EndpointUnknownName:    EndpointUnknown,
		EndpointDataChunksName: EndpointDataChunks,
		EndpointReportsName:    EndpointReports,
		EndpointFileStatusName: EndpointFileStatus,
	}
)

var (
	ErrEndpointIDBusy = fmt.Errorf("busy")
)

// RegisterEndpointID registers new EndpointID with specified name
func RegisterEndpointID(id EndpointIDType, name string) error {
	// Check availability
	if n, ok := endpointName[id]; ok {
		if n != name {
			return ErrEndpointIDBusy
		}
	}
	if v, ok := endpointValue[name]; ok {
		if v != id {
			return ErrEndpointIDBusy
		}
	}

	// Register
	endpointName[id] = name
	endpointValue[name] = id

	return nil
}
