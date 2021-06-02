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
func (m *Metadata) SetType(_type int32) *Metadata {
	if m.TypeOptional == nil {
		m.TypeOptional = new(Metadata_Type)
	}
	m.TypeOptional.(*Metadata_Type).Type = _type
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

// HasStatus
func (m *Metadata) HasStatus() bool {
	if m == nil {
		return false
	}
	return m.StatusOptional != nil
}

// SetStatus
func (m *Metadata) SetStatus(status int32) *Metadata {
	if m.StatusOptional == nil {
		m.StatusOptional = new(Metadata_Status)
	}
	m.StatusOptional.(*Metadata_Status).Status = status
	return m
}

// HasMode
func (m *Metadata) HasMode() bool {
	if m == nil {
		return false
	}
	return m.ModeOptional != nil
}

// SetMode
func (m *Metadata) SetMode(mode int32) *Metadata {
	if m.ModeOptional == nil {
		m.ModeOptional = new(Metadata_Mode)
	}
	m.ModeOptional.(*Metadata_Mode).Mode = mode
	return m
}

// HasTimestamp
func (m *Metadata) HasTimestamp() bool {
	if m == nil {
		return false
	}
	return m.TsOptional != nil
}

// SetTimestamp
func (m *Metadata) SetTimestamp(seconds int64, nanos int32) *Metadata {
	if m.TsOptional == nil {
		m.TsOptional = new(Metadata_Ts)
	}
	m.TsOptional.(*Metadata_Ts).Ts = new(timestamp.Timestamp)
	m.TsOptional.(*Metadata_Ts).Ts.Seconds = seconds
	m.TsOptional.(*Metadata_Ts).Ts.Nanos = nanos
	return m
}

// CreateTimestamp creates current timestamp
func (m *Metadata) CreateTimestamp() *Metadata {
	now := time.Now()
	seconds := now.Unix()           // seconds since 1970
	nanoseconds := now.Nanosecond() // nanosecond offset within the second
	return m.SetTimestamp(seconds, int32(nanoseconds))
}

// HasAddresses
func (m *Metadata) HasAddresses() bool {
	if m == nil {
		return false
	}
	return m.AddressesOptional != nil
}

// SetAddresses
func (m *Metadata) SetAddresses(addresses *AddressMap) *Metadata {
	if m == nil {
		return nil
	}
	if m.AddressesOptional == nil {
		m.AddressesOptional = new(Metadata_Addresses)
	}
	m.AddressesOptional.(*Metadata_Addresses).Addresses = addresses
	return m
}

// EnsureAddresses
func (m *Metadata) EnsureAddresses() *AddressMap {
	if m.HasAddresses() {
		return m.GetAddresses()
	}
	m.SetAddresses(NewAddressMap())
	return m.GetAddresses()
}

// HasPresentationOptions
func (m *Metadata) HasPresentationOptions() bool {
	if m == nil {
		return false
	}
	return m.PresentationOptionsOptional != nil
}

// SetPresentationOptions
func (m *Metadata) SetPresentationOptions(options *PresentationOptions) *Metadata {
	if m == nil {
		return nil
	}
	if m.PresentationOptionsOptional == nil {
		m.PresentationOptionsOptional = new(Metadata_PresentationOptions)
	}
	m.PresentationOptionsOptional.(*Metadata_PresentationOptions).PresentationOptions = options
	return m
}

// EnsurePresentationOptions
func (m *Metadata) EnsurePresentationOptions() *PresentationOptions {
	if m.HasPresentationOptions() {
		return m.GetPresentationOptions()
	}
	m.SetPresentationOptions(NewPresentationOptions())
	return m.GetPresentationOptions()
}

// HasDataChunkProperties
func (m *Metadata) HasDataChunkProperties() bool {
	if m == nil {
		return false
	}
	return m.DataChunkPropertiesOptional != nil
}

// SetDataChunkProperties
func (m *Metadata) SetDataChunkProperties(properties *DataChunkProperties) *Metadata {
	if m == nil {
		return nil
	}
	if m.DataChunkPropertiesOptional == nil {
		m.DataChunkPropertiesOptional = new(Metadata_DataChunkProperties)
	}
	m.DataChunkPropertiesOptional.(*Metadata_DataChunkProperties).DataChunkProperties = properties
	return m
}

// EnsureDataChunkProperties
func (m *Metadata) EnsureDataChunkProperties() *DataChunkProperties {
	if m.HasDataChunkProperties() {
		return m.GetDataChunkProperties()
	}
	m.SetDataChunkProperties(NewDataChunkProperties())
	return m.GetDataChunkProperties()
}

// Log
func (m *Metadata) Log() {
	if m == nil {
		return
	}
	log.Infof("metadata: %s", m.String())
}

// String
func (m *Metadata) String() string {
	return "to be implemented"
}
