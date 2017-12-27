//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	gcc "github.com/Mashatan/go-cryptoconditions"
	base58 "github.com/jbenet/go-base58"
	"golang.org/x/crypto/ed25519"
)

type input struct {
	inputItems []*inputItem
}

func (in *input) Generate() []JsonObj {
	arr := []JsonObj{}
	for _, item := range in.inputItems {
		arr = append(arr, item.Generate())
	}
	return arr
}

func (in *input) Sign(message string) error {
	for _, item := range in.inputItems {
		item.Sign(message)
	}
	return nil
}

func (in *input) Add(publicKey *[]PublicKey, privateKey *[]PrivateKey) {
	it := NewInputItem()
	it.ownerBefores = publicKey
	it.ownerPrivates = privateKey
	it.ownerSignatures = nil
	it.fulfills = &JsonObj{}
	in.inputItems = append(in.inputItems, &it)
}

type inputItem struct {
	ccEd            gcc.Ed25519Sha256
	ownerBefores    *[]PublicKey
	ownerPrivates   *[]PrivateKey
	ownerSignatures *[]Signature
	fulfills        *JsonObj
}

func NewInputItem() inputItem {
	in := inputItem{}
	in.ownerBefores = nil
	in.ownerPrivates = nil
	in.ownerSignatures = nil
	in.fulfills = nil
	return in
}

func (i *inputItem) Generate() JsonObj {
	return JsonObj{
		"fulfillment":   i.creatFulfillment(),
		"fulfills":      i.fulfills,
		"owners_before": i.ownerBefores,
	}
}

func (i *inputItem) Sign(message string) (Signature, error) {
	if i.ownerPrivates == nil {
		return nil, nil
	}
	priv := (*i.ownerPrivates)[0]
	sgn := ed25519.Sign(ed25519.PrivateKey(priv), []byte(message))
	if i.ownerSignatures == nil {
		i.ownerSignatures = new([]Signature)
		*(i.ownerSignatures) = append(*(i.ownerSignatures), sgn)
	}
	return sgn, nil
}

func (i *inputItem) creatFulfillment() *string {
	if i.ownerSignatures == nil {
		return nil
	}
	ee, _ := gcc.NewEd25519Sha256((*i.ownerBefores)[0], (*i.ownerSignatures)[0])
	tt, _ := ee.Encode()
	ss := string(base58.Encode(tt))
	return &ss
}
