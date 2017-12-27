//Date: 2017 Q4
//Email: ali.mashatan@gmail.com
//Author: Ali Mashatan

package GoBigChainDBDriver

import "encoding/json"

const (
	CREATE   = "CREATE"
	GENESIS  = "GENSIS"
	TRANSFER = "TRANSFER"
	VERSION  = "1.0"
)

type transaction struct {
	id        string
	asset     JsonObj
	input     input
	output    output
	metadata  JsonObj
	operation string
	version   string
}

func NewCreateTransaction(asset JsonObj, metadata JsonObj) transaction {
	trasaction := transaction{}
	trasaction.operation = CREATE
	trasaction.version = VERSION
	trasaction.asset = JsonObj{"data": asset}
	trasaction.metadata = metadata
	trasaction.input = input{}
	trasaction.output = output{}
	return trasaction
}

func NewTransferTransaction(asset JsonObj, metadata JsonObj) transaction {
	trasaction := transaction{}
	trasaction.operation = TRANSFER
	trasaction.version = VERSION
	trasaction.asset = JsonObj{"data": asset}
	trasaction.metadata = metadata
	trasaction.input = input{}
	trasaction.output = output{}
	return trasaction
}

func (t *transaction) AddOwnerAfter(publicKey *[]PublicKey, amount int) error {
	t.output.Add(publicKey, amount)
	return nil
}

func (t *transaction) AddOwnerBefore(publicKey *[]PublicKey, privateKey *[]PrivateKey) error {
	t.input.Add(publicKey, privateKey)
	return nil
}

func (t *transaction) Generate() (JsonObj, error) {

	tx := JsonObj{
		"asset":     t.asset,
		"inputs":    t.input.Generate(),
		"metadata":  t.metadata,
		"operation": t.operation,
		"outputs":   t.output.Generate(),
		"version":   VERSION,
	}
	return tx, nil
}

func (t *transaction) dump() string {
	jo, _ := t.Generate()
	b, _ := json.Marshal(jo)
	return string(b)
}

func (t *transaction) Sign() error {

	dm := t.dump()
	t.input.Sign(dm)
	t.output.Sign(dm)

	return nil
}
