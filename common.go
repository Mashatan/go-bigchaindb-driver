package GoBigChainDBDriver

import (
	"golang.org/x/crypto/sha3"
)

type JsonObj map[string]interface{}

type PublicKey []byte
type PrivateKey []byte
type Signature []byte

func HashData(p []byte) []byte {
	digest := sha3.Sum256(p)
	return digest[:]
}
