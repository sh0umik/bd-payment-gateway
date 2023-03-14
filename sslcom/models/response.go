package models

type SessionResponse struct {
	Status                   string `json:"status"`
	Failedreason             string `json:"failedreason"`
	Sessionkey               string `json:"sessionkey"`
	Gw                       Gw     `json:"gw"`
	RedirectGatewayURL       string `json:"redirectGatewayURL"`
	DirectPaymentURLBank     string `json:"directPaymentURLBank"`
	DirectPaymentURLCard     string `json:"directPaymentURLCard"`
	DirectPaymentURL         string `json:"directPaymentURL"`
	RedirectGatewayURLFailed string `json:"redirectGatewayURLFailed"`
	GatewayPageURL           string `json:"GatewayPageURL"`
	StoreBanner              string `json:"storeBanner"`
	StoreLogo                string `json:"storeLogo"`
	Desc                     []Desc `json:"desc"`
	IsDirectPayEnable        string `json:"is_direct_pay_enable"`
}

type Gw struct {
	Visa            string `json:"visa"`
	Master          string `json:"master"`
	Amex            string `json:"amex"`
	Othercards      string `json:"othercards"`
	Internetbanking string `json:"internetbanking"`
	Mobilebanking   string `json:"mobilebanking"`
}

type Desc struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
	Gw   string `json:"gw"`
}

type IpnResponse struct {
	VerifyKey             string `json:"verify_key"`
	CardIssuer            string `json:"card_issuer"`
	TranDate              string `json:"tran_date"`
	ValueA                string `json:"value_a"`
	ValueB                string `json:"value_b"`
	ValueC                string `json:"value_c"`
	ValueD                string `json:"value_d"`
	RiskTitle             string `json:"risk_title"`
	Code                  string `json:"code"`
	CardIssuerCountryCode string `json:"card_issuer_country_code"`
	CurrencyAmount        string `json:"currency_amount"`
	CurrencyRate          string `json:"currency_rate"`
	RiskLevel             string `json:"risk_level"`
	StoreAmount           string `json:"store_amount"`
	ValID                 string `json:"val_id"`
	Amount                string `json:"amount"`
	CardBrand             string `json:"card_brand"`
	VerifySignSha2        string `json:"verify_sign_sha_2"`
	CrrencyRate           string `json:"crrency_rate"`
	VerifySign            string `json:"verify_sign"`
	StoreId               string `json:"store_id"`
	BaseFair              string `json:"base_fair"`
	Currency              string `json:"currency"`
	CardIssuerCountry     string `json:"card_issuer_country"`
	CardNo                string `json:"card_no"`
	TranId                string `json:"tran_id"`
	CurrencyType          string `json:"currency_type"`
	BankTranId            string `json:"bank_tran_id"`
	CardType              string `json:"card_type"`
	Status                string `json:"status"`
	CustomerFax           string `json:"cus_fax"`
}

type OrderValidationResponse struct {
	Status                     string  `json:"status"`
	TranDate                   string  `json:"tran_date"`
	TranId                     string  `json:"tran_id"`
	ValId                      string  `json:"val_id"`
	Amount                     string  `json:"amount"`
	StoreAmount                string  `json:"store_amount"`
	Currency                   string  `json:"currency"`
	BankTranId                 string  `json:"bank_tran_id"`
	CardType                   string  `json:"card_type"`
	CardNo                     string  `json:"card_no"`
	CardIssuer                 string  `json:"card_issuer"`
	CardBrand                  string  `json:"card_brand"`
	CardCategory               string  `json:"card_category"`
	CardSubBrand               string  `json:"card_sub_brand"`
	CardIssuerCountry          string  `json:"card_issuer_country"`
	CardIssuerCountryCode      string  `json:"card_issuer_country_code"`
	CurrencyType               string  `json:"currency_type"`
	CurrencyAmount             string  `json:"currency_amount"`
	CurrencyRate               string  `json:"currency_rate"`
	AdditionCharge             string  `json:"addition_charge"`
	BaseFair                   string  `json:"base_fair"`
	ValueA                     string  `json:"value_a"`
	ValueB                     string  `json:"value_b"`
	ValueC                     string  `json:"value_c"`
	ValueD                     string  `json:"value_d"`
	EmiInstalment              string  `json:"emi_instalment"`
	EmiAmount                  string  `json:"emi_amount"`
	EmiDescription             string  `json:"emi_description"`
	EmiIssuer                  string  `json:"emi_issuer"`
	RiskTitle                  string  `json:"risk_title"`
	RiskLevel                  string  `json:"risk_level"`
	DiscountPercentage         string  `json:"discount_percentage"`
	DiscountRemarks            string  `json:"discount_remarks"`
	DiscountAmount             float64 `json:"discount_amount,string"`
	APIConnect                 string  `json:"APIConnect"`
	ValidatedOn                string  `json:"validated_on"`
	GwVersion                  string  `json:"gw_version"`
	NoSameCardTransWithinOffer int     `json:"no_same_card_trans_within_offer"`
	PaymentChannel             string  `json:"payment_channel"`
	CardRefId                  string  `json:"card_ref_id"`
	CusName                    string  `json:"cus_name"`
	CusEmail                   string  `json:"cus_email"`
	CusPhone                   string  `json:"cus_phone"`
	CampaignCode               string  `json:"campaign_code"`
}

