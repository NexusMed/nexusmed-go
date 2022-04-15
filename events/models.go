// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package events

import (
	"fmt"
	"io"
	"strconv"
)

type CreateEndpointInput struct {
	URL         string     `json:"url"`
	Resources   []Resource `json:"resources,omitempty"`
	Description *string    `json:"description,omitempty"`
}

type Data struct {
	Object map[string]interface{} `json:"object,omitempty"`
}

type Endpoint struct {
	ID          *string     `json:"id,omitempty"`
	URL         *string     `json:"url,omitempty"`
	Resources   []*Resource `json:"resources,omitempty"`
	Description *string     `json:"description,omitempty"`
}

type Event struct {
	ID        string  `json:"id"`
	Type      *string `json:"type,omitempty"`
	AccountID *string `json:"account_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Data      *Data   `json:"data,omitempty"`
	Delivered *bool   `json:"delivered,omitempty"`
}

type Events struct {
	Items     []*Event `json:"items,omitempty"`
	NextToken *string  `json:"nextToken,omitempty"`
}

type UpdateEndpointInput struct {
	ID          string     `json:"id"`
	URL         string     `json:"url"`
	Resources   []Resource `json:"resources,omitempty"`
	Description *string    `json:"description,omitempty"`
}

type Resource string

const (
	ResourceShipment      Resource = "shipment"
	ResourceConsultation  Resource = "consultation"
	ResourcePrescription  Resource = "prescription"
	ResourcePatient       Resource = "patient"
	ResourceQuestionnaire Resource = "questionnaire"
	ResourcePayment       Resource = "payment"
	ResourceAccount       Resource = "account"
	ResourceProduct       Resource = "product"
)

var AllResource = []Resource{
	ResourceShipment,
	ResourceConsultation,
	ResourcePrescription,
	ResourcePatient,
	ResourceQuestionnaire,
	ResourcePayment,
	ResourceAccount,
	ResourceProduct,
}

func (e Resource) IsValid() bool {
	switch e {
	case ResourceShipment, ResourceConsultation, ResourcePrescription, ResourcePatient, ResourceQuestionnaire, ResourcePayment, ResourceAccount, ResourceProduct:
		return true
	}
	return false
}

func (e Resource) String() string {
	return string(e)
}

func (e *Resource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Resource(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Resource", str)
	}
	return nil
}

func (e Resource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
