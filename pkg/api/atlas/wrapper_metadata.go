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

// NewMetadata
func NewMetadata() *Metadata {
	return new(Metadata)
}

// SetFilename
func (m *Metadata) SetFilename(filename string) {
	if filename == "" {
		return
	}

	if m.FilenameOptional == nil {
		m.FilenameOptional = new(Metadata_Filename)
	}
	m.FilenameOptional.(*Metadata_Filename).Filename = filename
}

// SetURL
func (m *Metadata) SetURL(url string) {
	if url == "" {
		return
	}

	if m.UrlOptional == nil {
		m.UrlOptional = new(Metadata_Url)
	}
	m.UrlOptional.(*Metadata_Url).Url = url
}
