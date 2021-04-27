package tests

import (
	"github.com/sh0umik/bd-payment-gateway/bkash"
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
