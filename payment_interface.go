package go_sslcom

import (
	"github.com/sh0umik/go-sslcom/models"
	"net/http"
)

type PaymentService interface {

	// Create session
	CreateSession(value *models.RequestValue, isLiveStore bool) (*models.SessionResponse, error)

	// Set up IPN Listener
	IPNListener(request *http.Request) (*models.IpnResponse, error)

	// Validate the IPN Response
	OrderValidation(valId string, isLiveStore bool) (*models.IpnResponse, error)

	// 	Order Validation
	CheckValidation(request *models.OrderValidationRequest, isLiveStore bool) (*models.OrderValidationResponse, error)

	// Transaction query by Transaction ID
	TransactionQueryByTID(request *models.TransactionQueryRequest, isLiveStore bool) (*models.TransactionQueryResponseTID, error)

	//Transaction query by Session Key
	TransactionQueryBySID(request *models.TransactionQueryRequest, isLiveStore bool) (*models.TransactionQueryResponseSID, error)
}
