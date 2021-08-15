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
	// General purpose domains
	DomainThis      = NewDomain("this")
	DomainSrc       = NewDomain("src")
	DomainDst       = NewDomain("dst")
	DomainReference = NewDomain("reference")
	DomainContext   = NewDomain("context")
	DomainTask      = NewDomain("task")
	DomainFile      = NewDomain("file")
	DomainStatus    = NewDomain("status")
	DomainParent    = NewDomain("parent")
	DomainReport    = NewDomain("report")
	DomainResult    = NewDomain("result")
	DomainInterim   = NewDomain("interim")
	// Predefined address domains
	DomainS3       = NewDomain("s3")
	DomainKafka    = NewDomain("kafka")
	DomainDigest   = NewDomain("digest")
	DomainUUID     = NewDomain("uuid")
	DomainUserID   = NewDomain("userid")
	DomainDirname  = NewDomain("dirname")
	DomainFilename = NewDomain("filename")
	DomainURL      = NewDomain("url")
	DomainDomain   = NewDomain("domain")
	DomainCustom   = NewDomain("custom")

	// List of all registered domains
	Domains = []*Domain{
		// General purpose domains
		DomainThis,
		DomainSrc,
		DomainDst,
		DomainReference,
		DomainContext,
		DomainTask,
		DomainFile,
		DomainStatus,
		DomainParent,
		DomainReport,
		DomainResult,
		DomainInterim,
		// Predefined address domains
		DomainS3,
		DomainKafka,
		DomainDigest,
		DomainUUID,
		DomainUserID,
		DomainDirname,
		DomainFilename,
		DomainURL,
		DomainDomain,
		DomainCustom,
	}
)

// RegisterDomain tries to register specified Domain.
// Domain must be non-equal to all registered domains.
// Returns nil in case Domain can not be registered, say it is equal to previously registered domain
func RegisterDomain(domain *Domain) *Domain {
	if FindDomain(domain) != nil {
		// Such a domain already exists
		return nil
	}
	Domains = append(Domains, domain)
	return domain
}

// MustRegisterDomain the same as RegisterDomain but with panic
func MustRegisterDomain(domain *Domain) {
	if RegisterDomain(domain) == nil {
		panic("unable to register domain")
	}
}

// FindDomain returns registered domain with the same string value as provided
func FindDomain(domain *Domain) *Domain {
	return DomainFromString(domain.Name)
}

// NormalizeDomain returns either registered domain with the same string value as provided domain or provided domain.
func NormalizeDomain(domain *Domain) *Domain {
	if f := FindDomain(domain); f != nil {
		// Return registered domain
		return f
	}
	// Unable to find registered domain, return provided domain
	return domain
}

// DomainFromString tries to find registered domain with specified string value
func DomainFromString(str string) *Domain {
	d := NewDomain().SetName(str)
	for _, domain := range Domains {
		if domain.Equals(d) {
			return domain
		}
	}
	return nil
}
