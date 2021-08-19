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

// StatusCode represents all codes of statuses
const (
	// Due to first enum value has to be zero in proto3
	StatusCodeReserved int32 = 0
	// Unspecified means we do not know its type
	StatusCodeUnspecified int32 = 100
	// Object found
	StatusCodeOK int32 = 200
	// Object created
	StatusCodeCreated int32 = 201
	// Object accepted
	StatusCodeAccepted int32 = 202
	// Not all parts/objects requested were found
	StatusCodePartial int32 = 206
	// All objects found
	StatusCodeFoundAll int32 = 220
	// Object is in progress of something
	StatusCodeInProgress int32 = 230
	// Object moved to other location
	StatusCodeMovedPermanently int32 = 301
	// Object not found
	StatusCodeNotFound int32 = 404
	// Object not ready
	StatusCodeNotReady int32 = 405
	// Object has failed due to internal error
	StatusCodeInternalError int32 = 500
	// Object failed somehow
	StatusCodeFailed int32 = 550
)

var StatusCodeEnum = NewEnum()

func init() {
	StatusCodeEnum.MustRegister("StatusCodeReserved", StatusCodeReserved)
	StatusCodeEnum.MustRegister("StatusCodeUnspecified", StatusCodeUnspecified)
	StatusCodeEnum.MustRegister("StatusCodeOK", StatusCodeOK)
	StatusCodeEnum.MustRegister("StatusCodeCreated", StatusCodeCreated)
	StatusCodeEnum.MustRegister("StatusCodeAccepted", StatusCodeAccepted)
	StatusCodeEnum.MustRegister("StatusCodePartial", StatusCodePartial)
	StatusCodeEnum.MustRegister("StatusCodeFoundAll", StatusCodeFoundAll)
	StatusCodeEnum.MustRegister("StatusCodeInProgress", StatusCodeInProgress)
	StatusCodeEnum.MustRegister("StatusCodeMovedPermanently", StatusCodeMovedPermanently)
	StatusCodeEnum.MustRegister("StatusCodeNotFound", StatusCodeNotFound)
	StatusCodeEnum.MustRegister("StatusCodeNotReady", StatusCodeNotReady)
	StatusCodeEnum.MustRegister("StatusCodeInternalError", StatusCodeInternalError)
	StatusCodeEnum.MustRegister("StatusCodeFailed", StatusCodeFailed)
}
