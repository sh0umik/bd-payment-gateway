package tests

import (
	"github.com/gin-gonic/gin"
	payment "github.com/sh0umik/go-sslcom"
	"os"
	"testing"
)

/*
Sample Request

curl -X POST localhost:8080/ipn -d 'tran_id=5a16c68b23783&val_id=1711231900331kHP17lnrr9T8Gt&amount=100&card_type=VISA-Dutch Bangla&store_amount=97&card_no=425272XXXXXX3456&bank_tran_id=1711231900331S0R8atkhAZksmM&status=VALID&tran_date=2017-11-23 18:59:55&currency=BDT&card_issuer=Standard Chartered Bank&card_brand=VISA&card_issuer_country=Bangladesh&card_issuer_country_code=BD&store_id=testbox&verify_sign=8070c0cefed9e629b01100d8a92afda2&verify_key=amount,bank_tran_id,base_fair,card_brand,card_issuer,card_issuer_country,card_issuer_country_code,card_no,card_type,currency,currency_amount,currency_rate,currency_type,risk_level,risk_title,status,store_amount,store_id,tran_date,tran_id,val_id,value_a,value_b,value_c,value_d&cus_fax=01711111111&currency_type=BDT&currency_amount=100.00&currency_rate=1.0000&base_fair=0.00&value_a=ref001_A&value_b=ref002_B&value_c=ref003_C&value_d=ref004_D&risk_level=0&risk_title=Safe'

*/
/*
	Sample IPN POST Request Response
	2/12/2018 1:06:13 PM INI IPN
	2/12/2018 1:06:13 PM verify_key [amount,bank_tran_id,base_fair,card_brand,card_issuer,card_issuer_country,card_issuer_country_code,card_no,card_type,currency,currency_amount,currency_rate,currency_type,risk_level,risk_title,status,store_amount,store_id,tran_date,tran_id,val_id,value_a,value_b,value_c,value_d]
	2/12/2018 1:06:13 PM card_issuer [BRAC BANK, LTD.]
	2/12/2018 1:06:13 PM tran_date [2018-02-12 13:00:49]
	2/12/2018 1:06:13 PM value_a [ref001_A]
	2/12/2018 1:06:13 PM risk_title [Safe]
	2/12/2018 1:06:13 PM status [VALID]
	2/12/2018 1:06:13 PM card_issuer_country_code [BD]
	2/12/2018 1:06:13 PM currency_amount [150.00]
	2/12/2018 1:06:13 PM risk_level [0]
	2/12/2018 1:06:13 PM store_amount [145.5]
	2/12/2018 1:06:13 PM val_id [180212130454LSoAaVvP591wES1]
	2/12/2018 1:06:13 PM value_d [ref004_D]
	2/12/2018 1:06:13 PM amount [150]
	2/12/2018 1:06:13 PM card_brand [MASTERCARD]
	2/12/2018 1:06:13 PM verify_sign_sha2 [933e92c724f9bd85d2423898cea9de1e6b549f273eebcfb95da769f864ad2634]
	2/12/2018 1:06:13 PM currency_rate [1.0000]
	2/12/2018 1:06:13 PM verify_sign [36f3f9b7231964216fc5fd7a6518b5ba]
	2/12/2018 1:06:13 PM store_id [test_my_ass]
	2/12/2018 1:06:13 PM base_fair [0.00]
	2/12/2018 1:06:13 PM currency [BDT]
	2/12/2018 1:06:13 PM card_issuer_country [Bangladesh]
	2/12/2018 1:06:13 PM value_c [ref003_C]
	2/12/2018 1:06:13 PM card_no [545610XXXXXX4362]
	2/12/2018 1:06:13 PM tran_id [CY123]
	2/12/2018 1:06:13 PM currency_type [BDT]
	2/12/2018 1:06:13 PM value_b [ref002_B]
	2/12/2018 1:06:13 PM bank_tran_id [180212130454TulpU0ADylAw4ci]
*/

func TestIPN(t *testing.T) {

	storeId := os.Getenv("SSLCOM_STORE_ID")
	storePass := os.Getenv("SSLCOM_STORE_PASSWORD")

	sslCom := payment.GetSslCommerz(storeId, storePass)

	router := gin.Default()

	router.POST("/ipn", func(c *gin.Context) {

		t.Log(c.Request.Form)
		t.Log(c.Request.URL.Query())

		t.Logf("%+v", c.Request)

		ipnResp, err := sslCom.IPNListener(c.Request)
		if err != nil {
			t.Error(err.Error())
		}

		t.Log("IPN TESTING RESPONSE : ", ipnResp)

	})
	router.Run(":8080")
}
