//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"testing"
)

func TestBigchain(t *testing.T) {
	var headers map[string]string
	headers = make(map[string]string)
	headers["app_id"] = "b1d63ff3"
	headers["app_key"] = "29913c6deb7ee2bd0709d6af3b382b44"
	bcdb := NewBigChainDB("https://test.bigchaindb.com/api/v1/", &headers)

	//data := JsonObj{"bicycle": JsonObj{"serial_number": "abcd1234", "manufacturer": "bkfab"}}

	/*info, _ := bcdb.GetServerInfo()
	{
		b, err1 := json.Marshal(info)
		if err1 != nil {
			//t.Fatal(err)
		}
		println("Info: ", string(b))
	}*/

	trans := NewCreateTransaction(JsonObj{"AssetKey": "AssetValue"}, JsonObj{"MetaDataKey": "MetaDataValue"})
	pub, priv := GenerateKeypair()
	alicePublic := []PublicKey{pub}
	alicePrivate := []PrivateKey{priv}
	trans.AddOwnerBefore(&alicePublic, &alicePrivate)
	trans.AddOwnerAfter(&alicePublic, 1)
	trans.Sign()
	tx, _ := trans.Generate(true, false)
	//{
	//b, err1 := json.Marshal(tx)
	//if err1 != nil {
	//t.Fatal(err)
	//}
	//println("TX: ", string(b), "\r\n++++\r\n")
	//}
	tx1, err := bcdb.NewTransaction(tx)
	if err != nil {
		println("Error :", err.Error())
	} else {
		println("Tx Id: ", string(tx1))
	}

}
