package GoBigChainDBDriver

type BigChainCore struct {
}

//////////////////////////////////////////////
/*
func FulfillTx(tx Data, fulfillments cc.Fulfillments) error {
	n := len(fulfillments)
	if n == 0 {
		return Error("no fulfillments")
	}
	inputs := GetTxInputs(tx)
	if n != len(inputs) {
		return Error("different number of fulfillments and inputs")
	}
	//println("Befor id ******\r\n", string(MustMarshalJSON(tx)), "\r\n***************\r\n")
	bb := MustMarshalJSON(tx)
	println(" \r\n Json final ======= \r\n", string(bb), " \r\n======= \r\n")
	tx.Set("id", BytesToHex(Checksum256(bb)))
	for i, fulfillment := range fulfillments {
		inputs[i].Set("fulfillment", fulfillment.String())
	}
	return nil
}
*/
