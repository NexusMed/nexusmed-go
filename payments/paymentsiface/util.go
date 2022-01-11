package paymentsiface

import (
	"encoding/json"

	"github.com/nexusmed/nexusmed-go/payments"
)

// CreatePayment converts payment.Payment to *payments.CreatePayment
func CreatePayment(in payments.Payment) *payments.CreatePayment {
	b, _ := json.Marshal(in)
	var out payments.CreatePayment
	json.Unmarshal(b, &out.CreatePayment)
	return &out
}
