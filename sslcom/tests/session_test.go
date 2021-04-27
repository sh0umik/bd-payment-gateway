package tests

import (
	"github.com/sh0umik/bd-payment-gateway/sslcom"
	models2 "github.com/sh0umik/bd-payment-gateway/sslcom/models"
	"os"
	"testing"
)

func TestCreateSession(t *testing.T) {

	storeId := os.Getenv("SSLCOM_STORE_ID")
	storePass := os.Getenv("SSLCOM_STORE_PASSWORD")

	sslCom := sslcom.GetSslCommerz(storeId, storePass)
	paymentService := sslcom.PaymentService(sslCom)

	sessionRequest := &models2.RequestValue{
		// Fill it in
		TotalAmount:      "1050",
		Currency:         "BDT",
		TranID:           "TRAN123456",
		SuccessURL:       "https://shikho.tech/payment/success",
		FailUrl:          "https://shikho.tech/payment/fail",
		CancelURL:        "https://shikho.tech/payment/cancel",
		CustomerName:     "Customer Name",
		CustomerEmail:    "cus@yahoo.com",
		CustomerAdd1:     "Dhaka",
		CustomerAdd2:     "Dhaka",
		CustomerCity:     "Dhaka",
		CustomerState:    "Dhaka",
		CustomerPostCode: "1000",
		CustomerCountry:  "Bangladesh",
		CustomerPhone:    "01711111111",
		CustomerFax:      "01711111111",
		ShipName:         "Ship Name",
		ShipAdd1:         "Dhaka",
		ShipAdd2:         "Dhaka",
		ShipCity:         "Dhaka",
		ShipState:        "Dhaka",
		ShipPostCode:     "1000",
		ShipCountry:      "Bangladesh",
		MultiCardName:    []string{"bkash"},
		ValueA:           "706633169",
		ValueB:           "288580",
		ValueC: 		  "3",
	}

	sessionResponse, err := paymentService.CreateSession(sessionRequest, false)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(sessionResponse)

	if sessionResponse.Status != "SUCCESS" {
		t.Errorf("Can Not Create a Session")
	}
}
