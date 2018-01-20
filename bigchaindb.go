//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"bytes"
	"crypto"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BigChainDB struct {
	RootEndpoint string
	Headers      *map[string]string
}

func NewBigChainDB(rootEndpoint string, headers *map[string]string) BigChainDB {
	bigChainDB := BigChainDB{}
	bigChainDB.RootEndpoint = rootEndpoint
	bigChainDB.Headers = headers
	return bigChainDB
}

func (bc *BigChainDB) request(action string, method string, sendData interface{}, reciveData interface{}) error {

	client := &http.Client{}
	url := bc.RootEndpoint + action
	var buf *bytes.Buffer = nil
	println("url: ", url)
	var req *http.Request
	var err error
	if sendData != nil {
		b, err := json.Marshal(sendData)
		if err != nil {
			return errors.New("error")
		}
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

		buf = bytes.NewBuffer(b)
		req, err = http.NewRequest(strings.ToUpper(method), url, buf)
	} else {
		req, err = http.NewRequest(strings.ToUpper(method), url, nil)
	}
	if err != nil {
		return errors.New("error")
	}
	if sendData != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if bc.Headers != nil {
		for k, v := range *bc.Headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode < 200 && resp.StatusCode > 202 {
		//fmt.Println(string(body))
		return errors.New(string(body))
	}
	json.Unmarshal([]byte(body), &reciveData)
	return nil
}

func (bc *BigChainDB) GetServerInfo() (JsonObj, error) {
	req := ""
	tx := make(JsonObj)
	if err := bc.request(req, "GET", nil, &tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// GET requests
func (bc *BigChainDB) GetTransaction(transaction_id string) (JsonObj, error) {
	req := "transactions/" + transaction_id
	tx := make(JsonObj)
	if err := bc.request(req, "GET", nil, &tx); err != nil {
		return nil, err
	}
	/***********************
	fulfilled, err := FulfilledTx(tx)
	if err != nil {
		return nil, err
	}
	if !fulfilled {
		return nil, error.Error("unfulfilled")
	}
	*****************************/
	return tx, nil
}

//operation: TRANSFER or CREATE
func (bc *BigChainDB) GetListTransactions(assetId string, operation string) ([]JsonObj, error) {
	req := fmt.Sprintf("transactions?operation=%v&asset_id=%v", operation, assetId)
	var txs []JsonObj
	if err := bc.request(req, "GET", nil, &txs); err != nil {
		return nil, err
	}
	/*******************************
	for _, tx := range txs {
		fulfilled, err := FulfilledTx(tx)
		if err != nil {
			return nil, err
		}
		if !fulfilled {
			return nil, common.Error("unfulfilled")
		}
	}
	************************************/
	return txs, nil
}

func (bc *BigChainDB) HttpGetOutputs(pubkey crypto.PublicKey, unspent bool) ([]string, []int, error) {
	req := fmt.Sprintf("outputs?public_key=%v&unspent=%v", pubkey, unspent)
	var links []string
	if err := bc.request(req, "GET", nil, &links); err != nil {
		return nil, nil, err
	}
	txIds := make([]string, len(links))
	outputs := make([]int, len(links))
	/************************
	for i, link := range links {
		submatch := common.SubmatchStr(`transactions/(.*?)/outputs/([0-9]{1,2})`, link)
		txIds[i], outputs[i] = submatch[1], common.MustAtoi(submatch[2])
	}
	*************************/
	return txIds, outputs, nil
}

func (bc *BigChainDB) HttpGetFilter(fn func(string) (JsonObj, error), pubkey crypto.PublicKey, unspent bool) ([]JsonObj, error) {
	txIds, _, err := bc.HttpGetOutputs(pubkey, unspent)
	if err != nil {
		return nil, err
	}
	var jsonObjs []JsonObj
	for _, txId := range txIds {
		//tx, err := fn(txId)
		_, err := fn(txId)
		if err == nil {
			//jsonObjs = append(jsonObjs, GetTxAssetData(tx))
		}
	}
	return jsonObjs, nil
}

// POST request

func (bc *BigChainDB) NewTransaction(transaction JsonObj) (string, error) {
	req := "transactions/"
	var response JsonObj
	response = make(JsonObj)
	if err := bc.request(req, "POST", transaction, &response); err != nil {
		return "", err
	}
	str := response["id"].(string)

	return str, nil
	//return "", nil
}
