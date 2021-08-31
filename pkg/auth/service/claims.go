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

package service_auth

import (
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// ScopeClaims
//	{
//		"aud":"https://atlas-aud",
//		"iss":"https://issuer.url/",
//		"sub":"sub@clients",
//		"iat":1630265536,
//		"exp":1630351936,
//		"azp":"some azp",
//		"gty":"client-credentials"
//		"scope":"permission-scope-read",
//	}
type ScopeClaims struct {
	jwt.StandardClaims
	Scope string `json:"scope"`
}

// Valid
func (c ScopeClaims) Valid() error {
	return c.StandardClaims.Valid()
}

// Dump
func (c ScopeClaims) Dump() {
	log.Infof("Customized Claims:")
	log.Infof("Audience  %s", c.Audience)
	log.Infof("ExpiresAt %d", c.ExpiresAt)
	log.Infof("Id        %s", c.Id)
	log.Infof("IssuedAt  %d", c.IssuedAt)
	log.Infof("Issuer    %s", c.Issuer)
	log.Infof("NotBefore %d", c.NotBefore)
	log.Infof("Subject   %s", c.Subject)
	log.Infof("Scope     %s", c.Scope)
}
