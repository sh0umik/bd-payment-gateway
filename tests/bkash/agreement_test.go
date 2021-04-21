package bkash

import (
	payment "github.com/sh0umik/go-sslcom"
	"github.com/sh0umik/go-sslcom/models"
	"os"
	"testing"
)

func TestAgreement(t *testing.T) {
	username := os.Getenv("BKASH_USERNAME")
	password := os.Getenv("BKASH_PASSWORD")
	appKey := os.Getenv("BKASH_APP_KEY")
	appSecret := os.Getenv("BKASH_APP_SECRET")

	bkash := payment.GetBkash(username, password, appKey, appSecret)
	paymentService := payment.BkashTokenizedCheckoutService(bkash)

	token, err := paymentService.GrantToken(false)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		t.Error("invalid token")
		t.Fail()
	}

	var createResponse *models.CreateAgreementResponse
	t.Run("test CreateAgreement", func(t *testing.T) {
		req := &models.CreateAgreementRequest{
			Mode:           "0000",
			PayerReference: "01537161343",
			CallbackUrl:    "https://api.shikho.net/payment",
			Currency:       "BDT",
			Intent:         "Shikho Subscription",
		}
		createAgreementResponse, err := paymentService.CreateAgreement(req, token, false)

		if err != nil {
			t.Error(err.Error())
			t.Fail()
		}

		if createAgreementResponse == nil || createAgreementResponse.StatusCode != "0000" {
			t.Fatal("Invalid create agreement response")
		}

		createResponse = createAgreementResponse
	})

	//var executeReponse *models.ExecuteAgreementResponse
	t.Run("test ExecuteAgreement", func(t *testing.T) {
		req := &models.ExecuteAgreementRequest{
			PaymentID: createResponse.PaymentID,
		}
		executeAgreementResponse, err := paymentService.ExecuteAgreement(req, token, false)

		if err != nil {
			t.Error(err.Error())
			t.Fail()
		}

		if executeAgreementResponse == nil || executeAgreementResponse.StatusCode != "0000" {
			t.Fatal("Invalid execute agreement response")
		}

		//executeReponse = executeAgreementResponse
	})
}
