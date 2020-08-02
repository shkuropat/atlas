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
	"time"
)

// NewMetadata
func NewMetadata() *Metadata {
	return new(Metadata)
}

// HasType
func (m *Metadata) HasType() bool {
	return m.TypeOptional != nil
}

// SetType
func (m *Metadata) SetType(_type MetadataType) *Metadata {
	if m.TypeOptional == nil {
		m.TypeOptional = new(Metadata_Type)
	}
	m.TypeOptional.(*Metadata_Type).Type = int32(_type)

	return m
}

// HasName
func (m *Metadata) HasName() bool {
	return m.NameOptional != nil
}

// SetName
func (m *Metadata) SetName(name string) *Metadata {
	if m.NameOptional == nil {
		m.NameOptional = new(Metadata_Name)
	}
	m.NameOptional.(*Metadata_Name).Name = name

	return m
}

// HasVersion
func (m *Metadata) HasVersion() bool {
	return m.VersionOptional != nil
}

// SetVersion
func (m *Metadata) SetVersion(version int32) *Metadata {
	if m.VersionOptional == nil {
		m.VersionOptional = new(Metadata_Version)
	}
	m.VersionOptional.(*Metadata_Version).Version = version

	return m
}

// HasUUID
func (m *Metadata) HasUUID() bool {
	return m.UuidOptional != nil
}

// SetUUID
func (m *Metadata) SetUUID(uuid string) *Metadata {
	if m.UuidOptional == nil {
		m.UuidOptional = new(Metadata_Uuid)
	}
	m.UuidOptional.(*Metadata_Uuid).Uuid = NewUUID(uuid)

	return m
}

// CreateUUID
func (m *Metadata) CreateUUID() *Metadata {
	return m.SetUUID(CreateNewUUID())
}

// HasUUIDReference
func (m *Metadata) HasUUIDReference() bool {
	return m.UuidReferenceOptional != nil
}

// SetUUIDReference
func (m *Metadata) SetUUIDReference(uuid string) *Metadata {
	if m.UuidReferenceOptional == nil {
		m.UuidReferenceOptional = new(Metadata_UuidReference)
	}
	m.UuidReferenceOptional.(*Metadata_UuidReference).UuidReference = NewUUID(uuid)

	return m
}

// HasTimestamp
func (m *Metadata) HasTimestamp() bool {
	return m.TimestampOptional != nil
}

// SetTimestamp
func (m *Metadata) SetTimestamp(seconds int64, nanos int32) *Metadata {
	if m.TimestampOptional == nil {
		m.TimestampOptional = new(Metadata_Ts)
	}
	m.TimestampOptional.(*Metadata_Ts).Ts = new(timestamp.Timestamp)
	m.TimestampOptional.(*Metadata_Ts).Ts.Seconds = seconds
	m.TimestampOptional.(*Metadata_Ts).Ts.Nanos = nanos

	return m
}

// CreateTimestamp
func (m *Metadata) CreateTimestamp() *Metadata {
	now := time.Now()
	seconds := now.Unix()           // seconds since 1970
	nanoseconds := now.Nanosecond() // nanosecond offset within the second

	return m.SetTimestamp(seconds, int32(nanoseconds))
}

// HasEncoding
func (m *Metadata) HasEncoding() bool {
	return m.EncodingOptional != nil
}

// SetEncoding
func (m *Metadata) SetEncoding(encoding string) *Metadata {
	if encoding == "" {
		return m
	}

	if m.EncodingOptional == nil {
		m.EncodingOptional = new(Metadata_Encoding)
	}
	m.EncodingOptional.(*Metadata_Encoding).Encoding = encoding

	return m
}

// HasCompression
func (m *Metadata) HasCompression() bool {
	return m.CompressionOptional != nil
}

// SetCompression
func (m *Metadata) SetCompression(compression string) *Metadata {
	if m.CompressionOptional == nil {
		m.CompressionOptional = new(Metadata_Compression)
	}
	m.CompressionOptional.(*Metadata_Compression).Compression = compression

	return m
}

// HasFilename
func (m *Metadata) HasFilename() bool {
	return m.FilenameOptional != nil
}

// SetFilename
func (m *Metadata) SetFilename(filename string) *Metadata {
	if filename == "" {
		return m
	}

	if m.FilenameOptional == nil {
		m.FilenameOptional = new(Metadata_Filename)
	}
	m.FilenameOptional.(*Metadata_Filename).Filename = filename

	return m
}

// HasURL
func (m *Metadata) HasURL() bool {
	return m.UrlOptional != nil
}

// SetURL
func (m *Metadata) SetURL(url string) *Metadata {
	if url == "" {
		return m
	}

	if m.UrlOptional == nil {
		m.UrlOptional = new(Metadata_Url)
	}
	m.UrlOptional.(*Metadata_Url).Url = url

	return m
}

// HasS3Address
func (m *Metadata) HasS3Address() bool {
	return m.S3AddressOptional != nil
}

// SetS3Address
func (m *Metadata) SetS3Address(s3address *S3Address) *Metadata {
	if s3address == nil {
		return m
	}

	if m.S3AddressOptional == nil {
		m.S3AddressOptional = new(Metadata_S3Address)
	}
	m.S3AddressOptional.(*Metadata_S3Address).S3Address = s3address

	return m
}

// HasDescription
func (m *Metadata) HasDescription() bool {
	return m.DescriptionOptional != nil
}

// SetDescription
func (m *Metadata) SetDescription(description string) *Metadata {
	if m.DescriptionOptional == nil {
		m.DescriptionOptional = new(Metadata_Description)
	}
	m.DescriptionOptional.(*Metadata_Description).Description = description

	return m
}

// HasLen
func (m *Metadata) HasLen() bool {
	return m.LenOptional != nil
}

// SetLen
func (m *Metadata) SetLen(len int64) *Metadata {
	if m.LenOptional == nil {
		m.LenOptional = new(Metadata_Len)
	}
	m.LenOptional.(*Metadata_Len).Len = len

	return m
}

// HasOffset
func (m *Metadata) HasOffset() bool {
	return m.OffsetOptional != nil
}

// SetOffset
func (m *Metadata) SetOffset(offset int64) *Metadata {
	if m.OffsetOptional == nil {
		m.OffsetOptional = new(Metadata_Offset)
	}
	m.OffsetOptional.(*Metadata_Offset).Offset = offset

	return m
}

// HasLast
func (m *Metadata) HasLast() bool {
	return m.LastOptional != nil
}

// SetLast
func (m *Metadata) SetLast(last bool) *Metadata {
	if m.LastOptional == nil {
		m.LastOptional = new(Metadata_Last)
	}
	m.LastOptional.(*Metadata_Last).Last = last

	return m
}

// Log
func (m *Metadata) Log() {
	log.Infof("metadata: %s", m.String())
}
