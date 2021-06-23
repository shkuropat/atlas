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

package util

import "path/filepath"

func FullPath(root, dir, file string) string {
	if filepath.IsAbs(file) {
		return filepath.Clean(file)
	}

	rel := filepath.Join(dir, file)
	if filepath.IsAbs(rel) {
		return filepath.Clean(rel)
	}

	return filepath.Join(root, dir, file)
}
