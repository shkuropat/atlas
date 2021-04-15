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

// NewStatusRequestMulti
func NewStatusRequestMulti() *StatusRequestMulti {
	return new(StatusRequestMulti)
}

// EnsureDomain
func (m *StatusRequestMulti) EnsureDomain() *Domain {
	if m == nil {
		return nil
	}
	if m.DomainOptional == nil {
		m.DomainOptional = new(StatusRequestMulti_Domain)
	}
	if m.DomainOptional.(*StatusRequestMulti_Domain).Domain == nil {
		m.DomainOptional.(*StatusRequestMulti_Domain).Domain = NewDomain()
	}
	return m.GetDomain()
}

// SetDomain
func (m *StatusRequestMulti) SetDomain(domain *Domain) *StatusRequestMulti {
	if m == nil {
		return nil
	}
	if m.DomainOptional == nil {
		m.DomainOptional = new(StatusRequestMulti_Domain)
	}
	m.DomainOptional.(*StatusRequestMulti_Domain).Domain = domain
	return m
}

// EnsureMode
func (m *StatusRequestMulti) EnsureMode() StatusRequestMode {
	if m == nil {
		return StatusRequestMode_RESERVED
	}
	if m.ModeOptional == nil {
		m.ModeOptional = new(StatusRequestMulti_Mode)
	}
	return m.GetMode()
}

// SetMode
func (m *StatusRequestMulti) SetMode(mode StatusRequestMode) *StatusRequestMulti {
	if m == nil {
		return nil
	}
	if m.ModeOptional == nil {
		m.ModeOptional = new(StatusRequestMulti_Mode)
	}
	m.ModeOptional.(*StatusRequestMulti_Mode).Mode = mode
	return m
}

// GetEntitiesNum
func (m *StatusRequestMulti) GetEntitiesNum() int {
	return len(m.GetEntities())
}

// String
func (m *StatusRequestMulti) String() string {
	return "to be implemented"
}
