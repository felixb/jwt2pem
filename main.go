package main

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

type jwt struct {
	Kty string `json:"kty"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type pkcs8 struct {
	Algo      pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

var (
	oidPublicKeyRSA = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
)

func main() {
	if _, err := os.Stdout.Write(encodeKey(parseKey(readToken(os.Stdin)))); err != nil {
		log.Fatalf("error writing encoded key to stdout: %v", err)
	}
}

func readToken(r io.Reader) *jwt {
	var token jwt
	if b, err := ioutil.ReadAll(r); err != nil {
		log.Fatalf("error reading jwt token from stdin: %v", err)
	} else if err := json.Unmarshal(b, &token); err != nil {
		log.Fatalf("error parsing token: %v", err)
	} else if token.Kty != "RSA" {
		log.Fatalf("unsupported key type: %s", token.Kty)
	}

	return &token
}

func parseBigInt(s string) *big.Int {
	i := new(big.Int)

	if b, err := base64.RawURLEncoding.DecodeString(s); err != nil {
		log.Fatalf("error decoding numbers: %v", err)
	} else if _, err := fmt.Sscan("0x"+hex.EncodeToString(b), i); err != nil {
		log.Fatalf("error parsing big int: %v", err)
	}

	return i
}

func parseKey(token *jwt) *rsa.PublicKey {
	key := rsa.PublicKey{}
	key.E = int(parseBigInt(token.E).Int64())
	key.N = parseBigInt(token.N)
	return &key
}

func encodeKey(key *rsa.PublicKey) []byte {
	container := pkcs8{
		Algo: pkix.AlgorithmIdentifier{
			Algorithm:  oidPublicKeyRSA,
			Parameters: asn1.NullRawValue,
		},
		PublicKey: asn1.BitString{Bytes: x509.MarshalPKCS1PublicKey(key)},
	}

	if encodedKey, err := asn1.Marshal(container); err != nil {
		log.Fatalf("error encoding public key with asn1: %v", err)
		return nil
	} else {
		block := pem.Block{Type: "PUBLIC KEY", Bytes: encodedKey}
		return pem.EncodeToMemory(&block)
	}
}
