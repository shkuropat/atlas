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

package controller_client

import (
	"bytes"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// DataExchangeResult represents result of data exchange
type DataExchangeResult struct {
	// Err is an error
	Error  error
	Errors []error

	// Send describes outgoing result
	Send struct {
		Data struct {
			Len int64
		}
	}

	// Recv describes incoming result
	Recv struct {
		Data struct {
			Len      int64
			Data     *bytes.Buffer
			Metadata *atlas.Metadata
		}
		ObjectStatus *atlas.ObjectStatus
		ObjectsList  *atlas.ObjectsList
	}
}

// NewDataExchangeResult
func NewDataExchangeResult() *DataExchangeResult {
	return &DataExchangeResult{}
}

// NewDataExchangeResultError
func NewDataExchangeResultError(err error) *DataExchangeResult {
	res := NewDataExchangeResult()
	res.Error = err
	return res
}
