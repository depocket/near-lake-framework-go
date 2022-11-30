package types

type TransactionAction string

const (
	FunctionCallAction = "FunctionCall"
	TransferAction     = "Transfer"
)

type SignedTransactionView struct {
	SignerId   string        `json:"signer_id"`
	PublicKey  string        `json:"public_key"`
	Nonce      uint64        `json:"nonce"`
	ReceiverId string        `json:"receiver_id"`
	Actions    []interface{} `json:"actions"`
	Signature  string        `json:"signature"`
	Hash       string        `json:"hash"`
}

func (s SignedTransactionView) LoopActions(f func(interface{})) {
	for _, action := range s.Actions {
		f(action)
	}
}
