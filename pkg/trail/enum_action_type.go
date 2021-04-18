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

package trail

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
)

const (
	ActionUnknown          int32 = 0
	ActionRequestStart     int32 = 1
	ActionSaveData         int32 = 2
	ActionSaveDataError    int32 = 3
	ActionProcessData      int32 = 4
	ActionProcessDataError int32 = 5
	ActionRequestCompleted int32 = 6
	ActionRequestError     int32 = 7
)

var (
	ActionTypeEnum = atlas.NewEnum()
)

func init() {
	ActionTypeEnum.MustRegister("ActionUnknown", ActionUnknown)
	ActionTypeEnum.MustRegister("ActionRequestStart", ActionRequestStart)
	ActionTypeEnum.MustRegister("ActionSaveData", ActionSaveData)
	ActionTypeEnum.MustRegister("ActionSaveDataError", ActionSaveDataError)
	ActionTypeEnum.MustRegister("ActionProcessData", ActionProcessData)
	ActionTypeEnum.MustRegister("ActionProcessDataError", ActionProcessDataError)
	ActionTypeEnum.MustRegister("ActionRequestCompleted", ActionRequestCompleted)
	ActionTypeEnum.MustRegister("ActionRequestError", ActionRequestError)
}
