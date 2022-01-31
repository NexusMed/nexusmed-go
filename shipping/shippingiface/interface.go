package paymentsiface

import (
	"github.com/nexusmed/nexusmed-go/shipping"
)

type ShippingAPI interface {
	// Create a new payment
	CreateShipment(shipping.CreateShipmentInput) (*shipping.CreateShipment, error)

	// Set the Authorization header
	SetApiKey(key string)
}
