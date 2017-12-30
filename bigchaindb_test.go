//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"encoding/json"
	"testing"
)

var (
	Alice = "3th33iKfYoPXQ6YL8mXcD3gzgMppEEHFBPFqch4Cn5d3"
)

func TestBigchain(t *testing.T) {
	var headers map[string]string
	headers = make(map[string]string)
	headers["app_id"] = "25977264"
	headers["app_key"] = "ca2e815e03f0595983034975cfee8c4b"
	bcdb := NewBigChainDB("https://test.ipdb.io/api/v1/", &headers)

	//data := JsonObj{"bicycle": JsonObj{"serial_number": "abcd1234", "manufacturer": "bkfab"}}

	info, _ := bcdb.GetServerInfo()
	{
		b, err1 := json.Marshal(info)
		if err1 != nil {
			//t.Fatal(err)
		}
		println("Info: ", string(b))
	}

	trans := NewCreateTransaction(JsonObj{"Test1": "Test2"}, JsonObj{"Data1": "Data2"})
	pub, priv := GenerateKeypair()
	alicePublic := []PublicKey{pub}
	alicePrivate := []PrivateKey{priv}
	trans.AddOwnerBefore(&alicePublic, &alicePrivate)
	trans.AddOwnerAfter(&alicePublic, 1)
	trans.Sign()
	tx, _ := trans.Generate()
	{
		b, err1 := json.Marshal(tx)
		if err1 != nil {
			//t.Fatal(err)
		}
		println("TX: ", string(b), "\r\n*****\r\n")
	}
	tx1, err := bcdb.NewTransaction(tx)
	{
		println("Error :", err.Error())
		println("Tx Id: ", string(tx1))
	}

}
