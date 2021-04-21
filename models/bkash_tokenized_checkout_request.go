package models

type CreateAgreementRequest struct {
	Mode                  string `json:"mode"`
	PayerReference        string `json:"payerReference"`
	CallbackUrl           string `json:"callbackURL"`
	Amount                string `json:"amount"`
	Currency              string `json:"currency"`
	Intent                string `json:"intent"`
	MerchantInvoiceNumber string `json:"merchantInvoiceNumber"`
}

type ExecuteAgreementRequest struct {
	PaymentID string `json:"paymentID"`
}

type QueryAgreementRequest struct {
	AgreementID string `json:"agreementID"`
}

type CancelAgreementRequest struct {
	AgreementID string `json:"agreementID"`
}

type CreatePaymentRequest struct {
	Mode                    string `json:"mode"`
	PayerReference          string `json:"payerReference"`
	CallbackURL             string `json:"callbackURL"`
	AgreementID             string `json:"agreementID"`
	Amount                  string `json:"amount"`
	Currency                string `json:"currency"`
	Intent                  string `json:"intent"`
	MerchantInvoiceNumber   string `json:"merchantInvoiceNumber"`
	MerchantAssociationInfo string `json:"merchantAssociationInfo"`
}

type ExecutePaymentRequest struct {
	PaymentID string `json:"paymentID"`
}

type QueryPaymentRequest struct {
	PaymentID string `json:"paymentID"`
}

type SearchTransactionRequest struct {
	TrxID string `json:"trxID"`
}
