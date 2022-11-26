package keepers

import (
	"time"
	"tracker/types"
)

type TransactionKeeper struct {
	store map[string]*types.Transaction //Map llamado store de tipo Transaction con indices strings
}

func NewTransactionKeeper() *TransactionKeeper {
	tk := &TransactionKeeper{}
	tk.store = make(map[string]*types.Transaction)

	//Nueva tx
	tk.store["idtk1k5ma4ksma6amgka34kasw"] = &types.Transaction{
		Id:       "idtk1k5ma4ksma6amgka34kasw",
		DateTime: time.Now(),
		Coin:     "osmosis",
		Value:    10000,
	}

	return tk
}

func (tk *TransactionKeeper) AddTx(transaction *types.Transaction) {
	tk.store[transaction.Id] = transaction
}

func (tk *TransactionKeeper) GetTxs() map[string]*types.Transaction {
	return tk.store
}

func (tk *TransactionKeeper) GetTxById(id string) *types.Transaction {
	return tk.store[id]
}

func (tk *TransactionKeeper) DeleteTx(id string) {

}

func (tk *TransactionKeeper) UpdateTx(id string) {

}
