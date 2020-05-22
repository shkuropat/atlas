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
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func parseRSAPublicKey(pem []byte) (*rsa.PublicKey, error) {
	var err error

	// Ensure PEM is reasonable
	err = verifyPEM(pem)
	if err != nil {
		return nil, fmt.Errorf("unable to verify PEM %v", err)
	}

	// Parse RSA Public Key
	return jwt.ParseRSAPublicKeyFromPEM(pem)
}

func verifyPEM(pem []byte) error {
	// Trim newlines from the end of the key
	trimmed := strings.TrimRight(string(pem), "\r\n")

	// Key must be enclosed into start/stop tags
	publicKeyPrefix := "-----BEGIN PUBLIC KEY-----"
	publicKeySuffix := "-----END PUBLIC KEY-----"
	err := fmt.Errorf("public key must be enclosed in %s %s", publicKeyPrefix, publicKeySuffix)

	if !strings.HasPrefix(trimmed, publicKeyPrefix) {
		return err
	}

	if !strings.HasSuffix(trimmed, publicKeySuffix) {
		return err
	}

	return nil
}
