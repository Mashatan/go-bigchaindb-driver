//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import (
	gcc "github.com/Mashatan/go-cryptoconditions"
	"golang.org/x/crypto/ed25519"
)

type input struct {
	inputItems []*inputItem
}

func (in *input) Generate(removeNull bool) []JsonObj {
	arr := []JsonObj{}
	for _, item := range in.inputItems {
		if !removeNull || item != nil {
			arr = append(arr, item.Generate(removeNull))
		}
	}
	return arr
}

func (in *input) Sign(message []byte) error {
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
	it.fulfills = nil
	in.inputItems = append(in.inputItems, &it)
}

type inputItem struct {
	ccEd            gcc.Ed25519Sha256
	ownerBefores    *[]PublicKey
	ownerPrivates   *[]PrivateKey
	ownerSignatures *[]Signature
	fulfills        *JsonObj `json:"fulfills,omitempty"`
}

func NewInputItem() inputItem {
	in := inputItem{}
	in.ownerBefores = nil
	in.ownerPrivates = nil
	in.ownerSignatures = nil
	in.fulfills = nil
	return in
}

func (i *inputItem) Generate(removeNull bool) JsonObj {
	var publicList []string
	for _, pk := range *i.ownerBefores {
		publicList = append(publicList, pk.String())
	}
	ii := JsonObj{}
	fulfilment := i.creatFulfillment()
	//if !removeNull || fulfilment != nil {
	ii["fulfillment"] = fulfilment
	//}

	//if !removeNull || i.fulfills != nil {
	ii["fulfills"] = i.fulfills
	//}
	if !removeNull || len(publicList) != 0 {
		ii["owners_before"] = publicList
	}
	return ii
}

func (i *inputItem) Sign(message []byte) (Signature, error) {
	if i.ownerPrivates == nil {
		return nil, nil
	}
	priv := (*i.ownerPrivates)[0]
	println("*\n message: ", string(message), "\n*\n")

	sgn := ed25519.Sign(ed25519.PrivateKey(priv), message)
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
	ss := string(Base64UrlEncode(tt))
	return &ss
}
