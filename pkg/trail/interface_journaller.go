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

// Journaller
type Journaller interface {
	SetContext(ctx Contexter) Journaller

	//
	// Expose direct access to storage via adapters.
	// Implement Adapter interface as wrappers over Adapter
	//
	NewEntry(action int32) *JournalEntry
	Adapter

	//
	// Common requests section
	//

	RequestStart()
	RequestEnd()
	RequestError(callErr error)

	//
	// In-request actions
	//

	SaveData(
		dataAddress *atlas.Address,
		dataSize int64,
		dataMetadata *atlas.Metadata,
		data []byte,
	)

	SaveDataError(callErr error)

	//
	//
	//

	ProcessData(
		dataAddress *atlas.Address,
		dataSize int64,
		dataMetadata *atlas.Metadata,
	)

	ProcessDataError(callErr error)

	//
	//
	//

	SaveTask(task *atlas.Task)

	SaveTaskError(task *atlas.Task, callErr error)

	//
	//
	//

	ProcessTask(task *atlas.Task)

	ProcessTaskError(task *atlas.Task, callErr error)
}
