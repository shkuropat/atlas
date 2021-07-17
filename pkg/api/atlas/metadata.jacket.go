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

//
//
// Wrap AddressMap
//
//

// Has wraps AddressMap.Has
func (m *Metadata) Has(domain ...*Domain) bool {
	return m.GetAddresses().Has(domain...)
}

// Set wraps AddressMap.Set
func (m *Metadata) Set(entities ...interface{}) *Metadata {
	m.EnsureAddresses().Set(entities...)
	return m
}

// Append wraps AddressMap.Append
func (m *Metadata) Append(entities ...interface{}) *Metadata {
	m.EnsureAddresses().Append(entities...)
	return m
}

//
//
// Address customization
//
//
// SetUUIDFromString
func (m *Metadata) SetUUIDFromString(id string, domain ...*Domain) *Metadata {
	var domain0 = DomainThis
	var domain1 = DomainUUID
	switch len(domain) {
	case 2:
		domain0 = domain[0]
		domain1 = domain[1]
	case 1:
		domain0 = domain[0]
	}
	m.Set(domain0, domain1, NewAddressUUIDFromString(id, domain1))
	return m
}

// SetRandomUUID domains are optional
func (m *Metadata) SetRandomUUID(domains ...*Domain) *Metadata {
	// Default values
	var domain0 = DomainThis
	var domain1 = DomainUUID

	switch len(domains) {
	case 2:
		domain0 = domains[0]
		domain1 = domains[1]
	case 1:
		domain0 = domains[0]
	}
	m.Set(domain0, domain1, NewAddressUUIDRandom(domain1))
	return m
}

//
//
// Wrap Address
//
//

// SetS3
func (m *Metadata) SetS3(bucket, object string) *Metadata {
	return m.Set(DomainThis, DomainS3, NewAddress().SetDomain(DomainS3).Set(NewS3Address(bucket, object)))
}

// SetKafka
func (m *Metadata) SetKafka(topic string, partition int32) *Metadata {
	return m.Set(DomainThis, DomainKafka, NewAddress().SetDomain(DomainKafka).Set(NewKafkaAddress(topic, partition)))
}

// SetUUID
func (m *Metadata) SetUUID(uuid *UUID) *Metadata {
	return m.Set(DomainThis, DomainUUID, NewAddress().SetDomain(DomainUUID).Set(uuid))
}

// GetUUID
func (m *Metadata) GetUUID() *UUID {
	return m.GetAddresses().First(DomainThis, DomainUUID).GetUUID()
}

// SetUserID
func (m *Metadata) SetUserID(userID *UserID) *Metadata {
	return m.Set(DomainThis, DomainUserID, NewAddress().SetDomain(DomainUserID).Set(userID))
}

// GetUserID
func (m *Metadata) GetUserID() *UserID {
	return m.GetAddresses().First(DomainThis, DomainUserID).GetUserId()
}

// SetDirname
func (m *Metadata) SetDirname(dirname string) *Metadata {
	return m.Set(DomainThis, DomainDirname, NewAddress().SetDomain(DomainDirname).Set(NewDirname(dirname)))
}

// GetDirname
func (m *Metadata) GetDirname() string {
	return m.GetAddresses().First(DomainThis, DomainDirname).GetDirname().String()
}

// SetFilename
func (m *Metadata) SetFilename(filename string) *Metadata {
	return m.Set(DomainThis, DomainFilename, NewAddress().SetDomain(DomainFilename).Set(NewFilename(filename)))
}

// GetFilename
func (m *Metadata) GetFilename() string {
	return m.GetAddresses().First(DomainThis, DomainFilename).GetFilename().String()
}

// SetURL
func (m *Metadata) SetURL(url string) *Metadata {
	return m.Set(DomainThis, DomainURL, NewAddress().SetDomain(DomainURL).Set(NewURL(url)))
}

