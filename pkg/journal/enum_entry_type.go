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
	EntryTypeUnknown          int32 = 0
	EntryTypeRequestStart     int32 = 1
	EntryTypeSaveData         int32 = 100
	EntryTypeSaveDataError    int32 = 101
	EntryTypeProcessData      int32 = 200
	EntryTypeProcessDataError int32 = 201
	EntryTypeResult           int32 = 300
	EntryTypeSaveTask         int32 = 400
	EntryTypeSaveTaskError    int32 = 401
	EntryTypeProcessTask      int32 = 500
	EntryTypeProcessTaskError int32 = 501
	EntryTypeLookup           int32 = 600
	EntryTypeLookupError      int32 = 601
	EntryTypeRequestCompleted int32 = 10000
	EntryTypeRequestError     int32 = 10001
)

var (
	EntryTypeEnum = atlas.NewEnum()
)

func init() {
	EntryTypeEnum.MustRegister("EntryTypeUnknown", EntryTypeUnknown)
	EntryTypeEnum.MustRegister("EntryTypeRequestStart", EntryTypeRequestStart)
	EntryTypeEnum.MustRegister("EntryTypeSaveData", EntryTypeSaveData)
	EntryTypeEnum.MustRegister("EntryTypeSaveDataError", EntryTypeSaveDataError)
	EntryTypeEnum.MustRegister("EntryTypeProcessData", EntryTypeProcessData)
	EntryTypeEnum.MustRegister("EntryTypeProcessDataError", EntryTypeProcessDataError)
	EntryTypeEnum.MustRegister("EntryTypeResult", EntryTypeResult)
	EntryTypeEnum.MustRegister("EntryTypeSaveTask", EntryTypeSaveTask)
	EntryTypeEnum.MustRegister("EntryTypeSaveTaskError", EntryTypeSaveTaskError)
	EntryTypeEnum.MustRegister("EntryTypeProcessTask", EntryTypeProcessTask)
	EntryTypeEnum.MustRegister("EntryTypeProcessTaskError", EntryTypeProcessTaskError)
	EntryTypeEnum.MustRegister("EntryTypeLookup", EntryTypeLookup)
	EntryTypeEnum.MustRegister("EntryTypeLookupError", EntryTypeLookupError)
	EntryTypeEnum.MustRegister("EntryTypeRequestCompleted", EntryTypeRequestCompleted)
	EntryTypeEnum.MustRegister("EntryTypeRequestError", EntryTypeRequestError)
}
