package types

type IndexerExecutionOutcomeWithReceipt struct {
	Transaction SignedTransactionView                      `json:"transaction"`
	Outcome     IndexerExecutionOutcomeWithOptionalReceipt `json:"outcome"`
}
