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

// AddressType represents all types of domain-specific addresses in the system
const (
	// Due to first enum value has to be zero in proto3
	AddressReserved int32 = 0
	// S3 and MinIO address
	AddressS3 int32 = 100
	// Kafka address
	AddressKafka int32 = 200
	// Digest-based address
	AddressDigest int32 = 300
	// UUID-based address
	AddressUUID int32 = 400
	// UserID-based address. Used to specify any related user (owner, sender, etc)
	AddressUserID int32 = 500
	// Dirname/path-based address
	AddressDirname int32 = 600
	// Filename/filepath-based address
	AddressFilename int32 = 700
	// URL address
	AddressURL int32 = 800
	// Domain address
	AddressDomain int32 = 900
	// Custom string
	AddressCustomString int32 = 1000
)

var AddressTypeEnum = NewEnum()

func init() {
	AddressTypeEnum.MustRegister("AddressReserved", AddressReserved)
	AddressTypeEnum.MustRegister("AddressS3", AddressS3)
	AddressTypeEnum.MustRegister("AddressKafka", AddressKafka)
	AddressTypeEnum.MustRegister("AddressDigest", AddressDigest)
	AddressTypeEnum.MustRegister("AddressUUID", AddressUUID)
	AddressTypeEnum.MustRegister("AddressUserID", AddressUserID)
	AddressTypeEnum.MustRegister("AddressDirname", AddressDirname)
	AddressTypeEnum.MustRegister("AddressFilename", AddressFilename)
	AddressTypeEnum.MustRegister("AddressURL", AddressURL)
	AddressTypeEnum.MustRegister("AddressDomain", AddressDomain)
	AddressTypeEnum.MustRegister("AddressCustomString", AddressCustomString)
}

// NewAddress creates new Address with specified domain
// Call example:
// NewAddress(domain, address)
func NewAddress(entities ...interface{}) *Address {
	var domain *Domain = nil
	var address interface{} = nil
	for _, entity := range entities {
		switch typed := entity.(type) {
		case *Domain:
			domain = typed
		default:
			address = entity
		}
	}

	res := new(Address)
	if domain != nil {
		res.SetDomain(domain)
	}
	if address != nil {
		res.Set(address)
	}
	return res
}

// NewAddressUUIDRandom creates new Address with specified Domain with random UUID
func NewAddressUUIDRandom(domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewUUIDRandom())
}

// NewAddressUUIDFromString creates new Address with specified Domain with UUID fetched from string
func NewAddressUUIDFromString(str string, domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewUUID().SetString(str))
}

// Ensure returns new or existing Address
func (m *Address) Ensure() *Address {
	if m == nil {
		return NewAddress()
	}
	return m
}

// SetDomain sets Domain of the Address
func (m *Address) SetDomain(domain *Domain) *Address {
	if m == nil {
		return nil
	}
	explicitDomain := new(Address_ExplicitDomain)
	explicitDomain.ExplicitDomain = domain
	m.DomainOptional = explicitDomain
	return m
}

// GetAddressDomain
func (m *Address) GetAddressDomain() *Domain {
	if m == nil {
		return nil
	}
	if explicit := m.GetExplicitDomain(); explicit != nil {
		return explicit
	}

	switch {
	case m.GetS3() != nil:
		return DomainS3
	case m.GetKafka() != nil:
		return DomainKafka
	case m.GetDigest() != nil:
		return DomainDigest
	case m.GetUUID() != nil:
		return DomainUUID
	case m.GetUserId() != nil:
		return DomainUserID
	case m.GetDirname() != nil:
		return DomainDirname
	case m.GetFilename() != nil:
		return DomainFilename
	case m.GetUrl() != nil:
		return DomainURL
	case m.GetDomain() != nil:
		return DomainDomain
	default:
		return DomainCustomString
	}
}

// Set sets value of the Address
func (m *Address) Set(address interface{}) *Address {
	if m == nil {
		return nil
	}
	switch typed := address.(type) {
	case isAddress_AddressOptional:
		m.AddressOptional = typed
	case *S3Address:
		i := new(Address_S3)
		i.S3 = typed
		m.AddressOptional = i
		m.SetDomain(DomainS3)
	case *KafkaAddress:
		i := new(Address_Kafka)
		i.Kafka = typed
		m.AddressOptional = i
		m.SetDomain(DomainKafka)
	case *Digest:
		i := new(Address_Digest)
		i.Digest = typed
		m.AddressOptional = i
		m.SetDomain(DomainDigest)
	case *UUID:
		i := new(Address_Uuid)
		i.Uuid = typed
		m.AddressOptional = i
		m.SetDomain(DomainUUID)
	case *UserID:
		i := new(Address_UserId)
		i.UserId = typed
		m.AddressOptional = i
		m.SetDomain(DomainUserID)
	case *Dirname:
		i := new(Address_Dirname)
		i.Dirname = typed
		m.AddressOptional = i
		m.SetDomain(DomainDirname)
	case *Filename:
		i := new(Address_Filename)
		i.Filename = typed
		m.AddressOptional = i
		m.SetDomain(DomainFilename)
	case *URL:
		i := new(Address_Url)
		i.Url = typed
		m.AddressOptional = i
		m.SetDomain(DomainURL)
	case *Domain:
		i := new(Address_Domain)
		i.Domain = typed
		m.AddressOptional = i
		m.SetDomain(DomainDomain)
	case string:
		i := new(Address_CustomString)
		i.CustomString = typed
		m.AddressOptional = i
		m.SetDomain(DomainCustomString)
	}
	return m
}

// String
func (m *Address) String() string {
	if m == nil {
		return ""
	}

	if m.GetAddressOptional() == nil {
		return "address unspecified"
	}

	switch {
	case m.GetS3() != nil:
		return m.GetS3().String()
	case m.GetKafka() != nil:
		return m.GetKafka().String()
	case m.GetDigest() != nil:
		return m.GetDigest().String()
	case m.GetUUID() != nil:
		return m.GetUUID().String()
	case m.GetUserId() != nil:
		return m.GetUserId().String()
	case m.GetDirname() != nil:
		return m.GetDirname().String()
	case m.GetFilename() != nil:
		return m.GetFilename().String()
	case m.GetUrl() != nil:
		return m.GetUrl().String()
	case m.GetDomain() != nil:
		return m.GetDomain().String()
	default:
		return m.GetCustomString()
	}
}
