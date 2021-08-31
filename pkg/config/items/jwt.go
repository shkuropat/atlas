// Copyright 2021 The Atlas Authors. All rights reserved.
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

package items

import (
	"bytes"
	"fmt"
	jwks2 "github.com/binarly-io/atlas/pkg/auth/jwks"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type JWT struct {
	Audience string      `mapstructure:"audience"`
	Issuer   string      `mapstructure:"issuer"`
	JWKS     *jwks2.JWKS `mapstructure:"jwks"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewJWT
func NewJWT() *JWT {
	return new(JWT)
}

// GetAudience
func (o *JWT) GetAudience() string {
	if o == nil {
		return ""
	}
	return o.Audience
}

// GetIssuer
func (o *JWT) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

// GetJWKS
func (o *JWT) GetJWKS() *jwks2.JWKS {
	if o == nil {
		return nil
	}
	return o.JWKS
}

// String
func (o *JWT) String() string {
	if o == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Audience: %v\n", o.Audience)
	_, _ = fmt.Fprintf(b, "Issuer: %v\n", o.Issuer)
	_, _ = fmt.Fprintf(b, "JWKS: %v\n", o.JWKS)

	return b.String()
}
