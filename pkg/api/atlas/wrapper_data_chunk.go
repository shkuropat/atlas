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

// NewDataChunk
func NewDataChunk(transportMD, payloadMD *Metadata, offset *int64, last bool, data []byte) *DataChunk {
	dc := &DataChunk{
		Header: NewMetadata().SetType(MetadataType(DataChunkType_DATA_CHUNK_TYPE_DATA)),
		Bytes:  data,
	}

	if transportMD != nil {
		dc.SetTransportMetadata(transportMD)
	}

	if payloadMD != nil {
		dc.SetPayloadMetadata(payloadMD)
	}

	if offset != nil {
		dc.Header.SetOffset(*offset)
	}

	if last {
		dc.Header.SetLast(last)
	}

	return dc
}

// ensureTransportMetadata
func (dc *DataChunk) ensureTransportMetadata() {
	if dc.TransportMetadataOptional == nil {
		dc.TransportMetadataOptional = new(DataChunk_TransportMetadata)
	}
	if dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata == nil {
		dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata = new(Metadata)
	}
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
func (dc *DataChunk) SetTransportMetadata(metadata *Metadata) {
	if dc.TransportMetadataOptional == nil {
		dc.TransportMetadataOptional = new(DataChunk_TransportMetadata)
	}
	dc.TransportMetadataOptional.(*DataChunk_TransportMetadata).TransportMetadata = metadata
}

// ensurePayloadMetadata
func (dc *DataChunk) ensurePayloadMetadata() {
	if dc.PayloadMetadataOptional == nil {
		dc.PayloadMetadataOptional = new(DataChunk_PayloadMetadata)
	}
	if dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata == nil {
		dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata = new(Metadata)
	}
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
func (dc *DataChunk) SetPayloadMetadata(metadata *Metadata) {
	if dc.PayloadMetadataOptional == nil {
		dc.PayloadMetadataOptional = new(DataChunk_PayloadMetadata)
	}
	dc.PayloadMetadataOptional.(*DataChunk_PayloadMetadata).PayloadMetadata = metadata
}

func (dc *DataChunk) GetOffset() int64 {
	return dc.GetHeader().GetOffset()
}
