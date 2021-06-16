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

// Mode represents all types of modes in the system
const (
	ModeReserved    int32 = 0
	ModeUnspecified int32 = 100
	ModeAll         int32 = 200
	ModeAny         int32 = 300
)

var ModeEnum = NewEnum()

func init() {
	ModeEnum.MustRegister("ModeReserved", ModeReserved)
	ModeEnum.MustRegister("ModeUnspecified", ModeUnspecified)
	ModeEnum.MustRegister("ModeAll", ModeAll)
	ModeEnum.MustRegister("ModeAny", ModeAny)
}
