package tests

import (
	payment "github.com/sh0umik/go-sslcom"
	"github.com/sh0umik/go-sslcom/models"
	"os"
	"testing"
)

func TestTransactionQueryByTranID(t *testing.T) {

	storeId := os.Getenv("SSLCOM_STORE_ID")
	storePass := os.Getenv("SSLCOM_STORE_PASSWORD")

	sslCom := payment.GetSslCommerz(storeId, storePass)
	paymentService := payment.PaymentService(sslCom)

	data := models.TransactionQueryRequest{
		SessionKey: "73E5E656F8EFE6665fdf08A9165349E67C7",
		TranId:     "59C2A4F6432F8",
		V:          1,
		Format:     "json",
	}

	tranResponse, err := paymentService.TransactionQueryByTID(&data)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(tranResponse)

}
