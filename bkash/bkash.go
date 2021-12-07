package bkash

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/sh0umik/bd-payment-gateway/bkash/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const BKASH_SANDBOX_GATEWAY = "https://tokenized.sandbox.bka.sh/v1.2.0-beta"
const BKASH_LIVE_GATEWAY = "https://tokenized.pay.bka.sh/v1.2.0-beta"
const BKASH_GRANT_TOKEN_URI = "/tokenized/checkout/token/grant"
const BKASH_REFRESH_TOKEN_URI = "/tokenized/checkout/token/refresh"
const BKASH_CREATE_AGREEMENT_URI = "/tokenized/checkout/create"
const BKASH_EXECUTE_AGREEMENT_URI = "/tokenized/checkout/execute"
const BKASH_QUERY_AGREEMENT_URI = "/tokenized/checkout/agreement/status"
const BKASH_CANCEL_AGREEMENT_URI = "/tokenized/checkout/agreement/cancel"
const BKASH_CREATE_PAYMENT_URI = "/tokenized/checkout/create"
const BKASH_EXECUTE_PAYMENT_URI = "/tokenized/checkout/execute"
const BKASH_QUERY_PAYMENT_URI = "/tokenized/checkout/payment/status"

var EMPTY_REQUIRED_FIELD = errors.New("empty required field")
var TIMEOUT_ERROR = errors.New("api request timeout")

type Bkash struct {
	Username  string
	Password  string
	AppKey    string
	AppSecret string
}

