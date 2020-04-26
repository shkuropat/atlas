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

package mservice

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

func NewHeader(
	Type int32,
	name string,
	version int32,
	uuid string,
	uuidReference string,
	seconds int64,
	nanoSeconds int32,
	description string,
) *Header {
	h := new(Header)

	h.SetTypeOrName(Type, name)

	if version > 0 {
		h.SetVersion(version)
	}

	if uuid == "" {
		h.SetUUID(CreateNewUUID())
	} else {
		h.SetUUID(uuid)
	}

	if uuidReference != "" {
		h.SetUUIDReference(uuidReference)
	}

	if seconds > 0 {
		h.SetTimestamp(seconds, nanoSeconds)
	}

	if description != "" {
		h.SetDescription(description)
	}

	return h
}

func (h *Header) ensureTypeName() {
	if h.TypeName == nil {
		h.TypeName = new(TypeName)
	}
}

func (h *Header) SetTypeOrName(Type int32, name string) {
	if Type > 0 {
		h.SetType(Type)
	} else {
		h.SetName(name)
	}
}

func (h *Header) SetType(Type int32) {
	h.ensureTypeName()
	if h.TypeName.TypeOptional == nil {
		h.TypeName.TypeOptional = new(TypeName_Type)
	}
	h.TypeName.TypeOptional.(*TypeName_Type).Type = Type
}

func (h *Header) SetName(name string) {
	h.ensureTypeName()
	if h.TypeName.NameOptional == nil {
		h.TypeName.NameOptional = new(TypeName_Name)
	}
	h.TypeName.NameOptional.(*TypeName_Name).Name = name
}

func (h *Header) SetVersion(version int32) {
	if h.VersionOptional == nil {
		h.VersionOptional = new(Header_Version)
		h.VersionOptional.(*Header_Version).Version = version
	}
}

func (h *Header) SetUUID(uuid string) {
	if h.UuidOptional == nil {
		h.UuidOptional = new(Header_Uuid)
		if h.UuidOptional.(*Header_Uuid).Uuid == nil {
			h.UuidOptional.(*Header_Uuid).Uuid = NewUUID(uuid)
		}

	}
}

func (h *Header) SetUUIDReference(uuid string) {
	if h.UuidReferenceOptional == nil {
		h.UuidReferenceOptional = new(Header_UuidReference)
		if h.UuidReferenceOptional.(*Header_UuidReference).UuidReference == nil {
			h.UuidReferenceOptional.(*Header_UuidReference).UuidReference = NewUUID(uuid)
		}
	}
}

func (h *Header) SetTimestamp(seconds int64, nanos int32) {
	if h.TimestampOptional == nil {
		h.TimestampOptional = new(Header_Ts)
		if h.TimestampOptional.(*Header_Ts).Ts == nil {
			h.TimestampOptional.(*Header_Ts).Ts = new(timestamp.Timestamp)
			h.TimestampOptional.(*Header_Ts).Ts.Seconds = seconds
			h.TimestampOptional.(*Header_Ts).Ts.Nanos = nanos
		}
	}
}

func (h *Header) SetDescription(description string) {
	if h.DescriptionOptional == nil {
		h.DescriptionOptional = new(Header_Description)
		h.DescriptionOptional.(*Header_Description).Description = description
	}
}
