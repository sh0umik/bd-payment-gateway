package tests

import (
	"github.com/sh0umik/bd-payment-gateway/sslcom"
	models2 "github.com/sh0umik/bd-payment-gateway/sslcom/models"
	"os"
	"testing"
)

func TestTransactionQueryByTranID(t *testing.T) {

	storeId := os.Getenv("SSLCOM_STORE_ID")
	storePass := os.Getenv("SSLCOM_STORE_PASSWORD")

	sslCom := sslcom.GetSslCommerz(storeId, storePass)
	paymentService := sslcom.PaymentService(sslCom)

	data := models2.TransactionQueryRequest{
		SessionKey: "73E5E656F8EFE6665fdf08A9165349E67C7",
		TranId:     "59C2A4F6432F8",
		V:          1,
		Format:     "json",
	}

	tranResponse, err := paymentService.TransactionQueryByTID(&data, false)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(tranResponse)

}
