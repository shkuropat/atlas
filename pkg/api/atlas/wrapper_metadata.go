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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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

// HasUserID
func (m *Metadata) HasUserID() bool {
	if m == nil {
		return false
	}
	return m.UserIdOptional != nil
}

// SetUserID
func (m *Metadata) SetUserID(id *UserID) *Metadata {
	if m.UserIdOptional == nil {
		m.UserIdOptional = new(Metadata_UserId)
	}
	m.UserIdOptional.(*Metadata_UserId).UserId = id

	return m
}

// SetUserIDFromString
func (m *Metadata) SetUserIDFromString(id string) *Metadata {
	return m.SetUserID(NewUserID().SetString(id))
}

// HasID
func (m *Metadata) HasID() bool {
	if m == nil {
		return false
	}
	return m.IdOptional != nil
}

// SetID
func (m *Metadata) SetID(id *UUID) *Metadata {
	if m.IdOptional == nil {
		m.IdOptional = new(Metadata_Id)
	}
	m.IdOptional.(*Metadata_Id).Id = id

	return m
}

// SetIDFromString
func (m *Metadata) SetIDFromString(id string) *Metadata {
	return m.SetID(NewUUID().SetString(id))
}

// CreateID
func (m *Metadata) CreateID() *Metadata {
	return m.SetID(CreateUUID())
}

// HasReferenceID
func (m *Metadata) HasReferenceID() bool {
	if m == nil {
		return false
	}
	return m.ReferenceIdOptional != nil
}

// SetReferenceID
func (m *Metadata) SetReferenceID(id *UUID) *Metadata {
	if m.ReferenceIdOptional == nil {
		m.ReferenceIdOptional = new(Metadata_ReferenceId)
	}
	m.ReferenceIdOptional.(*Metadata_ReferenceId).ReferenceId = id

	return m
}

// SetReferenceIDFromString
func (m *Metadata) SetReferenceIDFromString(id string) *Metadata {
	return m.SetReferenceID(NewUUID().SetString(id))
}

// HasTimestamp
func (m *Metadata) HasTimestamp() bool {
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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

// HasDomain
func (m *Metadata) HasDomain() bool {
	if m == nil {
		return false
	}
	return m.DomainOptional != nil
}

// SetDomain
func (m *Metadata) SetDomain(domain *Domain) *Metadata {
	if domain == nil {
		return m
	}

	if m.DomainOptional == nil {
		m.DomainOptional = new(Metadata_Domain)
	}
	m.DomainOptional.(*Metadata_Domain).Domain = domain

	return m
}

// HasDigest
func (m *Metadata) HasDigest() bool {
	if m == nil {
		return false
	}
	return m.DigestOptional != nil
}

// SetDigest
func (m *Metadata) SetDigest(digest *Digest) *Metadata {
	if digest == nil {
		return m
	}

	if m.DigestOptional == nil {
		m.DigestOptional = new(Metadata_Digest)
	}
	m.DigestOptional.(*Metadata_Digest).Digest = digest

	return m
}

// HasDescription
func (m *Metadata) HasDescription() bool {
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return false
	}
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
	if m == nil {
		return
	}
	log.Infof("metadata: %s", m.String())
}
