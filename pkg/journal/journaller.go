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
	"github.com/binarly-io/atlas/pkg/context"
)

// Journaller
type Journaller interface {
	RequestStart(ctx *context.Context)

	SaveData(
		ctx *context.Context,

		dataS3Address *atlas.S3Address,
		dataSize int64,
		dataMetadata *atlas.Metadata,
		data []byte,
	)

	SaveDataError(
		ctx *context.Context,
		callErr error,
	)

	ProcessData(
		ctx *context.Context,

		dataS3Address *atlas.S3Address,
		dataSize int64,
		dataMetadata *atlas.Metadata,
	)

	ProcessDataError(
		ctx *context.Context,
		callErr error,
	)

	RequestCompleted(
		ctx *context.Context,
	)

	RequestError(
		ctx *context.Context,
		callErr error,
	)
}
