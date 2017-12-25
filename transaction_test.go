//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"encoding/json"
	"testing"
)

func TestTransaction(t *testing.T) {
	trans := NewCreateTransaction(JsonObj{"Test1": "Test2"}, JsonObj{"Data1": "Data2"})
	a1 := []PublicKey{{'g', 'o', 'l', 'a', 'n', 'g'}}
	trans.AddOwnerAfter(a1, 1)
	trans.AddOwnerBefore(a1, JsonObj{})

	obj, _ := trans.CreateTransaction()
	{
		b, err1 := json.Marshal(obj)
		if err1 != nil {
			//t.Fatal(err)
		}
		println("TX: ", string(b), "\r\n*****\r\n")
	}
}