// SetDomain
func (m *Metadata) SetDomain(domain *Domain) *Metadata {
	return m.Set(DomainThis, DomainDomain, NewAddress().SetDomain(DomainDomain).Set(domain))
}

// GetDomain
func (m *Metadata) GetDomain() *Domain {
	return m.GetAddresses().First(DomainThis, DomainDomain).GetDomain()
}

// SetCustomString
func (m *Metadata) SetCustom(s string) *Metadata {
	return m.Set(DomainThis, DomainCustom, NewAddress().SetDomain(DomainCustom).Set(s))
}

// GetContextUUID
func (m *Metadata) GetContextUUID() *UUID {
	return m.GetAddresses().First(DomainContext, DomainUUID).GetUUID()
}

// SetContextUUID
func (m *Metadata) SetContextUUID(uuid *UUID) *Metadata {
	return m.Set(DomainContext, DomainUUID, NewAddress().SetDomain(DomainUUID).Set(uuid))
}

// GetTaskUUID
func (m *Metadata) GetTaskUUID() *UUID {
	return m.GetAddresses().First(DomainTask, DomainUUID).GetUUID()
}

// SetTaskUUID
func (m *Metadata) SetTaskUUID(uuid *UUID) *Metadata {
	return m.Set(DomainTask, DomainUUID, NewAddress().SetDomain(DomainUUID).Set(uuid))
}

// GetResultDomain
func (m *Metadata) GetResultDomain() *Domain {
	return m.GetAddresses().First(DomainResult, DomainDomain).GetDomain()
}

// SetResultDomain
func (m *Metadata) SetResultDomain(domain *Domain) *Metadata {
	return m.Set(DomainResult, DomainDomain, NewAddress().SetDomain(DomainDomain).Set(domain))
}

//
//
// Wrap PresentationOptions
//
//

// HasDigest
func (m *Metadata) HasDigest() bool {
	return m.GetPresentationOptions().HasDigest()
}

// GetDigest
func (m *Metadata) GetDigest() *Digest {
	return m.GetPresentationOptions().GetDigest()
}

// SetDigest
func (m *Metadata) SetDigest(digest *Digest) *Metadata {
	m.EnsurePresentationOptions().SetDigest(digest)
	return m
}

// HasCompression
func (m *Metadata) HasCompression() bool {
	return m.GetPresentationOptions().HasCompression()
}

// GetCompression
func (m *Metadata) GetCompression() *Compression {
	return m.GetPresentationOptions().GetCompression()
}

// SetCompression
func (m *Metadata) SetCompression(compression *Compression) *Metadata {
	m.EnsurePresentationOptions().SetCompression(compression)
	return m
}

//
//
// Wrap DataChunkProperties
//
//

// HasLen
func (m *Metadata) HasLen() bool {
	return m.GetDataChunkProperties().HasLen()
}

// GetLen
func (m *Metadata) GetLen() int64 {
	return m.GetDataChunkProperties().GetLen()
}

// SetLen
func (m *Metadata) SetLen(len int64) *Metadata {
	m.EnsureDataChunkProperties().SetLen(len)
	return m
}

// HasOffset
func (m *Metadata) HasOffset() bool {
	return m.GetDataChunkProperties().HasOffset()
}

// GetOffset
func (m *Metadata) GetOffset() int64 {
	return m.GetDataChunkProperties().GetOffset()
}

// SetOffset
func (m *Metadata) SetOffset(offset int64) *Metadata {
	m.EnsureDataChunkProperties().SetOffset(offset)
	return m
}

// HasLast
func (m *Metadata) HasLast() bool {
	return m.GetDataChunkProperties().HasLast()
}

// GetLast
func (m *Metadata) GetLast() bool {
	return m.GetDataChunkProperties().GetLast()
}

// SetLast
func (m *Metadata) SetLast(last bool) *Metadata {
	m.EnsureDataChunkProperties().SetLast(last)
	return m
}
