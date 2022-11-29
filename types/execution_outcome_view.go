package types

type ExecutionOutcomeView struct {
	Logs        []string               `json:"logs"`
	ReceiptIds  []string               `json:"receipt_ids"`
	GasBurnt    uint64                 `json:"gas_burnt"`
	TokensBurnt *BigInt                `json:"tokens_burnt"`
	ExecutorId  string                 `json:"executor_id"`
	Status      map[string]interface{} `json:"status"`
	Metadata    ExecutionMetadataView  `json:"metadata"`
}