type TransactionQueryResponseTID struct {
	APIConnect     string           `json:"APIConnect"`
	NoOfTransFound int              `json:"no_of_trans_found"`
	Element        []ElementDetails `json:"element"`
}

type ElementDetails struct {
	ValID                 string `json:"val_id"`
	Status                string `json:"status"`
	ValidatedOn           string `json:"validated_on"`
	CurrencyType          string `json:"currency_type"`
	CurrencyAmount        string `json:"currency_amount"`
	CurrencyRate          string `json:"currency_rate"`
	BaseFair              string `json:"base_fair"`
	ValueA                string `json:"value_a"`
	ValueB                string `json:"value_b"`
	ValueC                string `json:"value_c"`
	ValueD                string `json:"value_d"`
	TranDate              string `json:"tran_date"`
	TranID                string `json:"tran_id"`
	Amount                string `json:"amount"`
	StoreAmount           string `json:"store_amount"`
	BankTranID            string `json:"bank_tran_id"`
	CardType              string `json:"card_type"`
	RiskTitle             string `json:"risk_title"`
	RiskLevel             string `json:"risk_level"`
	Currency              string `json:"currency"`
	BankGw                string `json:"bank_gw"`
	CardNo                string `json:"card_no"`
	CardIssuer            string `json:"card_issuer"`
	CardBrand             string `json:"card_brand"`
	CardIssuerCountry     string `json:"card_issuer_country"`
	CardIssuerCountryCode string `json:"card_issuer_country_code"`
	GwVersion             string `json:"gw_version"`
	EmiInstalment         string `json:"emi_instalment"`
	EmiAmount             string `json:"emi_amount"`
	EmiDescription        string `json:"emi_description"`
	EmiIssuer             string `json:"emi_issuer"`
	Error                 string `json:"error"`
}

type TransactionQueryResponseSID struct {
	APIConnect            string `json:"api_connect"`
	Status                string `json:"status"`
	Sessionkey            string `json:"sessionkey"`
	TranDate              string `json:"tran_date"`
	TranID                string `json:"tran_id"`
	ValID                 string `json:"val_id"`
	Amount                string `json:"amount"`
	StoreAmount           string `json:"store_amount"`
	BankTranID            string `json:"bank_tran_id"`
	CardType              string `json:"card_type"`
	CardNo                string `json:"card_no"`
	CardIssuer            string `json:"card_issuer"`
	CardBrand             string `json:"card_brand"`
	CardIssuerCountry     string `json:"card_issuer_country"`
	CardIssuerCountryCode string `json:"card_issuer_country_code"`
	CurrencyType          string `json:"currency_type"`
	CurrencyAmount        string `json:"currency_amount"`
	CurrencyRate          string `json:"currency_rate"`
	BaseFair              string `json:"base_fair"`
	ValueA                string `json:"value_a"`
	ValueB                string `json:"value_b"`
	ValueC                string `json:"value_c"`
	ValueD                string `json:"value_d"`
	RiskTitle             string `json:"risk_title"`
	RiskLevel             string `json:"risk_level"`
	ValidatedOn           string `json:"validated_on"`
	GwVersion             string `json:"gw_version"`
}

type RefundResponse struct {
	APIConnect  string `json:"APIConnect"`
	BankTranID  string `json:"bank_tran_id"`
	TransID     string `json:"trans_id"`
	RefundRefID string `json:"refund_ref_id"`
	Status      string `json:"status"`
	ErrorReason string `json:"errorReason"`
}
