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

import "strings"

// NewS3Address
func NewS3Address(bucket, object string) *S3Address {
	return &S3Address{
		Bucket: bucket,
		Object: object,
	}
}

// NewS3AddressFromString
func NewS3AddressFromString(str string) *S3Address {
	parts := strings.SplitN(str, "/", 2)
	if len(parts) != 2 {
		return nil
	}
	bucket := parts[0]
	object := parts[1]
	return NewS3Address(bucket, object)
}

// SetBucket
func (m *S3Address) SetBucket(bucket string) *S3Address {
	if m == nil {
		return nil
	}
	m.Bucket = bucket
	return m
}

// SetObject
func (m *S3Address) SetObject(object string) *S3Address {
	if m == nil {
		return nil
	}
	m.Object = object
	return m
}

// String
func (m *S3Address) String() string {
	if m != nil {
		return m.Bucket + "/" + m.Object
	}
	return ""
}
