//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

func GenerateKeypair() (PublicKey, PrivateKey) {
	//pubInner, privInner, _ := ed25519.GenerateKey(nil)

	const publicStr = "a9b030a8738d36b9d17d2ade627a5971e0a1e7bdb95e51d7ae54387f48adad8a"
	const privateStr = "51d8d046a2fed79881ad7e4a4d5eded047cf33743a49f40ea1511d36fac98be8a9b030a8738d36b9d17d2ade627a5971e0a1e7bdb95e51d7ae54387f48adad8a"
	pubInner, _ := hex.DecodeString(publicStr)
	privInner, _ := hex.DecodeString(privateStr)
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
		println("TX: ", string(b), "\r\n+++++++\r\n")
	}
}
