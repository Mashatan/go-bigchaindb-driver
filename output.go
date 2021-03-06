//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"bytes"
	"errors"
	"strconv"
	"strings"

	gcc "github.com/Mashatan/go-cryptoconditions"
)

type output struct {
	outputItems []*outputItem
}

func (ou *output) Generate(removeNull bool) []JsonObj {
	arr := []JsonObj{}
	for _, item := range ou.outputItems {
		it, _ := item.Generate(removeNull)
		if !removeNull || len(it) != 0 {
			arr = append(arr, it)
		}
	}
	return arr
}

func (ou *output) Sign(message []byte) error {
	for _, item := range ou.outputItems {
		item.Sign(message)
	}
	return nil
}

func (ou *output) Add(publicKey *[]PublicKey, amount int) {
	ot := outputItem{}
	ot.ownerAfters = publicKey
	ot.amount = amount
	ou.outputItems = append(ou.outputItems, &ot)
}

type outputItem struct {
	amount      int
	ownerAfters *[]PublicKey
	condition   *gcc.Conditions
}

func (o *outputItem) Generate(removeNull bool) (JsonObj, error) {
	if o.ownerAfters == nil {
		return nil, nil
	}
	n := len(*o.ownerAfters)
	if n == 0 {
		return nil, errors.New("no ownersAfter")
	}
	var publicList []string

	for _, pk := range *o.ownerAfters {
		publicList = append(publicList, pk.String())
	}

	if n == 1 {
		oi := JsonObj{}
		oi["amount"] = strconv.Itoa(o.amount)
		conditions := o.creatCondition()
		if !removeNull || conditions != nil {
			oi["condition"] = conditions
		}
		if !removeNull || len(publicList) != 0 {
			oi["public_keys"] = publicList
		}
		return oi, nil
	}
	return nil, nil
	/// NO SUPPORT YET
	/*
		fulfillment, err := cc.DefaultFulfillmentThresholdFromPubkeys(ownersAfter)
		if err != nil {
			return nil, err
		}
		return JsonObj{
			"amount":      strconv.Itoa(amount),
			"condition":   fulfillment.Data(),
			"public_keys": ownersAfter,
		}, nil
	*/
}

func (o *outputItem) Sign(message []byte) error {

	ee, _ := gcc.NewEd25519Sha256((*o.ownerAfters)[0], nil)
	o.condition = ee.Condition()
	return nil
}

func (o *outputItem) creatCondition() JsonObj {

	var typestr string
	var uri string
	var pk PublicKey
	if o.ownerAfters != nil {
		if len(*o.ownerAfters) > 0 {
			pk = (*o.ownerAfters)[0]
		}
	}
	if o.condition != nil {
		typestr = strings.ToLower(o.condition.Type().String())
		uri = o.condition.URI()
		b := []byte(uri)
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
		uri = string(b)
	}
	return JsonObj{
		"details": JsonObj{
			"public_key": pk.String(),
			"type":       typestr,
		},
		"uri": uri,
	}
}
