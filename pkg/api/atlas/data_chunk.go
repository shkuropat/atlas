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

const (
	// Due to first enum value has to be zero in proto3
	DataChunkTypeReserved int32 = 0
	// Unspecified means data chunk type is unknown
	DataChunkTypeUnspecified int32 = 100
	// Data chunk
	DataChunkTypeData int32 = 200
)

var DataChunkTypeEnum = NewEnum()

func init() {
	DataChunkTypeEnum.MustRegister("DataChunkTypeReserved", DataChunkTypeReserved)
	DataChunkTypeEnum.MustRegister("DataChunkTypeUnspecified", DataChunkTypeUnspecified)
	DataChunkTypeEnum.MustRegister("DataChunkTypeData", DataChunkTypeData)
}

// NewDataChunk
func NewDataChunk() *DataChunk {
	return &DataChunk{
		Header: NewMetadata().SetType(DataChunkTypeData),
	}
}

// ensureTransportMetadata
func (m *DataChunk) ensureTransportMetadata() *DataChunk {
	if m.TransportMetadataOptional == nil {
		m.TransportMetadataOptional = new(DataChunk_TransportMetadata)
	}
	if m.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata == nil {
		m.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata = new(Metadata)
	}
	return m
}

// HasTransportMetadata
func (m *DataChunk) HasTransportMetadata() bool {
	if m.TransportMetadataOptional == nil {
		return false
	}
	if m.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata == nil {
		return false
	}
	return true
}

// SetTransportMetadata
func (m *DataChunk) SetTransportMetadata(metadata *Metadata) *DataChunk {
	if m.TransportMetadataOptional == nil {
		m.TransportMetadataOptional = new(DataChunk_TransportMetadata)
	}
	m.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata = metadata
	return m
}

// ensurePayloadMetadata
func (m *DataChunk) ensurePayloadMetadata() *DataChunk {
	if m.PayloadMetadataOptional == nil {
		m.PayloadMetadataOptional = new(DataChunk_PayloadMetadata)
	}
	if m.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata == nil {
		m.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata = new(Metadata)
	}
	return m
}

// HasPayloadMetadata
func (m *DataChunk) HasPayloadMetadata() bool {
	if m.PayloadMetadataOptional == nil {
		return false
	}
	if m.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata == nil {
		return false
	}
	return true
}

// SetPayloadMetadata
func (m *DataChunk) SetPayloadMetadata(metadata *Metadata) *DataChunk {
	if m.PayloadMetadataOptional == nil {
		m.PayloadMetadataOptional = new(DataChunk_PayloadMetadata)
	}
	m.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata = metadata
	return m
}

// SetData
func (m *DataChunk) SetData(data []byte) *DataChunk {
	if m != nil {
		m.Data = data
	}
	return m
}

// String
func (m *DataChunk) String() string {
	return "no be implemented"
}
