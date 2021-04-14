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

const (
	ReportTypeReserved    = 0
	ReportTypeUnspecified = 100
)

var ReportTypeEnum = NewEnum()

func init() {
	ReportTypeEnum.MustRegister("ReportTypeReserved", ReportTypeReserved)
	ReportTypeEnum.MustRegister("ReportTypeUnspecified", ReportTypeUnspecified)
}

// NewReport
func NewReport() *Report {
	return new(Report)
}

// EnsureHeader
func (m *Report) EnsureHeader() *Metadata {
	if m.HeaderOptional == nil {
		m.HeaderOptional = new(Report_Header)
	}
	if m.HeaderOptional.(*Report_Header).Header == nil {
		m.HeaderOptional.(*Report_Header).Header = new(Metadata)
	}
	return m.HeaderOptional.(*Report_Header).Header
}

// SetBytes
func (m *Report) SetBytes(bytes []byte) *Report {
	if m == nil {
		return nil
	}
	m.Bytes = bytes
	return m
}

// String
func (m *Report) String() string {
	return "to be implemented"
}
