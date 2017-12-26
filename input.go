//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

type input struct {
	inputItems []inputItem
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

func (in *input) Add(publicKey []PublicKey, fulfill JsonObj) {
	it := inputItem{}
	it.ownerBefores = publicKey
	it.fulfills = fulfill
	in.inputItems = append(in.inputItems, it)
}

type inputItem struct {
	ownerBefores []PublicKey
	fulfills     JsonObj
}

func (i *inputItem) Generate() JsonObj {
	return JsonObj{
		"fulfillment":   i.creatFulfillment(),
		"fulfills":      i.fulfills,
		"owners_before": i.ownerBefores,
	}
}

func (i *inputItem) Sign(message string) error {

	return nil
}

func (i *inputItem) creatFulfillment() string {

	return ""
}
