package tests

import (
	"encoding/json"
	"github.com/sh0umik/bd-payment-gateway/bkash"
	"github.com/sh0umik/bd-payment-gateway/bkash/models"
	"testing"
)

func TestMessageSignatureValidation(t *testing.T) {
	notificationJson := `
	{
	  "Type": "Notification",
	  "MessageId": "c92ab22a-0aad-5643-a7c5-eb34dbe36d7d",
	  "Token": "",
	  "TopicArn": "arn:aws:sns:ap-southeast-1:988117108360:MockBkash",
	  "Message": "{\n       \"dateTime\":\"20180419122246\",\n       \"debitMSISDN\":\"8801700000001\",\n       \"creditOrganizationName\":\"Org 01\",\n       \"creditShortCode\":\"ORG001\",\n       \"trxID\":\"4J420ANOXC\",\n       \"transactionStatus\":\"Completed\",\n       \"transactionType\":\"1003\",\n       \"amount\":\"100\",\n       \"currency\":\"BDT\",\n       \"transactionReference\":\"User inputed reference value.\" \n}",
	  "Timestamp": "2021-05-19T08:40:46.447Z",
	  "SignatureVersion": "1",
	  "Signature": "wqjA4qxKUJMUPKPR9KFMMsSA4F/66+ok2huJ/voad+zbI7aCd2lBR8pyYem6HizSfGGodFYgcBvj2u+UblNKKwXuzBWl//keUlkbEyq6WURuID4k9m4BE/By010v3HQmnuVDfArJD0PsnhUtnPbXL6031+nFr/h2A1lsuyqON4L9SN6n2WRzNHOMMZW6H6vIYXyrpSBN7f0ruut4jYvwGJOzy4fEAyrNForOHm9cX8YyWsZGiFMy0p9IhLt3hcTM5d1j7Ld49gDhV6f0TwGzG8vmaGN7776N3GNj6CiGRD3aG8LHeU0CSlIQndOrmKen8IBnLnrq372VSboAacLnYg==",
	  "SigningCertURL": "https://sns.ap-southeast-1.amazonaws.com/SimpleNotificationService-010a507c1833636cd94bdb98bd93083a.pem",
	  "UnsubscribeURL": "https://sns.ap-southeast-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:ap-southeast-1:988117108360:MockBkash:57bbfe14-f761-4bfc-9a5e-29ee932765f4",
	  "SubscribeURL": "",
	  "Subject": ""
	}`
	var notificationPayload models.BkashIPNPayload

	err := json.Unmarshal([]byte(notificationJson), &notificationPayload)
	if err != nil {
		t.Fatal(err)
	}

	verifyErr := bkash.IsMessageSignatureValid(&notificationPayload)
	if verifyErr != nil {
		t.Fatal(verifyErr)
	}
	t.Log("valid payload")
}
