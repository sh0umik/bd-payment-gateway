package models

type WebhookData struct {
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	Token            string `json:"Token"`
	TopicArn         string `json:"TopicArn"`
	Message          string `json:"Message"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
	SubscribeURL     string `json:"SubscribeURL"`
	Subject          string `json:"Subject"`
}

type Message struct {
	DateTime               string `json:"dateTime"`
	DebitMSISDN            string `json:"debitMSISDN"`
	CreditOrganizationName string `json:"creditOrganizationName"`
	CreditShortCode        string `json:"creditShortCode"`
	TrxID                  string `json:"trxID"`
	TransactionStatus      string `json:"transactionStatus"`
	TransactionType        string `json:"transactionType"`
	Amount                 string `json:"amount"`
	Currency               string `json:"currency"`
	TransactionReference   string `json:"transactionReference"`
}
