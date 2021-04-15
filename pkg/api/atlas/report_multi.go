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

// NewReportMulti
func NewReportMulti() *ReportMulti {
	return new(ReportMulti)
}

// EnsureHeader
func (m *ReportMulti) EnsureHeader() *Metadata {
	if m == nil {
		return nil
	}
	if m.HeaderOptional == nil {
		m.HeaderOptional = new(ReportMulti_Header)
	}
	if m.HeaderOptional.(*ReportMulti_Header).Header == nil {
		m.HeaderOptional.(*ReportMulti_Header).Header = NewMetadata()
	}
	return m.GetHeader()
}

// AddReport
func (m *ReportMulti) AddReport(report *Report) *ReportMulti {
	if m == nil {
		return nil
	}
	m.Reports = append(m.Reports, report)
	return m
}

// AddReports
func (m *ReportMulti) AddReports(reports ...*Report) *ReportMulti {
	if m == nil {
		return nil
	}
	m.Reports = append(m.Reports, reports...)
	return m
}

// String
func (m *ReportMulti) String() string {
	return "to be implemented"
}
