package main

import (
	"crypto/rsa"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseKey(t *testing.T) {
	token := jwt{
		Kty: "RSA",
		E:   "EAE",
		N:   "AQAB",
	}

	key := parseKey(&token)

	assert.Equal(t, 4097, key.E, "E does not match")
	assert.Equal(t, big.NewInt(65537), key.N, "N does not match")
}

func Test_encodeKey(t *testing.T) {
	key := rsa.PublicKey{
		E: 4097,
		N: big.NewInt(65537),
	}

	encoded := string(encodeKey(&key))

	assert.Equal(t, `-----BEGIN PUBLIC KEY-----
MB0wDQYJKoZIhvcNAQEBBQADDAAwCQIDAQABAgIQAQ==
-----END PUBLIC KEY-----
`, encoded)
}
