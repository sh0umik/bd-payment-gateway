package tests

import (
	"github.com/sh0umik/bd-payment-gateway/bkash"
	"github.com/sh0umik/bd-payment-gateway/bkash/models"
	"os"
	"testing"
)

func TestGrantToken(t *testing.T) {
	username := os.Getenv("BKASH_USERNAME")
	password := os.Getenv("BKASH_PASSWORD")
	appKey := os.Getenv("BKASH_APP_KEY")
	appSecret := os.Getenv("BKASH_APP_SECRET")

	bkashService := bkash.GetBkash(username, password, appKey, appSecret)
	paymentService := bkash.BkashTokenizedCheckoutService(bkashService)

	token, err := paymentService.GrantToken(false)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		t.Error("invalid token")
	}
}

func TestBkashJourney(t *testing.T) {
	username := os.Getenv("BKASH_USERNAME")
	password := os.Getenv("BKASH_PASSWORD")
	appKey := os.Getenv("BKASH_APP_KEY")
	appSecret := os.Getenv("BKASH_APP_SECRET")

	bkashService := bkash.GetBkash(username, password, appKey, appSecret)
	paymentService := bkash.BkashTokenizedCheckoutService(bkashService)

	token, err := paymentService.GrantToken(false)
	if err != nil {
		t.Fatal(err)
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		t.Fatalf("StatusCode: %v, Status Message: %v\n", token.StatusCode, token.StatusMessage)
	}

	queryPaymentRes, err := paymentService.QueryPayment(&models.QueryPaymentRequest{PaymentID: "TR0001541629364022912"}, token, false)
	if err != nil {
		t.Fatal(err)
	}

	if queryPaymentRes.StatusCode != "0000" {
		t.Fatalf("StatusCode: %v, Status Message: %v\n", queryPaymentRes.StatusCode, queryPaymentRes.StatusMessage)
	}
}

func getEnv(key, defaults string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaults
}
