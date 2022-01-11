package paymentsiface

import "github.com/nexusmed/nexusmed-go/payments"

type PaymentsAPI interface {
	// Create a new payment
	CreatePayment(payments.CreatePaymentInput) (*payments.CreatePayment, error)

	// Set the Authorization header
	SetApiKey(key string)
}
