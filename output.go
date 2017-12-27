//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	"errors"
	"strconv"
)

type output struct {
	outputItems []*outputItem
}

func (ou *output) Generate() []JsonObj {
	arr := []JsonObj{}
	for _, item := range ou.outputItems {
		it, _ := item.Generate()
		arr = append(arr, it)
	}
	return arr
}

func (ou *output) Sign(message string) error {
	for _, item := range ou.outputItems {
		item.Sign(message)
	}
	return nil
}

func (ou *output) Add(publicKey *[]PublicKey, amount int) {
	ot := outputItem{}
	ot.ownersAfters = publicKey
	ot.amount = amount
	ou.outputItems = append(ou.outputItems, &ot)
}

type outputItem struct {
	amount       int
	ownersAfters *[]PublicKey
}

func (o *outputItem) Generate() (JsonObj, error) {
	if o.ownersAfters == nil {
		return nil, nil
	}
	n := len(*o.ownersAfters)
	if n == 0 {
		return nil, errors.New("no ownersAfter")
	}
	if n == 1 {
		return JsonObj{
			"amount":      strconv.Itoa(o.amount),
			"condition":   o.creatCondition(),
			"public_keys": o.ownersAfters,
		}, nil
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

func (o *outputItem) Sign(message string) error {

	return nil
}

func (o *outputItem) creatCondition() JsonObj {

	return JsonObj{
		"detial": JsonObj{
			"public_key": "",
			"type":       "",
		},
		"uri": "",
	}
}