func GetBkash(username, password, appKey, appSecret string) *Bkash {
	return &Bkash{
		Username:  username,
		Password:  password,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (b *Bkash) GrantToken(isLiveStore bool) (*models.Token, error) {
	// Mandatory field validation
	if b.AppKey == "" || b.AppSecret == "" || b.Username == "" || b.Password == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var data = make(map[string]string)

	data["app_key"] = b.AppKey
	data["app_secret"] = b.AppSecret

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_GRANT_TOKEN_URI

	grantTokenURL := u.String()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", grantTokenURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("username", b.Username)
	r.Header.Add("password", b.Password)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.Token
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) RefreshToken(token *models.Token, isLiveStore bool) (*models.Token, error) {
	// Mandatory field validation
	if b.AppKey == "" || b.AppSecret == "" || token.RefreshToken == "" || b.Username == "" || b.Password == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var data = make(map[string]string)

	data["app_key"] = b.AppKey
	data["app_secret"] = b.AppSecret
	data["refresh_token"] = token.RefreshToken

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_REFRESH_TOKEN_URI

	refreshTokenURL := u.String()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", refreshTokenURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("username", b.Username)
	r.Header.Add("password", b.Password)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.Token
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) CreateAgreement(request *models.CreateAgreementRequest, token *models.Token, isLiveStore bool) (*models.CreateAgreementResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.Mode == "" || request.CallbackUrl == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	// Mode validation
	if request.Mode != "0000" {
		return nil, errors.New("invalid mode value")
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_CREATE_AGREEMENT_URI
	//u.RawQuery = data.Encode()

	createAgreementURL := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", createAgreementURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.CreateAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) CreateAgreementValidationListener(r *http.Request) (*models.CreateAgreementValidationResponse, error) {
	if r.Method != "POST" {
		return nil, errors.New("method not allowed")
	}

	var agreementTValidationResponse models.CreateAgreementValidationResponse

	err := json.NewDecoder(r.Body).Decode(&agreementTValidationResponse)
	if err != nil {
		return nil, err
	}

	return &agreementTValidationResponse, nil
}

func (b *Bkash) ExecuteAgreement(request *models.ExecuteAgreementRequest, token *models.Token, isLiveStore bool) (*models.ExecuteAgreementResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_EXECUTE_AGREEMENT_URI
	//u.RawQuery = data.Encode()

	executeAgreementURL := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", executeAgreementURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.ExecuteAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) QueryAgreement(request *models.QueryAgreementRequest, token *models.Token, isLiveStore bool) (*models.QueryAgreementResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.AgreementID == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_QUERY_AGREEMENT_URI
	//u.RawQuery = data.Encode()

	queryAgreementURL := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", queryAgreementURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.QueryAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) CancelAgreement(request *models.CancelAgreementRequest, token *models.Token, isLiveStore bool) (*models.CancelAgreementResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.AgreementID == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_CANCEL_AGREEMENT_URI
	//u.RawQuery = data.Encode()

	cancelAgreementURL := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", cancelAgreementURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.CancelAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) CreatePayment(request *models.CreatePaymentRequest, token *models.Token, isLiveStore bool) (*models.CreatePaymentResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.Mode == "" || request.CallbackURL == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	// Mode validation
	if request.Mode != "0001" && request.Mode != "0011" {
		return nil, errors.New("invalid mode value")
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_CREATE_PAYMENT_URI

	createPaymentURL := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", createPaymentURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.CreatePaymentResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) ExecutePayment(request *models.ExecutePaymentRequest, token *models.Token, isLiveStore, debug bool) (*models.ExecutePaymentResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_EXECUTE_PAYMENT_URI

	executePayment := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", executePayment, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	if debug {
		time.Sleep(time.Second * 40)
	}

	response, err := client.Do(r)
	if err != nil {
		// if error is timeout then call query payment
		// if complete return success payload (*models.ExecutePaymentResponse, nil)
		// if initiated - return something that should be handled by client (maybe return some kind of timeout error)
		if errors.Is(err, context.DeadlineExceeded) {
			queryResp, err := b.QueryPayment(&models.QueryPaymentRequest{PaymentID: request.PaymentID}, token, isLiveStore)
			if err != nil {
				return nil, err
			}

			if queryResp.StatusCode == "0000" && queryResp.TransactionStatus == "Completed" {
				return &models.ExecutePaymentResponse{
					PaymentID:             queryResp.PaymentID,
					PayerReference:        queryResp.PayerReference,
					PaymentExecuteTime:    queryResp.PaymentExecuteTime,
					TrxID:                 queryResp.TrxID,
					TransactionStatus:     queryResp.TransactionStatus,
					Amount:                queryResp.Amount,
					Currency:              queryResp.Currency,
					Intent:                queryResp.Intent,
					MerchantInvoiceNumber: queryResp.MerchantInvoiceNumber,
					StatusCode:            queryResp.StatusCode,
					StatusMessage:         queryResp.StatusMessage,
					//AgreementID:           "",
					//CustomerMsisdn:        "",
					//AgreementExecuteTime:  "",
					//AgreementStatus:       "",
				}, nil
			} else {
				return nil, TIMEOUT_ERROR
			}
		} else {
			return nil, err
		}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.ExecutePaymentResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (b *Bkash) QueryPayment(request *models.QueryPaymentRequest, token *models.Token, isLiveStore bool) (*models.QueryPaymentResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, EMPTY_REQUIRED_FIELD
	}

	var storeUrl string
	if isLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_QUERY_PAYMENT_URI
	//u.RawQuery = data.Encode()

	queryPaymentURL := u.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", queryPaymentURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.IdToken))
	r.Header.Add("X-APP-Key", b.AppKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.QueryPaymentResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// getMessageBytesToSign returns a byte array containing a signature usable for signature verification
func getMessageBytesToSign(msg *models.BkashIPNPayload) []byte {
	var builtSignature bytes.Buffer
	signableKeys := []string{"Message", "MessageId", "Subject", "SubscribeURL", "Timestamp", "Token", "TopicArn", "Type"}
	for _, key := range signableKeys {
		reflectedStruct := reflect.ValueOf(msg)
		field := reflect.Indirect(reflectedStruct).FieldByName(key)
		value := field.String()
		if field.IsValid() && value != "" {
			builtSignature.WriteString(key + "\n")
			builtSignature.WriteString(value + "\n")
		}
	}
	return builtSignature.Bytes()
}

// IsMessageSignatureValid validates bkash IPN message signature. Returns true, nil if ok,
// otherwise returns false, error
func IsMessageSignatureValid(msg *models.BkashIPNPayload) error {
	resp, err := http.Get(msg.SigningCertURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("unable to get certificate err: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	p, _ := pem.Decode(body)
	cert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return err
	}

	base64DecodedSignature, err := base64.StdEncoding.DecodeString(msg.Signature)
	if err != nil {
		return err
	}

	if err := cert.CheckSignature(x509.SHA1WithRSA, getMessageBytesToSign(msg), base64DecodedSignature); err != nil {
		return err
	}

	return nil
}
