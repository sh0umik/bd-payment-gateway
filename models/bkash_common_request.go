package models

type RefundTransactionRequest struct {
	PaymentID string `json:"paymentID"`
	Amount    string `json:"amount"`
	TrxID     string `json:"trxID"`
	Sku       string `json:"sku"`
	Reason    string `json:"reason"`
}

type RefundStatusRequest struct {
	PaymentID string `json:"paymentID"`
	TrxID     string `json:"trxID"`
}
