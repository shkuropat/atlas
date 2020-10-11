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
	"github.com/binarly-io/atlas/pkg/rpc_context"
)

// Journaller
type Journaller interface {
	//
	// Common requests section
	//

	RequestStart(ctx *rpc_context.RPCContext)
	RequestCompleted(
		ctx *rpc_context.RPCContext,
	)
	RequestError(
		ctx *rpc_context.RPCContext,
		callErr error,
	)

	NewEntry(ctxID *atlas.UUID, action ActionType) *Entry
	Insert(entry *Entry) error
	FindAll(entry *Entry) ([]ClickHouseEntry, error)

	//
	// In-request actions
	//

	SaveData(
		ctx *rpc_context.RPCContext,

		dataS3Address *atlas.S3Address,
		dataSize int64,
		dataMetadata *atlas.Metadata,
		data []byte,
	)

	SaveDataError(
		ctx *rpc_context.RPCContext,
		callErr error,
	)

	//
	//
	//

	ProcessData(
		ctx *rpc_context.RPCContext,

		dataS3Address *atlas.S3Address,
		dataSize int64,
		dataMetadata *atlas.Metadata,
	)

	ProcessDataError(
		ctx *rpc_context.RPCContext,
		callErr error,
	)
}
