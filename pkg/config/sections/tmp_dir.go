// Copyright 2021 The Atlas Authors. All rights reserved.
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

package sections

import (
	"github.com/binarly-io/atlas/pkg/config/items"
)

// TmpDirConfigurator
type TmpDirConfigurator interface {
	GetTmpDirDir() string
	GetTmpDirPattern() string
}

// Interface compatibility
var _ TmpDirConfigurator = TmpDir{}

// TmpDir
type TmpDir struct {
	TmpDir *items.TmpItem `mapstructure:"tmpdir"`
}

// TmpDirNormalize
func (c TmpDir) TmpDirNormalize() {
	if c.TmpDir == nil {
		c.TmpDir = items.NewTmpItem()
	}
}

// GetTmpDirDir
func (c TmpDir) GetTmpDirDir() string {
	return c.TmpDir.GetDir()
}

// GetTmpDirPattern
func (c TmpDir) GetTmpDirPattern() string {
	return c.TmpDir.GetPattern()
}
