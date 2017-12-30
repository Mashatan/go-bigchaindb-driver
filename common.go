package GoBigChainDBDriver

import (
	"encoding/base64"
	"strings"

	base58 "github.com/jbenet/go-base58"
	"golang.org/x/crypto/sha3"
)

type JsonObj map[string]interface{}

type PublicKey []byte
type PrivateKey []byte
type Signature []byte

func Base64UrlEncode(p []byte) string {
	str := base64.RawURLEncoding.EncodeToString(p)
	{
		str = strings.Replace(str, "+", "-", -1)
		str = strings.Replace(str, "/", "_", -1)
		str = strings.Replace(str, "=", "", -1)
	}
	return str
}

func HashData(p []byte) []byte {
	digest := sha3.Sum256(p)
	return digest[:]
}

func (p PrivateKey) String() string {
	return base58.Encode([]byte(p))
}

func (p PublicKey) String() string {
	return base58.Encode([]byte(p))
}
