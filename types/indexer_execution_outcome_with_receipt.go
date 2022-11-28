package types

type IndexerExecutionOutcomeWithReceipt struct {
	transaction SignedTransactionView                      `json:"transaction"`
	outcome     IndexerExecutionOutcomeWithOptionalReceipt `json:"outcome"`
}
