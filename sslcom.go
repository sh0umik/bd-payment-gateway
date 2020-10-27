package go_sslcom

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sh0umik/go-sslcom/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const SANDBOX_GATEWAY = "https://sandbox.sslcommerz.com"
const LIVE_GATEWAY = "https://securepay.sslcommerz.com"
const SESSION_URI = "gwprocess/v3/api.php"
const ORDER_VALIDATION_URI = "validator/api/validationserverAPI.php"
const TRANSACTION_QUERY_URI_TID = "validator/api/merchantTransIDvalidationAPI.php"
const TRANSACTION_QUERY_URI_SK = "validator/api/merchantTransIDvalidationAPI.php"
const REFUNDING_URI = "validator/api/merchantTransIDvalidationAPI.php"

type SslCommerz struct {
	StorePass string
	StoreId   string
}

func GetSslCommerz(storeID string, storePass string) *SslCommerz {
	return &SslCommerz{
		StoreId:   storeID,
		StorePass: storePass,
	}
}

func (s *SslCommerz) CreateSession(req *models.RequestValue, isLiveStore bool) (*models.SessionResponse, error) {

	data := url.Values{}

	data.Set("store_id", s.StoreId)
	data.Set("store_passwd", s.StorePass)
	data.Set("total_amount", req.TotalAmount)
	data.Set("currency", req.Currency)
	data.Set("tran_id", req.TranID)
	data.Set("success_url", req.SuccessURL)
	data.Set("fail_url", req.FailUrl)
	data.Set("cancel_url", req.CancelURL)
	data.Set("cus_name", req.CustomerName)
	data.Set("cus_email", req.CustomerEmail)
	data.Set("cus_add1", req.CustomerAdd1)
	data.Set("cus_add2", req.CustomerAdd2)
	data.Set("cus_city", req.CustomerCity)
	data.Set("cus_state", req.CustomerState)
	data.Set("cus_postcode", req.CustomerPostCode)
	data.Set("cus_country", req.CustomerCountry)
	data.Set("cus_phone", req.CustomerPhone)
	data.Set("cus_fax", req.CustomerFax)
	data.Set("ship_name", req.ShipName)
	data.Set("ship_add1 ", req.ShipAdd1)
	data.Set("ship_add2", req.ShipAdd2)
	data.Set("ship_city", req.ShipCity)
	data.Set("ship_state", req.ShipState)
	data.Set("ship_postcode", req.ShipPostCode)
	data.Set("ship_country", req.ShipCountry)
	data.Set("multi_card_name", strings.Join(req.MultiCardName, ",")) //8
	data.Set("value_a", req.ValueA)
	data.Set("value_b", req.ValueB)
	data.Set("value_c", req.ValueC)
	data.Set("value_d", req.ValueD)

	/*
		data.Set("multi_card_name", "mastercard,visacard,amexcard,brac_visa,dbbl_visa,city_visa,ebl_visa,sbl_visa,brac_master,dbbl_master,city_master,ebl_master,sbl_master,city_amex,qcash,dbbl_nexus,bankasia,abbank,ibbl,mtbl,bkash,dbblmobilebanking,city")
	*/

	var storeUrl string
	if isLiveStore {
		storeUrl = LIVE_GATEWAY
	}else {
		storeUrl = SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path = SESSION_URI
	//u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}
	r, err := http.NewRequest("POST", sessionURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	var resp models.SessionResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *SslCommerz) IPNListener(request *http.Request) (*models.IpnResponse, error) {

	// IPN sends the data to IPN URL as Query Request URI so ,

	var ipnResponse models.IpnResponse
	if err := request.ParseForm(); err != nil {
		return nil, err
	}
	for key, values := range request.PostForm {

		switch key {
		case "tran_id":
			ipnResponse.TranId = values[0]
			break
		case "val_id":
			ipnResponse.ValID = values[0]
			break
		case "amount":
			ipnResponse.Amount = values[0]
			break
		case "card_type":
			ipnResponse.CardType = values[0]
			break
		case "store_amount":
			ipnResponse.StoreAmount = values[0]
			break
		case "card_no":
			ipnResponse.CardNo = values[0]
			break
		case "bank_tran_id":
			ipnResponse.BankTranId = values[0]
			break
		case "status":
			ipnResponse.Status = values[0]
			break
		case "tran_date":
			ipnResponse.TranDate = values[0]
			break
		case "currency":
			ipnResponse.Currency = values[0]
			break
		case "card_issuer":
			ipnResponse.CardIssuer = values[0]
			break
		case "card_brand":
			ipnResponse.CardBrand = values[0]
			break
		case "card_issuer_country":
			ipnResponse.CardIssuerCountry = values[0]
			break
		case "card_issuer_country_code":
			ipnResponse.CardIssuerCountryCode = values[0]
			break
		case "store_id":
			ipnResponse.StoreId = values[0]
			break
		case "verify_sign":
			ipnResponse.VerifySign = values[0]
			break
		case "verify_key":
			ipnResponse.VerifyKey = values[0]
			break
		case "cus_fax":
			ipnResponse.CustomerFax = values[0]
			break
		case "currency_type":
			ipnResponse.CurrencyType = values[0]
			break
		case "currency_amount":
			ipnResponse.CurrencyAmount = values[0]
			break
		case "currency_rate":
			ipnResponse.CurrencyRate = values[0]
			break
		case "base_fair":
			ipnResponse.BaseFair = values[0]
			break
		case "value_a":
			ipnResponse.ValueA = values[0]
			break
		case "value_b":
			ipnResponse.ValueB = values[0]
			break
		case "value_c":
			ipnResponse.ValueC = values[0]
			break
		case "value_d":
			ipnResponse.ValueD = values[0]
			break
		case "risk_level":
			ipnResponse.RiskLevel = values[0]
			break
		case "risk_title":
			ipnResponse.RiskTitle = values[0]
			break
		}
	}

	// Not Needed AT THIS TIME
	/*if ipnResponse.VerifyKey != ""{

	hasher := md5.New()

	// Hash Verification
	*/ /*
		Method to validate the hash
		Catch two POST parameters verify_sign and verify_key.
		Explode the parameter verify_key by comma (,). Here, this parameter contents all the parameters which are returned from SSLCommerz and combination of these value generates the verify_sign value.
		Serialize all these parameters by their name
		Example: value of verify_key (before the serial): key1,key3,key2,key5,key4
		Add the store_passwd and your store password value with verify_key
		Example: Now the verify_key (after adding store_passwd): key1,key3,key2,key5,key4,store_passwd
		Example: value of verify_key (after the serialized): key1,key2,key3,key4,key5,store_passwd
		Make a string by combining the parameters' key and value. Example: key1=value1&key2=value2&key3=value3&key4=value4&key5=value5&store_passwd=Your Store Password
		Generate md5 hash of the value and match with verify_sign
	*/ /*

		keys := strings.Split(ipnResponse.VerifyKey, ",")

		var queryStrings []string
		for _, key := range keys {
			if found, ok := ipnResponse[key]; ok {
				queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, found.([]string)[0]))
			}
		}

		hasher.Write([]byte(s.StorePass))
		passmd5Hash := hex.EncodeToString(hasher.Sum(nil))
		// add store pass as per doc said
		queryStrings = append(queryStrings, fmt.Sprintf("store_passwd=%s", passdsmd5Hash))

		sort.Strings(queryStrings) // sort

		queryString := strings.Join(queryStrings, "&")

		hasher.Write([]byte(queryString))
		md5Hash := hex.EncodeToString(hasher.Sum(nil))

		if ipnResponse["verify_sign"] == md5Hash {
			fmt.Println("verified")
		}

		fmt.Println(queryString)
		fmt.Println(md5Hash)

		return ipnResponse, nil
	}*/

	return &ipnResponse, nil
}

func (s *SslCommerz) OrderValidation(ipnValId string, isLiveStore bool) (*models.IpnResponse, error) {
	data := url.Values{}

	data.Set("val_id", ipnValId)
	data.Set("store_id", s.StoreId)
	data.Set("store_passwd", s.StorePass)
	data.Set("v", "1")
	data.Set("format", "json")

	var storeUrl string
	if isLiveStore {
		storeUrl = LIVE_GATEWAY
	}else {
		storeUrl = SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path = ORDER_VALIDATION_URI
	u.RawQuery = data.Encode()
	sessionURL := u.String()

	client := &http.Client{}
	r, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp models.IpnResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "VALID" || resp.Status == "VALIDATED" {
		return &resp, nil
	}

	return nil, errors.New(fmt.Sprintf("Transaction is not valid : %s", resp.Status))
}

func (s *SslCommerz) CheckValidation(request *models.OrderValidationRequest, isLiveStore bool) (*models.OrderValidationResponse, error) {

	data := url.Values{}

	data.Set("val_id", request.ValId)
	data.Set("store_id", s.StoreId)
	data.Set("store_passwd", s.StorePass)
	data.Set("v", strconv.Itoa(request.V))
	data.Set("format", request.Format)

	var storeUrl string
	if isLiveStore {
		storeUrl = LIVE_GATEWAY
	}else {
		storeUrl = SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path = ORDER_VALIDATION_URI
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}
	r, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	var resp models.OrderValidationResponse
	json.Unmarshal(body, &resp)

	return &resp, nil
}

func (s *SslCommerz) TransactionQueryByTID(request *models.TransactionQueryRequest, isLiveStore bool) (*models.TransactionQueryResponseTID, error) {

	data := url.Values{}

	data.Set("tran_id", request.TranId)
	data.Set("store_id", s.StoreId)
	data.Set("store_passwd", s.StorePass)
	data.Set("v", strconv.Itoa(request.V))
	data.Set("format", request.Format)

	var storeUrl string
	if isLiveStore {
		storeUrl = LIVE_GATEWAY
	}else {
		storeUrl = SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path = TRANSACTION_QUERY_URI_TID
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}

	var resp models.TransactionQueryResponseTID

	r, err := http.NewRequest("GET", sessionURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	json.Unmarshal(body, &resp)

	return &resp, nil
}

func (s *SslCommerz) TransactionQueryBySID(request *models.TransactionQueryRequest, isLiveStore bool) (*models.TransactionQueryResponseSID, error) {

	data := url.Values{}

	data.Set("sessionkey", request.SessionKey)
	data.Set("store_id", s.StoreId)
	data.Set("store_passwd", s.StorePass)
	//data.Set("v", request.V)
	//data.Set("format", request.Format)

	var storeUrl string
	if isLiveStore {
		storeUrl = LIVE_GATEWAY
	}else {
		storeUrl = SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path = TRANSACTION_QUERY_URI_SK
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}

	var resp models.TransactionQueryResponseSID

	r, err := http.NewRequest("GET", sessionURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	json.Unmarshal(body, &resp)

	return &resp, nil
}

func (s *SslCommerz) InitiateRefunding(request *models.RefundApiRequest, isLiveStore bool) (models.RefundResponse, error) {

	data := url.Values{}

	data.Set("bank_tran_id", request.BankTranId)
	data.Set("store_id", s.StoreId)
	data.Set("store_passwd", s.StorePass)
	data.Set("refund_amount", floattostr(request.RefundAmount))
	data.Set("refund_remarks", request.RefundRemarks)
	data.Set("refe_id", request.RefId)
	data.Set("format", request.Format)

	var storeUrl string
	if isLiveStore {
		storeUrl = LIVE_GATEWAY
	}else {
		storeUrl = SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path = REFUNDING_URI
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}

	var resp models.RefundResponse

	r, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return resp, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return resp, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	json.Unmarshal(body, &resp)

	return resp, nil

}

func floattostr(input_num float64) string {
	return strconv.FormatFloat(input_num, 'g', 1, 64)
}
