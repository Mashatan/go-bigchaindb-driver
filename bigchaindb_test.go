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
	headers["app_id"] = ""
	headers["app_key"] = ""
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

}
