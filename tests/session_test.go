package tests

import (
	payment "github.com/sh0umik/go-sslcom"
	"github.com/sh0umik/go-sslcom/models"
	"os"
	"testing"
)

func TestCreateSession(t *testing.T) {

	storeId := os.Getenv("SSLCOM_STORE_ID")
	storePass := os.Getenv("SSLCOM_STORE_PASSWORD")

	sslCom := payment.GetSslCommerz(storeId, storePass)
	paymentService := payment.PaymentService(sslCom)

	sessionRequest := &models.RequestValue{
		// Fill it in
		TotalAmount:      "1050",
		Currency:         "BDT",
		TranID:           "TRAN123456",
		SuccessURL:       "http://yoursite.com/success.php",
		FailUrl:          "http://yoursite.com/fail.php",
		CancelURL:        "http://yoursite.com/cancel.php",
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
