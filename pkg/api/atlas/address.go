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
	// Filename/filepath-based address
	AddressFilename int32 = 600
	// URL address
	AddressURL int32 = 700
	// Domain address
	AddressDomain int32 = 800
	// Custom string
	AddressCustomString int32 = 900
)

var AddressTypeEnum = NewEnum()

func init() {
	AddressTypeEnum.MustRegister("AddressReserved", AddressReserved)
	AddressTypeEnum.MustRegister("AddressS3", AddressS3)
	AddressTypeEnum.MustRegister("AddressKafka", AddressKafka)
	AddressTypeEnum.MustRegister("AddressDigest", AddressDigest)
	AddressTypeEnum.MustRegister("AddressUUID", AddressUUID)
	AddressTypeEnum.MustRegister("AddressUserID", AddressUserID)
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
	return NewAddress(domain...).Set(NewRandomUUID())
}

// NewAddressUUIDFromString creates new Address with specified Domain with UUID fetched from string
func NewAddressUUIDFromString(str string, domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewUUID().SetString(str))
}

// SetDomain sets Domain of the Address
func (m *Address) SetDomain(domain *Domain) *Address {
	if m == nil {
		return nil
	}
	addressDomain := new(Address_AddressDomain)
	addressDomain.AddressDomain = domain
	m.DomainOptional = addressDomain
	return m
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
	case *KafkaAddress:
		i := new(Address_Kafka)
		i.Kafka = typed
		m.AddressOptional = i
	case *Digest:
		i := new(Address_Digest)
		i.Digest = typed
		m.AddressOptional = i
	case *UUID:
		i := new(Address_Uuid)
		i.Uuid = typed
		m.AddressOptional = i
	case *UserID:
		i := new(Address_UserId)
		i.UserId = typed
		m.AddressOptional = i
	case *Filename:
		i := new(Address_Filename)
		i.Filename = typed
		m.AddressOptional = i
	case *URL:
		i := new(Address_Url)
		i.Url = typed
		m.AddressOptional = i
	case *Domain:
		i := new(Address_Domain)
		i.Domain = typed
		m.AddressOptional = i
	case string:
		i := new(Address_CustomString)
		i.CustomString = typed
		m.AddressOptional = i
	}
	return m
}

// Printable
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
		return "digest printable not implemented"
	case m.GetUuid() != nil:
		return "uuid printable not implemented"
	case m.GetUserId() != nil:
		return "userid printable not implemented"
	case m.GetFilename() != nil:
		return m.GetFilename().String()
	case m.GetUrl() != nil:
		return m.GetUrl().String()
	case m.GetDomain() != nil:
		return "domain printable not implemented"
	default:
		return m.GetCustomString()
	}
}
