package tests

//
//import (
//	"github.com/sh0umik/go-sslcom"
//	"github.com/sh0umik/go-sslcom/models"
//	"testing"
//)
//
//func TestRefundAPI(t *testing.T) {
//
//	sslCom := payment.GetSslCommerz("testbox", "test")
//
//	paymentService := payment.PaymentService(sslCom)
//
//	data := models.RefundApiRequest{
//		BankTranId:    "123454899",
//		RefundAmount:  5.50,
//		RefundRemarks: "Remarks string",
//		RefId:         "Sample referance",
//		Format:        "json",
//	}
//
//	refundResp, err := paymentService.InitiateRefunding(&data)
//	if err != nil {
//		t.Logf(err.Error())
//	}
//
//	t.Log(refundResp)
//
//	// #todo check response and error
//	if refundResp.Status != "success" {
//		t.Errorf("Refund request is failed")
//	}
//
//}
