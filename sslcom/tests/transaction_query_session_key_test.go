package tests

import (
	"github.com/sh0umik/bd-payment-gateway/sslcom"
	models2 "github.com/sh0umik/bd-payment-gateway/sslcom/models"
	"os"
	"testing"
)

/*
Sample request

curl -X POST https://sandbox.sslcommerz.com/validator/api/merchantTransIDvalidationAPI.php?wsdl -d 'tran_id=5a16c6fd8b23783&val_id=1711231900331kHP17lnrr9T8Gt&amount=100&card_type=VISA-Dutch Bangla&store_amount=97&card_no=425272XXXXXX3456&
bank_tran_id=1711231900331S0R8atkhAZksmM&status=VALID&tran_date=2017-11-23 18:59:55&currency=BDT&card_issuer=Standard Chartered Bank&card_brand=VISA&
card_issuer_country=Bangladesh&card_issuer_country_code=BD&store_id=testbox&verify_sign=8070c0cefed9e629b01100d8a92afda2&verify_key=amount,bank_tran_id,base_fair,card_brand,card_issuer,card_issuer_country,card_issuer_country_code,card_no,card_type,currency,currency_amount,currency_rate,currency_type,risk_level,risk_title,status,store_amount,store_id,tran_date,tran_id,val_id,value_a,value_b,value_c,value_d&
cus_fax=01711111111&currency_type=BDT&currency_amount=100.00&currency_rate=1.0000&base_fair=0.00&value_a=ref001_A&value_b=ref002_B&value_c=ref003_C&value_d=ref004_D&
risk_level=0&risk_title=Safe'

*/

func TestTransactionQueryBySID(t *testing.T) {

	storeId := os.Getenv("SSLCOM_STORE_ID")
	storePass := os.Getenv("SSLCOM_STORE_PASSWORD")

	sslCom := sslcom.GetSslCommerz(storeId, storePass)
	paymentService := sslcom.PaymentService(sslCom)

	data := models2.TransactionQueryRequest{
		SessionKey: "017453EB886ffDD96218F50704D588BE4C",
		TranId:     "REF123",
		V:          1,
		Format:     "json",
	}

	tranResponse, err := paymentService.TransactionQueryBySID(&data, false)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(tranResponse)

	if tranResponse.Status != "VALID" {
		t.Logf(" Transaction is %s", tranResponse.Status)
		return
	}

}
