package models

// Response models for TOKENIZED CHECKOUT

type BkashError struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type Token struct {
	TokenType     string `json:"token_type"`
	ExpiresIn     int    `json:"expires_in"`
	IdToken       string `json:"id_token"`
	RefreshToken  string `json:"refresh_token"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

type CreateAgreementResponse struct {
	PaymentID            string `json:"paymentID"`
	BkashURL             string `json:"bkashURL"`
	CallbackURL          string `json:"callbackURL"`
	SuccessCallbackURL   string `json:"successCallbackURL"`
	FailureCallbackURL   string `json:"failureCallbackURL"`
	CancelledCallbackURL string `json:"cancelledCallbackURL"`
	StatusCode           string `json:"statusCode"`
	StatusMessage        string `json:"statusMessage"`
}

type ExecuteAgreementResponse struct {
	PaymentID            string `json:"paymentID"`
	AgreementID          string `json:"agreementID"`
	CustomerMsisdn       string `json:"customerMsisdn"`
	PayerReference       string `json:"payerReference"`
	AgreementExecuteTime string `json:"agreementExecuteTime"`
	AgreementStatus      string `json:"agreementStatus"`
	StatusCode           string `json:"statusCode"`
	StatusMessage        string `json:"statusMessage"`
}

type QueryAgreementResponse struct {
	PaymentID            string `json:"paymentID"`
	AgreementID          string `json:"agreementID"`
	PayerReference       string `json:"payerReference"`
	CustomerMsisdn       string `json:"customerMsisdn"`
	AgreementCreateTime  string `json:"agreementCreateTime"`
	AgreementExecuteTime string `json:"agreementExecuteTime"`
	AgreementVoidTime    string `json:"agreementVoidTime"`
	AgreementStatus      string `json:"agreementStatus"`
	StatusCode           string `json:"statusCode"`
	StatusMessage        string `json:"statusMessage"`
}

type CancelAgreementResponse struct {
	PaymentID         string `json:"paymentID"`
	AgreementID       string `json:"agreementID"`
	PayerReference    string `json:"payerReference"`
	AgreementVoidTime string `json:"agreementVoidTime"`
	AgreementStatus   string `json:"agreementStatus"`
	StatusCode        string `json:"statusCode"`
	StatusMessage     string `json:"statusMessage"`
}

type CreatePaymentResponse struct {
	PaymentID             string `json:"paymentID"`
	AgreementID           string `json:"agreementID"`
	PaymentCreateTime     string `json:"paymentCreateTime"`
	TransactionStatus     string `json:"transactionStatus"`
	Amount                string `json:"amount"`
	Currency              string `json:"currency"`
	Intent                string `json:"intent"`
	MerchantInvoiceNumber string `json:"merchantInvoiceNumber"`
	BkashURL              string `json:"bkashURL"`
	CallbackURL           string `json:"callbackURL"`
	SuccessCallbackURL    string `json:"successCallbackURL"`
	FailureCallbackURL    string `json:"failureCallbackURL"`
	CancelledCallbackURL  string `json:"cancelledCallbackURL"`
	StatusCode            string `json:"statusCode"`
	StatusMessage         string `json:"statusMessage"`
}

type ExecutePaymentResponse struct {
	PaymentID             string `json:"paymentID"`
	AgreementID           string `json:"agreementID"`
	CustomerMsisdn        string `json:"customerMsisdn"`
	PayerReference        string `json:"payerReference"`
	AgreementExecuteTime  string `json:"agreementExecuteTime"`
	AgreementStatus       string `json:"agreementStatus"`
	PaymentExecuteTime    string `json:"paymentExecuteTime"`
	TrxID                 string `json:"trxID"`
	TransactionStatus     string `json:"transaction_status"`
	Amount                string `json:"amount"`
	Currency              string `json:"currency"`
	Intent                string `json:"intent"`
	MerchantInvoiceNumber string `json:"merchantInvoiceNumber"`
	StatusCode            string `json:"statusCode"`
	StatusMessage         string `json:"statusMessage"`
}

type QueryPaymentResponse struct {
	PaymentID              string `json:"paymentID"`
	Mode                   string `json:"mode"`
	PayerReference         string `json:"payerReference"`
	PaymentCreateTime      string `json:"paymentCreateTime"`
	PaymentExecuteTime     string `json:"paymentExecuteTime"`
	TrxID                  string `json:"trxID"`
	TransactionStatus      string `json:"transaction_status"`
	Amount                 string `json:"amount"`
	Currency               string `json:"currency"`
	Intent                 string `json:"intent"`
	MerchantInvoiceNumber  string `json:"merchantInvoiceNumber"`
	UserVerificationStatus string `json:"userVerificationStatus"`
	StatusCode             string `json:"statusCode"`
	StatusMessage          string `json:"statusMessage"`
}

type SearchTransactionResponse struct {
	Amount                string `json:"amount"`
	CompletedTime         string `json:"completed_time"`
	Currency              string `json:"currency"`
	CustomerMsisdn        string `json:"customerMsisdn"`
	InitiationTime        string `json:"initiationTime"`
	OrganizationShortCode string `json:"organizationShortCode"`
	TransactionReference  string `json:"transactionReference"`
	TransactionStatus     string `json:"transactionStatus"`
	TransactionType       string `json:"transactionType"`
	StatusCode            string `json:"statusCode"`
	StatusMessage         string `json:"statusMessage"`
	TrxID                 string `json:"trxID"`
}

type CreateAgreementValidationResponse struct {
	PaymentID string `json:"paymentID"`
	Status    string `json:"status"`
}
