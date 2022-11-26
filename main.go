package main

import (
	"log"
	"time"
	"tracker/keepers"
	"tracker/types"
)

func main() {

	tk := keepers.NewTransactionKeeper()

	txList := tk.GetTxs()

	//printTx(txList)

	tk.AddTx(&types.Transaction{
		Id:       "lkajsdlkjgiasdlkajsdu12",
		DateTime: time.Now(),
		Coin:     "Bitcoin",
		Value:    163234523.52,
	})

	printTx(txList)
}

func printTx(txList map[string]*types.Transaction) {
	for _, tx := range txList {
		log.Printf("tx: %v\n", tx)
	}
}
