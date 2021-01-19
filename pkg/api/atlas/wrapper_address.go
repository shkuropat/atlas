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

func NewAddress(a interface{}) *Address {
	if a, ok := a.(isAddress_Address); ok {
		return &Address{
			Address: a,
		}
	}

	return &Address{}
}

func (m *Address) Printable() string {
	if m == nil {
		return ""
	}
	switch {
	case m.GetS3() != nil:
		return m.GetS3().Printable()
	case m.GetKafka() != nil:
		return m.GetKafka().Printable()
	}
	return "UNKNOWN ADDRESS"
}
