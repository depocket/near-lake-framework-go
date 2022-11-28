package types

type ReceiptView struct {
	PredecessorId string                 `json:"predecessor_id"`
	ReceiverId    string                 `json:"receiver_id"`
	ReceiptId     string                 `json:"receipt_id"`
	Receipt       map[string]interface{} `json:"receipt"`
}
