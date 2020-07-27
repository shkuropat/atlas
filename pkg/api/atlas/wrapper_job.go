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

package atlas

import (
	"fmt"
	"strings"
)

// NewJob
func NewJob(addresses []*S3Address) *Job {
	return &Job{
		Addresses: addresses,
	}
}

// AddAddress
func (m *Job) AddAddress(address *S3Address) {
	m.Addresses = append(m.Addresses, address)
}

// Printable
func (m *Job) Printable() string {
	if m == nil {
		return "nil"
	}

	var parts []string
	if _type := m.GetType(); _type > 0 {
		parts = append(parts, "type:"+fmt.Sprintf("%d", _type))
	}
	if name := m.GetName(); name != "" {
		parts = append(parts, "name:"+name)
	}
	if version := m.GetVersion(); version > 0 {
		parts = append(parts, "version:"+fmt.Sprintf("%d", version))
	}
	if uuid := m.GetUuid(); uuid != nil {
		if uuid.String() != "" {
			parts = append(parts, "uuid:"+uuid.String())
		}
	}
	if uuid := m.GetUuidReference(); uuid != nil {
		if uuid.String() != "" {
			parts = append(parts, "uuidref:"+uuid.String())
		}
	}
	if timestamp := m.GetTs(); timestamp != nil {
		if timestamp.String() != "" {
			parts = append(parts, "timestamp:"+timestamp.String())
		}
	}
	if len(m.GetAddresses()) > 0 {
		parts = append(parts, "addresses:")
		for _, address := range m.GetAddresses() {
			parts = append(parts, address.Printable())
		}
	}

	return strings.Join(parts, " ")
}
