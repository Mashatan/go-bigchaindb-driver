//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"encoding/json"
	"testing"

	"golang.org/x/crypto/ed25519"
)

func GenerateKeypair() (PublicKey, PrivateKey) {
	pubInner, privInner, _ := ed25519.GenerateKey(nil)
	priv := PrivateKey(privInner)
	pub := PublicKey(pubInner)
	return pub, priv
}
func TestTransaction(t *testing.T) {

	trans := NewCreateTransaction(JsonObj{"Test1": "Test2"}, JsonObj{"Data1": "Data2"})
	pub, priv := GenerateKeypair()
	alicePublic := []PublicKey{pub}
	alicePrivate := []PrivateKey{priv}
	trans.AddOwnerBefore(&alicePublic, &alicePrivate)
	trans.AddOwnerAfter(&alicePublic, 1)

	trans.Sign()

	obj, _ := trans.Generate()
	{
		b, err1 := json.Marshal(obj)
		if err1 != nil {
			//t.Fatal(err)
		}
		println("TX: ", string(b), "\r\n*****\r\n")
	}
}
