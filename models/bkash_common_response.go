package models

type RefundTransactionResponse struct {
	CompletedTime     string `json:"completedTime,omitempty"`
	TransactionStatus string `json:"transactionStatus,omitempty"`
	OriginalTrxID     string `json:"originalTrxID,omitempty"`
	RefundTrxID       string `json:"refundTrxID,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Currency          string `json:"currency,omitempty"`
	Charge            string `json:"charge,omitempty"`
}

type RefundStatusResponse struct {
	RefundTransactionResponse
}
