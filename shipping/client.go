// Code generated by Nexus Automation, DO NOT EDIT.

package shipping

import (
	"github.com/nexusmed/nexusmed-go/client"
)

type Client struct {
	*client.Client
}

func New(interceptors ...client.RequestInterceptor) *Client {
	return &Client{client.New("/shipping/graphql", interceptors...)}
}

type Query struct {
	GetShipment *Shipment "json:\"getShipment\" graphql:\"getShipment\""
}
type Mutation struct {
	CreateShipment *Shipment "json:\"createShipment\" graphql:\"createShipment\""
}
type ProductParts struct {
	ID       *string "json:\"id\" graphql:\"id\""
	Name     *string "json:\"name\" graphql:\"name\""
	Quantity *int    "json:\"quantity\" graphql:\"quantity\""
}
type PatientParts struct {
	ID   *string "json:\"id\" graphql:\"id\""
	Name *struct {
		Title      *string "json:\"title\" graphql:\"title\""
		GivenName  *string "json:\"given_name\" graphql:\"given_name\""
		FamilyName *string "json:\"family_name\" graphql:\"family_name\""
	} "json:\"name\" graphql:\"name\""
}
type AddressParts struct {
	Line1      *string "json:\"line1\" graphql:\"line1\""
	Line2      *string "json:\"line2\" graphql:\"line2\""
	City       *string "json:\"city\" graphql:\"city\""
	PostalCode *string "json:\"postal_code\" graphql:\"postal_code\""
}
type SenderParts struct {
	Pharmacy struct {
		ID      string        "json:\"id\" graphql:\"id\""
		Address *AddressParts "json:\"address\" graphql:\"address\""
	} "graphql:\"... on Pharmacy\""
}
type GetShipment struct {
	GetShipment *struct {
		ID        string          "json:\"id\" graphql:\"id\""
		Products  []*ProductParts "json:\"products\" graphql:\"products\""
		Patient   *PatientParts   "json:\"patient\" graphql:\"patient\""
		Address   *AddressParts   "json:\"address\" graphql:\"address\""
		Courier   *Courier        "json:\"courier\" graphql:\"courier\""
		Service   *Service        "json:\"service\" graphql:\"service\""
		Status    *ShipmentStatus "json:\"status\" graphql:\"status\""
		Sender    *SenderParts    "json:\"sender\" graphql:\"sender\""
		CreatedAt *string         "json:\"created_at\" graphql:\"created_at\""
	} "json:\"getShipment\" graphql:\"getShipment\""
}
type CreateShipment struct {
	CreateShipment *struct {
		ID        string          "json:\"id\" graphql:\"id\""
		Products  []*ProductParts "json:\"products\" graphql:\"products\""
		Patient   *PatientParts   "json:\"patient\" graphql:\"patient\""
		Address   *AddressParts   "json:\"address\" graphql:\"address\""
		Courier   *Courier        "json:\"courier\" graphql:\"courier\""
		Service   *Service        "json:\"service\" graphql:\"service\""
		Status    *ShipmentStatus "json:\"status\" graphql:\"status\""
		Sender    *SenderParts    "json:\"sender\" graphql:\"sender\""
		CreatedAt *string         "json:\"created_at\" graphql:\"created_at\""
	} "json:\"createShipment\" graphql:\"createShipment\""
}

const GetShipmentDocument = `query GetShipment ($id: ID!) {
	getShipment(id: $id) {
		id
		products {
			... ProductParts
		}
		patient {
			... PatientParts
		}
		address {
			... AddressParts
		}
		courier
		service
		status
		sender {
			... SenderParts
		}
		created_at
	}
}
fragment ProductParts on Product {
	id
	name
	quantity
}
fragment PatientParts on Patient {
	id
	name {
		title
		given_name
		family_name
	}
}
fragment AddressParts on Address {
	line1
	line2
	city
	postal_code
}
fragment SenderParts on Sender {
	... on Pharmacy {
		id
		address {
			... AddressParts
		}
	}
}
`

func (c *Client) GetShipment(id string, interceptors ...client.RequestInterceptor) (*GetShipment, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetShipment
	if err := c.Client.Post("GetShipment", GetShipmentDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateShipmentDocument = `mutation CreateShipment ($input: CreateShipmentInput!) {
	createShipment(input: $input) {
		id
		products {
			... ProductParts
		}
		patient {
			... PatientParts
		}
		address {
			... AddressParts
		}
		courier
		service
		status
		sender {
			... SenderParts
		}
		created_at
	}
}
fragment SenderParts on Sender {
	... on Pharmacy {
		id
		address {
			... AddressParts
		}
	}
}
fragment ProductParts on Product {
	id
	name
	quantity
}
fragment PatientParts on Patient {
	id
	name {
		title
		given_name
		family_name
	}
}
fragment AddressParts on Address {
	line1
	line2
	city
	postal_code
}
`

func (c *Client) CreateShipment(input CreateShipmentInput, interceptors ...client.RequestInterceptor) (*CreateShipment, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateShipment
	if err := c.Client.Post("CreateShipment", CreateShipmentDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
