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

package pem

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	prefix = "-----BEGIN PUBLIC KEY-----"
	suffix = "-----END PUBLIC KEY-----"
)

var (
	ErrEnclosed = fmt.Errorf("public key must be enclosed in %s %s", prefix, suffix)
)

// EnsurePEM ensures cert is wrapped into -----BEGIN CERTIFICATE----- and -----END CERTIFICATE-----
func EnsurePEM(cert string) string {
	cert = trim(cert)
	if strings.HasPrefix(cert, prefix) {
		return cert
	}
	return fmt.Sprintf("%s\n%s\n%s", prefix, cert, suffix)
}

// ParseRSAPublicKeyFromPEM
func ParseRSAPublicKeyFromPEM(pem []byte) (*rsa.PublicKey, error) {
	// Ensure PEM is reasonable
	err := VerifyPEMFromBytes(pem)
	if err != nil {
		return nil, fmt.Errorf("unable to verify PEM %v", err)
	}

	// Parse RSA Public Key
	return jwt.ParseRSAPublicKeyFromPEM(pem)
}

// VerifyPEMFromBytes
func VerifyPEMFromBytes(cert []byte) error {
	trimmed := trim(string(cert))
	// Key must be enclosed into start/stop tags
	if !strings.HasPrefix(trimmed, prefix) || !strings.HasSuffix(trimmed, suffix) {
		return ErrEnclosed
	}
	return nil
}

// VerifyPEMFromString
func VerifyPEMFromString(cert string) error {
	trimmed := trim(cert)
	// Key must be enclosed into start/stop tags
	if !strings.HasPrefix(trimmed, prefix) || !strings.HasSuffix(trimmed, suffix) {
		return ErrEnclosed
	}
	return nil
}

// Trim trims PEM prefix/suffix
func Trim(cert string) string {
	cert = trim(cert)
	cert = strings.TrimPrefix(cert, prefix)
	cert = trim(cert)
	cert = strings.TrimSuffix(cert, suffix)
	return trim(cert)
}

// trim newlines from the end of the key
func trim(cert string) string {
	return strings.TrimRight(cert, "\r\n")
}
