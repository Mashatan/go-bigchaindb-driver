//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"testing"
	"time"
)

func TestBigchain(t *testing.T) {
	var headers map[string]string
	headers = make(map[string]string)
	headers["app_id"] = "b1d63ff3"
	headers["app_key"] = "29913c6deb7ee2bd0709d6af3b382b44"
	bcdb := NewBigChainDB("https://test.bigchaindb.com/api/v1/", &headers)

	trans := NewCreateTransaction(JsonObj{"AssetKey": "AssetValue"}, JsonObj{"MetaDataKey": "MetaDataValue"})
	pub, priv := GenerateKeypair()
	alicePublic := []PublicKey{pub}
	alicePrivate := []PrivateKey{priv}
	trans.AddOwnerBefore(&alicePublic, &alicePrivate, nil)
	trans.AddOwnerAfter(&alicePublic, 1)
	trans.Sign()
	tx, _ := trans.Generate(true, false)

	txId, err := bcdb.NewTransaction(tx)

	if err != nil {
		println("Error :", err.Error())
	} else {
		println("Tx Id: ", string(txId))
	}

	var count int
	for ok := true; ok; ok = (count < 60) {
		status := bcdb.TransactionStatus(txId)
		if status {
			break
		} else {
			time.Sleep(1000 * time.Millisecond)
			count++
		}
	}
	bcdb.GetTransaction(txId)

	pub, priv = GenerateKeypair()
	bobPublic := []PublicKey{pub}
	//bobPrivate := []PrivateKey{priv}
	trans_trans := NewTransferTransaction(txId, JsonObj{"TransMetaDataKey": "TransMetaDataValue"})

	trans_trans.AddOwnerBefore(&alicePublic, &alicePrivate, &JsonObj{"transaction_id": txId, "output_index": 0})
	trans_trans.AddOwnerAfter(&bobPublic, 1)
	trans_trans.Sign()
	trans_tx, _ := trans_trans.Generate(true, false)
	txId1, err := bcdb.NewTransaction(trans_tx)
	if err != nil {
		println("Error :", err.Error())
	} else {
		println("Tx Id: ", string(txId1))
	}
}
