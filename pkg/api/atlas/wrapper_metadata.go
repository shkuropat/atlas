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

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
)

// NewMetadata
func NewMetadata(
	_type int32,
	name string,
	version int32,
	uuid string,
	uuidReference string,
	seconds int64,
	nanoSeconds int32,
	description string,
) *Metadata {
	md := new(Metadata)

	if _type > 0 {
		md.SetType(_type)
	}

	if name != "" {
		md.SetName(name)
	}

	if version > 0 {
		md.SetVersion(version)
	}

	if uuid == "" {
		md.SetUUID(CreateNewUUID())
	} else {
		md.SetUUID(uuid)
	}

	if uuidReference != "" {
		md.SetUUIDReference(uuidReference)
	}

	if seconds > 0 {
		md.SetTimestamp(seconds, nanoSeconds)
	}

	if description != "" {
		md.SetDescription(description)
	}

	return md
}

// HasType
func (m *Metadata) HasType() bool {
	return m.TypeOptional != nil
}

// SetType
func (m *Metadata) SetType(_type int32) {
	if m.TypeOptional == nil {
		m.TypeOptional = new(Metadata_Type)
	}
	m.TypeOptional.(*Metadata_Type).Type = _type
}

// HasName
func (m *Metadata) HasName() bool {
	return m.NameOptional != nil
}

// SetName
func (m *Metadata) SetName(name string) {
	if m.NameOptional == nil {
		m.NameOptional = new(Metadata_Name)
	}
	m.NameOptional.(*Metadata_Name).Name = name
}

// HasVersion
func (m *Metadata) HasVersion() bool {
	return m.VersionOptional != nil
}

// SetVersion
func (m *Metadata) SetVersion(version int32) {
	if m.VersionOptional == nil {
		m.VersionOptional = new(Metadata_Version)
	}
	m.VersionOptional.(*Metadata_Version).Version = version
}

// HasUUID
func (m *Metadata) HasUUID() bool {
	return m.UuidOptional != nil
}

// SetUUID
func (m *Metadata) SetUUID(uuid string) {
	if m.UuidOptional == nil {
		m.UuidOptional = new(Metadata_Uuid)
	}
	m.UuidOptional.(*Metadata_Uuid).Uuid = NewUUID(uuid)
}

// HasUUIDReference
func (m *Metadata) HasUUIDReference() bool {
	return m.UuidReferenceOptional != nil
}

// SetUUIDReference
func (m *Metadata) SetUUIDReference(uuid string) {
	if m.UuidReferenceOptional == nil {
		m.UuidReferenceOptional = new(Metadata_UuidReference)
	}
	m.UuidReferenceOptional.(*Metadata_UuidReference).UuidReference = NewUUID(uuid)
}

// HasTimestamp
func (m *Metadata) HasTimestamp() bool {
	return m.TimestampOptional != nil
}

// SetTimestamp
func (m *Metadata) SetTimestamp(seconds int64, nanos int32) {
	if m.TimestampOptional == nil {
		m.TimestampOptional = new(Metadata_Ts)
	}
	m.TimestampOptional.(*Metadata_Ts).Ts = new(timestamp.Timestamp)
	m.TimestampOptional.(*Metadata_Ts).Ts.Seconds = seconds
	m.TimestampOptional.(*Metadata_Ts).Ts.Nanos = nanos
}

// HasEncoding
func (m *Metadata) HasEncoding() bool {
	return m.EncodingOptional != nil
}

// SetEncoding
func (m *Metadata) SetEncoding(encoding string) {
	if encoding == "" {
		return
	}

	if m.EncodingOptional == nil {
		m.EncodingOptional = new(Metadata_Encoding)
	}
	m.EncodingOptional.(*Metadata_Encoding).Encoding = encoding
}

// HasCompression
func (m *Metadata) HasCompression() bool {
	return m.CompressionOptional != nil
}

// SetCompression
func (m *Metadata) SetCompression(compression string) {
	if m.CompressionOptional == nil {
		m.CompressionOptional = new(Metadata_Compression)
	}
	m.CompressionOptional.(*Metadata_Compression).Compression = compression
}

// HasFilename
func (m *Metadata) HasFilename() bool {
	return m.FilenameOptional != nil
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

// HasURL
func (m *Metadata) HasURL() bool {
	return m.UrlOptional != nil
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

// HasDescription
func (m *Metadata) HasDescription() bool {
	return m.DescriptionOptional != nil
}

// SetDescription
func (m *Metadata) SetDescription(description string) {
	if m.DescriptionOptional == nil {
		m.DescriptionOptional = new(Metadata_Description)
	}
	m.DescriptionOptional.(*Metadata_Description).Description = description
}

// HasLen
func (m *Metadata) HasLen() bool {
	return m.LenOptional != nil
}

// SetLen
func (m *Metadata) SetLen(len int64) {
	if m.LenOptional == nil {
		m.LenOptional = new(Metadata_Len)
	}
	m.LenOptional.(*Metadata_Len).Len = len
}

// HasOffset
func (m *Metadata) HasOffset() bool {
	return m.OffsetOptional != nil
}

// SetOffset
func (m *Metadata) SetOffset(offset int64) {
	if m.OffsetOptional == nil {
		m.OffsetOptional = new(Metadata_Offset)
	}
	m.OffsetOptional.(*Metadata_Offset).Offset = offset
}

// HasLast
func (m *Metadata) HasLast() bool {
	return m.LastOptional != nil
}

// SetLast
func (m *Metadata) SetLast(last bool) {
	if m.LastOptional == nil {
		m.LastOptional = new(Metadata_Last)
	}
	m.LastOptional.(*Metadata_Last).Last = last
}

// Log
func (m *Metadata) Log() {
	log.Infof("metadata: %s", m.String())
}
