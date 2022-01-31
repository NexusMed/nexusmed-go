package paymentsiface

import (
	"encoding/json"

	"github.com/nexusmed/nexusmed-go/shipping"
)

// CreateShipment converts shipping.Shipment to *payments.CreateShipment
func CreateShipment(in shipping.Shipment) *shipping.CreateShipment {
	b, _ := json.Marshal(in)
	var out shipping.CreateShipment
	json.Unmarshal(b, &out.CreateShipment)
	return &out
}
