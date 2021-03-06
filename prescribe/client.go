// Code generated by Nexus Automation, DO NOT EDIT.

package prescribe

import (
	"github.com/nexusmed/nexusmed-go/client"
)

type Client struct {
	*client.Client
}

func New(interceptors ...client.RequestInterceptor) *Client {
	return &Client{client.New("/prescribe/graphql", interceptors...)}
}

type Query struct {
	GetPrescription       *Prescription "json:\"getPrescription\" graphql:\"getPrescription\""
	GetPrescriptionsCount *int          "json:\"getPrescriptionsCount\" graphql:\"getPrescriptionsCount\""
}
type Mutation struct {
	CreatePrescription *Prescription "json:\"createPrescription\" graphql:\"createPrescription\""
}
type GetPrescription struct {
	GetPrescription *struct {
		ID      string "json:\"id\" graphql:\"id\""
		Patient *struct {
			ID   string "json:\"id\" graphql:\"id\""
			Name struct {
				Title      *string "json:\"title\" graphql:\"title\""
				GivenName  *string "json:\"given_name\" graphql:\"given_name\""
				FamilyName *string "json:\"family_name\" graphql:\"family_name\""
			} "json:\"name\" graphql:\"name\""
		} "json:\"patient\" graphql:\"patient\""
		Prescriber *struct {
			ID   string "json:\"id\" graphql:\"id\""
			Name *struct {
				Title      *string "json:\"title\" graphql:\"title\""
				GivenName  *string "json:\"given_name\" graphql:\"given_name\""
				FamilyName *string "json:\"family_name\" graphql:\"family_name\""
			} "json:\"name\" graphql:\"name\""
			Register *struct {
				Type  *RegisterType "json:\"type\" graphql:\"type\""
				Value *string       "json:\"value\" graphql:\"value\""
			} "json:\"register\" graphql:\"register\""
		} "json:\"prescriber\" graphql:\"prescriber\""
		Pharmacy *struct {
			ID       string "json:\"id\" graphql:\"id\""
			Register *struct {
				Type  *RegisterType "json:\"type\" graphql:\"type\""
				Value *string       "json:\"value\" graphql:\"value\""
			} "json:\"register\" graphql:\"register\""
			Business *struct {
				Name    *string "json:\"name\" graphql:\"name\""
				Address *struct {
					Line1      *string "json:\"line1\" graphql:\"line1\""
					Line2      *string "json:\"line2\" graphql:\"line2\""
					City       *string "json:\"city\" graphql:\"city\""
					PostalCode *string "json:\"postal_code\" graphql:\"postal_code\""
				} "json:\"address\" graphql:\"address\""
			} "json:\"business\" graphql:\"business\""
		} "json:\"pharmacy\" graphql:\"pharmacy\""
		Products []*struct {
			ID         string "json:\"id\" graphql:\"id\""
			Name       string "json:\"name\" graphql:\"name\""
			Medication *struct {
				Name   *string "json:\"name\" graphql:\"name\""
				Dosage *struct {
					Quantity *float64    "json:\"quantity\" graphql:\"quantity\""
					Unit     *DosageUnit "json:\"unit\" graphql:\"unit\""
				} "json:\"dosage\" graphql:\"dosage\""
				Quantity *int "json:\"quantity\" graphql:\"quantity\""
			} "json:\"medication\" graphql:\"medication\""
		} "json:\"products\" graphql:\"products\""
		CreatedAt *string             "json:\"created_at\" graphql:\"created_at\""
		Status    *PrescriptionStatus "json:\"status\" graphql:\"status\""
	} "json:\"getPrescription\" graphql:\"getPrescription\""
}

const GetPrescriptionDocument = `query GetPrescription ($id: ID!) {
	getPrescription(id: $id) {
		id
		patient {
			id
			name {
				title
				given_name
				family_name
			}
		}
		prescriber {
			id
			name {
				title
				given_name
				family_name
			}
			register {
				type
				value
			}
		}
		pharmacy {
			id
			register {
				type
				value
			}
			business {
				name
				address {
					line1
					line2
					city
					postal_code
				}
			}
		}
		products {
			id
			name
			medication {
				name
				dosage {
					quantity
					unit
				}
				quantity
			}
		}
		created_at
		status
	}
}
`

func (c *Client) GetPrescription(id string, interceptors ...client.RequestInterceptor) (*GetPrescription, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetPrescription
	if err := c.Client.Post("GetPrescription", GetPrescriptionDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
