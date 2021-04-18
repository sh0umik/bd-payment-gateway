package models

type RefundTransactionResponse struct {
	CompletedTime     string `json:"completedTime"`
	TransactionStatus string `json:"transactionStatus"`
	OriginalTrxID     string `json:"originalTrxID"`
	RefundTrxID       string `json:"refundTrxID"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Charge            string `json:"charge"`
}

type RefundStatusResponse struct {
	RefundTransactionResponse
}
