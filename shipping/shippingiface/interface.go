package shippingiface

import (
	"github.com/nexusmed/nexusmed-go/client"
	"github.com/nexusmed/nexusmed-go/shipping"
)

type ShippingAPI interface {
	// Get a shipment
	GetShipment(id string, interceptors ...client.RequestInterceptor) (*shipping.GetShipment, error)

	// Create a new shipment
	CreateShipment(shipping.CreateShipmentInput, ...client.RequestInterceptor) (*shipping.CreateShipment, error)

	// Set the Authorization header
	SetApiKey(key string)
}
