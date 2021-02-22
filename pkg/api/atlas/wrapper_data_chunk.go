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

var DataChunkTypeEnum = NewEnum()

func init() {
	DataChunkTypeEnum.MustRegister("DATA_CHUNK_TYPE_RESERVED", int32(DataChunkType_DATA_CHUNK_TYPE_RESERVED))
	DataChunkTypeEnum.MustRegister("DATA_CHUNK_TYPE_UNSPECIFIED", int32(DataChunkType_DATA_CHUNK_TYPE_UNSPECIFIED))
	DataChunkTypeEnum.MustRegister("DATA_CHUNK_TYPE_DATA", int32(DataChunkType_DATA_CHUNK_TYPE_DATA))
}

// NewDataChunk
func NewDataChunk() *DataChunk {
	return &DataChunk{
		Header: NewMetadata().SetType(int32(DataChunkType_DATA_CHUNK_TYPE_DATA)),
	}
}

// ensureTransportMetadata
func (dc *DataChunk) ensureTransportMetadata() *DataChunk {
	if dc.TransportMetadataOptional == nil {
		dc.TransportMetadataOptional = new(DataChunk_TransportMetadata)
	}
	if dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata == nil {
		dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata = new(Metadata)
	}
	return dc
}

// HasTransportMetadata
func (dc *DataChunk) HasTransportMetadata() bool {
	if dc.TransportMetadataOptional == nil {
		return false
	}
	if dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata == nil {
		return false
	}
	return true
}

// SetTransportMetadata
func (dc *DataChunk) SetTransportMetadata(metadata *Metadata) *DataChunk {
	if dc.TransportMetadataOptional == nil {
		dc.TransportMetadataOptional = new(DataChunk_TransportMetadata)
	}
	dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata = metadata
	return dc
}

// ensurePayloadMetadata
func (dc *DataChunk) ensurePayloadMetadata() *DataChunk {
	if dc.PayloadMetadataOptional == nil {
		dc.PayloadMetadataOptional = new(DataChunk_PayloadMetadata)
	}
	if dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata == nil {
		dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata = new(Metadata)
	}
	return dc
}

// HasPayloadMetadata
func (dc *DataChunk) HasPayloadMetadata() bool {
	if dc.PayloadMetadataOptional == nil {
		return false
	}
	if dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata == nil {
		return false
	}
	return true
}

// SetPayloadMetadata
func (dc *DataChunk) SetPayloadMetadata(metadata *Metadata) *DataChunk {
	if dc.PayloadMetadataOptional == nil {
		dc.PayloadMetadataOptional = new(DataChunk_PayloadMetadata)
	}
	dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata = metadata
	return dc
}

// SetData
func (dc *DataChunk) SetData(data []byte) *DataChunk {
	if dc != nil {
		dc.Data = data
	}
	return dc
}

// SetType
func (dc *DataChunk) SetType(_type DataChunkType) *DataChunk {
	dc.GetHeader().SetType(int32(_type))
	return dc
}

// GetType
func (dc *DataChunk) GetType() DataChunkType {
	return DataChunkType(dc.GetHeader().GetType())
}

// SetOffset
func (dc *DataChunk) SetOffset(offset int64) *DataChunk {
	dc.GetHeader().SetOffset(offset)
	return dc
}

// GetOffset
func (dc *DataChunk) GetOffset() int64 {
	return dc.GetHeader().GetOffset()
}

// SetLast
func (dc *DataChunk) SetLast(last bool) *DataChunk {
	dc.GetHeader().SetLast(last)
	return dc
}

// GetLast
func (dc *DataChunk) GetLast() bool {
	return dc.GetHeader().GetLast()
}
