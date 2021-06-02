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

var (
	// Most popular predefined root domains
	DomainThis      = NewDomain().Set("this")
	DomainSrc       = NewDomain().Set("src")
	DomainDst       = NewDomain().Set("dst")
	DomainReference = NewDomain().Set("reference")
	DomainContext   = NewDomain().Set("context")
	DomainTask      = NewDomain().Set("task")
	DomainParent    = NewDomain().Set("parent")
	// Most popular predefined nested domains
	DomainS3           = NewDomain().Set("s3")
	DomainKafka        = NewDomain().Set("kafka")
	DomainDigest       = NewDomain().Set("digest")
	DomainUUID         = NewDomain().Set("uuid")
	DomainUserID       = NewDomain().Set("userid")
	DomainDirname      = NewDomain().Set("dirname")
	DomainFilename     = NewDomain().Set("filename")
	DomainURL          = NewDomain().Set("url")
	DomainDomain       = NewDomain().Set("domain")
	DomainCustomString = NewDomain().Set("custom string")
)
