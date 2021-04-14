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

// NewDataChunkProperties
func NewDataChunkProperties() *DataChunkProperties {
	return &DataChunkProperties{}
}

// HasDigest
func (m *DataChunkProperties) HasDigest() bool {
	if m == nil {
		return false
	}
	return m.DigestOptional != nil
}

// SetDigest
func (m *DataChunkProperties) SetDigest(digest *Digest) *DataChunkProperties {
	if m == nil {
		return nil
	}
	if m.DigestOptional == nil {
		m.DigestOptional = new(DataChunkProperties_Digest)
	}
	m.DigestOptional.(*DataChunkProperties_Digest).Digest = digest
	return m
}

// HasLen
func (m *DataChunkProperties) HasLen() bool {
	if m == nil {
		return false
	}
	return m.LenOptional != nil
}

// SetLen
func (m *DataChunkProperties) SetLen(len int64) *DataChunkProperties {
	if m == nil {
		return nil
	}
	if m.LenOptional == nil {
		m.LenOptional = new(DataChunkProperties_Len)
	}
	m.LenOptional.(*DataChunkProperties_Len).Len = len
	return m
}

// HasOffset
func (m *DataChunkProperties) HasOffset() bool {
	if m == nil {
		return false
	}
	return m.OffsetOptional != nil
}

// SetOffset
func (m *DataChunkProperties) SetOffset(offset int64) *DataChunkProperties {
	if m == nil {
		return nil
	}
	if m.OffsetOptional == nil {
		m.OffsetOptional = new(DataChunkProperties_Offset)
	}
	m.OffsetOptional.(*DataChunkProperties_Offset).Offset = offset
	return m
}

// HasLast
func (m *DataChunkProperties) HasLast() bool {
	if m == nil {
		return false
	}
	return m.LastOptional != nil
}

// SetLast
func (m *DataChunkProperties) SetLast(last bool) *DataChunkProperties {
	if m == nil {
		return nil
	}
	if m.LastOptional == nil {
		m.LastOptional = new(DataChunkProperties_Last)
	}
	m.LastOptional.(*DataChunkProperties_Last).Last = last
	return m
}

// String
func (m *DataChunkProperties) String() string {
	return "no be implemented"
}
