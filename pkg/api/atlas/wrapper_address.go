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

var AddressTypeEnum = NewEnum()

func init() {
	AddressTypeEnum.MustRegister("ADDRESS_RESERVED", int32(AddressType_ADDRESS_RESERVED))
	AddressTypeEnum.MustRegister("ADDRESS_S3", int32(AddressType_ADDRESS_S3))
	AddressTypeEnum.MustRegister("ADDRESS_KAFKA", int32(AddressType_ADDRESS_KAFKA))
}

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
