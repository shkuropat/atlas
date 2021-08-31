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

package jwks

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWKS specifies JSON Web Key Set (JWKS)
// JWKS is a set of public keys used to verify JWT issued by authorization server
// https://auth0.com/docs/security/tokens/json-web-tokens/json-web-key-sets
type JWKS struct {
	Keys []*JWK `json:"keys,omitempty" yaml:"keys,omitempty"`
	File string `json:"file,omitempty" yaml:"file,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
	Data string `json:"data,omitempty" yaml:"data,omitempty"`
}

// New creates new JWKS
func New() *JWKS {
	return new(JWKS)
}

// NewFromReader creates new JWKS from reader which providers JSON
func NewFromReader(jsonJWKS io.Reader) (*JWKS, error) {
	jwks := New()
	err := json.NewDecoder(jsonJWKS).Decode(jwks)
	if err != nil {
		return nil, err
	}
	return jwks, nil
}

// NewFromString creates new JWKS from JSON string
func NewFromString(jsonJWKS string) (*JWKS, error) {
	return NewFromReader(strings.NewReader(jsonJWKS))
}

// NewFromBytes creates new JWKS from JSON bytes
func NewFromBytes(json []byte) (*JWKS, error) {
	return NewFromString(string(json))
}

// NewFromURL creates new JWKS from URL which providers JSON
func NewFromURL(url string) (*JWKS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return NewFromReader(resp.Body)
}

// NewFromFile
func NewFromFile(filename string) (*JWKS, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read JWKS from file '%s': %v", filename, err)
	}

	return NewFromBytes(bytes)
}

// Append appends JWK
func (jwks *JWKS) Append(jwk ...*JWK) *JWKS {
	if len(jwk) == 0 {
		// Nothing to append
		return jwks
	}
	if jwks == nil {
		jwks = New()
	}
	jwks.Keys = append(jwks.Keys, jwk...)
	return jwks
}

// ReadIn
func (jwks *JWKS) ReadIn() *JWKS {
	if jwks == nil {
		return nil
	}
	if jwks.File != "" {
		if file, err := NewFromFile(jwks.File); (file != nil) && (err == nil) {
			jwks.Append(file.Keys...)
		}
	}
	if jwks.URL != "" {
		if url, err := NewFromURL(jwks.URL); (url != nil) && (err == nil) {
			jwks.Append(url.Keys...)
		}
	}
	if jwks.Data != "" {
		if data, err := NewFromString(jwks.Data); (data != nil) && (err == nil) {
			jwks.Append(data.Keys...)
		}
	}

	for i := range jwks.Keys {
		if new := jwks.Keys[i].ReadIn(); new != nil {
			jwks.Keys[i] = new
		}
	}

	return jwks
}

// GetDefaultKey gets default key
func (jwks *JWKS) GetDefaultKey() *JWK {
	if jwks == nil {
		return nil
	}
	if len(jwks.Keys) < 1 {
		return nil
	}
	return jwks.Keys[0]
}

// GetVerificationPublicKey searches for cert specified in token's header
func (jwks *JWKS) GetVerificationPublicKey(token *jwt.Token, fallbackToDefault bool) (*rsa.PublicKey, error) {
	for i := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[i].Kid {
			return jwks.Keys[i].PublicKey, nil
		}
	}
	if fallbackToDefault {
		if def := jwks.GetDefaultKey(); def != nil {
			return def.PublicKey, nil
		}
	}
	return nil, errors.New("unable to find appropriate key")
}

func (jwks *JWKS) String() string {
	return "jwks string rep"
}
