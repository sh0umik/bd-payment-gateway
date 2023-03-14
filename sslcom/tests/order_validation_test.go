package tests

import (
	"github.com/sh0umik/bd-payment-gateway/sslcom"
	models2 "github.com/sh0umik/bd-payment-gateway/sslcom/models"
	"testing"
)

// Sample Request
/*
curl -X POST https://sandbox.sslcommerz.com/gwprocess/v3/api.php -d 'store_id=testbox&store_passwd=testpass&total_amount=100&currency=EUR&tran_id=REF123&success_url=http://yoursite.com/success.php&fail_url=http://yoursite.com/fail.php&cancel_url=http://yoursite.com/cancel.php&cus_name=Customer Name&cus_email=cust@yahoo.com&cus_add1=Dhaka&cus_add2=Dhaka&cus_city=Dhaka&cus_state=Dhaka&cus_postcode=1000&cus_country=Bangladesh&cus_phone=01711111111&cus_fax=01711111111&
ship_name=Customer Name&ship_add1 =Dhaka&ship_add2=Dhaka&ship_city=Dhaka&ship_state=Dhaka&ship_postcode=1000&ship_country=Bangladesh&multi_card_name=mastercard,visacard,amexcard&value_a=ref001_A&value_b=ref002_B&
value_c=ref003_C&value_d=ref004_D'
*/

func TestOrderValidation(t *testing.T) {

	storeId := "shikh5f43ed6eb1bc8"
	storePass := "shikh5f43ed6eb1bc8@ssl"

	sslCom := sslcom.GetSslCommerz(storeId, storePass, false)
	paymentService := sslcom.PaymentService(sslCom)

	data := models2.OrderValidationRequest{
		ValId:  "230314111712tvDmtdz89RYvRhG",
		Format: "json",
		V:      1,
	}

	orderResponse, err := paymentService.CheckValidation(&data)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(orderResponse)

	// #todo check response and error
	if orderResponse.Status != "VALID" && orderResponse.Status != "VALIDATED" {
		t.Errorf("Order validation not successful")
	}

	orderResp, err := paymentService.OrderValidation("230314111712tvDmtdz89RYvRhG")
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(orderResp)

	// #todo check response and error
	if orderResp.Status != "VALID" && orderResp.Status != "VALIDATED" {
		t.Errorf("Order validation not successful")
	}
}
