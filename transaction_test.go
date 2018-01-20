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

	const publicStr = "cdedae9c1fed4227e6209f594f0791e9fa816ec0e27491e72a934906cdbd645e" //"a9b030a8738d36b9d17d2ade627a5971e0a1e7bdb95e51d7ae54387f48adad8a"
	const privateStr = "3c75939b824e878ece0a4c36163759b64d830fe92064124867c6032258745243cdedae9c1fed4227e6209f594f0791e9fa816ec0e27491e72a934906cdbd645e"

	pubInner, _ := hex.DecodeString(publicStr)
	pub := PublicKey(pubInner)

	privInner, _ := hex.DecodeString(privateStr)
	priv := PrivateKey(privInner)
	return pub, priv
}
func TestTransaction(t *testing.T) {

	trans := NewCreateTransaction(JsonObj{"AssetKey": "AssetValue"}, JsonObj{"MetaDataKey": "MetaDataValue"})
	pub, priv := GenerateKeypair()
	alicePublic := []PublicKey{pub}
	alicePrivate := []PrivateKey{priv}
	trans.AddOwnerBefore(&alicePublic, &alicePrivate)
	trans.AddOwnerAfter(&alicePublic, 1)

	trans.Sign()

	obj, _ := trans.Generate(true, false)
	{
		b, err1 := json.Marshal(obj)
		if err1 != nil {
			//t.Fatal(err)
		}
		println("TX: ", string(b), "\r\n+++++++\r\n")
	}
}
