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

type BkashTokenizedCheckoutService interface {
	// GrantToken creates a access token using bkash credentials
	GrantToken(isLiveStore bool) (*models.Token, error)

	// RefreshToken refreshes the access token
	RefreshToken(token *models.Token, isLiveStore bool) (*models.Token, error)

	// CreateAgreement Initiates an agreement request for a customer.
	// Pass phone number of the customer as PayerReference
	CreateAgreement(request *models.CreateAgreementRequest, token *models.Token, isLiveStore bool) (*models.CreateAgreementResponse, error)

	// CreateAgreementValidationListener is a handler func that receives paymentID & status
	// as a json post request and returns CreateAgreementValidationResponse object
	CreateAgreementValidationListener(r *http.Request) (*models.CreateAgreementValidationResponse, error)

	// ExecuteAgreement executes the agreement using the paymentID received from CreateAgreementValidationResponse
	ExecuteAgreement(request *models.ExecuteAgreementRequest, token *models.Token, isLiveStore bool) (*models.ExecuteAgreementResponse, error)

	// QueryAgreement query agreement by agreementID
	QueryAgreement(request *models.QueryAgreementRequest, token *models.Token, isLiveStore bool) (*models.QueryAgreementResponse, error)

	// CancelAgreement cancels an agreement by agreementID
	CancelAgreement(request *models.CancelAgreementRequest, token *models.Token, isLiveStore bool) (*models.CancelAgreementResponse, error)
}
