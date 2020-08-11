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

import "github.com/google/uuid"

// NewUUIDFromString
func NewUUIDFromString(uuid string) *UUID {
	return &UUID{
		StringValue: uuid,
	}
}

// SetFromString
func (id *UUID) SetFromString(uuid string) {
	id.StringValue = uuid
}

// CreateUUID
func CreateUUID() *UUID {
	return NewUUIDFromString(uuid.New().String())
}
