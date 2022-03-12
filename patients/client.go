// Code generated by Nexus Automation, DO NOT EDIT.

package patients

import (
	"github.com/nexusmed/nexusmed-go/client"
)

type Client struct {
	*client.Client
}

func New(interceptors ...client.RequestInterceptor) *Client {
	return &Client{client.New("/patients/graphql", interceptors...)}
}

type Query struct {
	GetPatient  *Patient  "json:\"getPatient\" graphql:\"getPatient\""
	GetPatients *Patients "json:\"getPatients\" graphql:\"getPatients\""
}
type Mutation struct {
	CreatePatient *Patient "json:\"createPatient\" graphql:\"createPatient\""
}
type PatientParts struct {
	ID          string    "json:\"id\" graphql:\"id\""
	Name        NameParts "json:\"name\" graphql:\"name\""
	Email       *string   "json:\"email\" graphql:\"email\""
	Phone       *string   "json:\"phone\" graphql:\"phone\""
	DateOfBirth *struct {
		Day   *int "json:\"day\" graphql:\"day\""
		Month *int "json:\"month\" graphql:\"month\""
		Year  *int "json:\"year\" graphql:\"year\""
	} "json:\"date_of_birth\" graphql:\"date_of_birth\""
	Sex     *SexAssignment  "json:\"sex\" graphql:\"sex\""
	Gender  *GenderIdentity "json:\"gender\" graphql:\"gender\""
	Address *struct {
		Line1      *string "json:\"line1\" graphql:\"line1\""
		Line2      *string "json:\"line2\" graphql:\"line2\""
		City       *string "json:\"city\" graphql:\"city\""
		PostalCode *string "json:\"postal_code\" graphql:\"postal_code\""
		Country    *string "json:\"country\" graphql:\"country\""
	} "json:\"address\" graphql:\"address\""
	Stripe *struct {
		ID *string "json:\"id\" graphql:\"id\""
	} "json:\"stripe\" graphql:\"stripe\""
}
type NameParts struct {
	Title      *string "json:\"title\" graphql:\"title\""
	GivenName  string  "json:\"given_name\" graphql:\"given_name\""
	FamilyName string  "json:\"family_name\" graphql:\"family_name\""
}
type PrescriptionParts struct {
	ID         string "json:\"id\" graphql:\"id\""
	CreatedAt  string "json:\"created_at\" graphql:\"created_at\""
	Prescriber struct {
		ID   string    "json:\"id\" graphql:\"id\""
		Name NameParts "json:\"name\" graphql:\"name\""
	} "json:\"prescriber\" graphql:\"prescriber\""
}
type GetPatient struct {
	GetPatient *struct {
		ID          string    "json:\"id\" graphql:\"id\""
		Name        NameParts "json:\"name\" graphql:\"name\""
		Email       *string   "json:\"email\" graphql:\"email\""
		Phone       *string   "json:\"phone\" graphql:\"phone\""
		DateOfBirth *struct {
			Day   *int "json:\"day\" graphql:\"day\""
			Month *int "json:\"month\" graphql:\"month\""
			Year  *int "json:\"year\" graphql:\"year\""
		} "json:\"date_of_birth\" graphql:\"date_of_birth\""
		Sex     *SexAssignment  "json:\"sex\" graphql:\"sex\""
		Gender  *GenderIdentity "json:\"gender\" graphql:\"gender\""
		Address *struct {
			Line1      *string "json:\"line1\" graphql:\"line1\""
			Line2      *string "json:\"line2\" graphql:\"line2\""
			City       *string "json:\"city\" graphql:\"city\""
			PostalCode *string "json:\"postal_code\" graphql:\"postal_code\""
			Country    *string "json:\"country\" graphql:\"country\""
		} "json:\"address\" graphql:\"address\""
		Stripe *struct {
			ID *string "json:\"id\" graphql:\"id\""
		} "json:\"stripe\" graphql:\"stripe\""
		Consultations []*struct {
			ID        string "json:\"id\" graphql:\"id\""
			CreatedAt string "json:\"created_at\" graphql:\"created_at\""
			Assignee  *struct {
				ID   string    "json:\"id\" graphql:\"id\""
				Name NameParts "json:\"name\" graphql:\"name\""
			} "json:\"assignee\" graphql:\"assignee\""
			Prescriptions []*PrescriptionParts "json:\"prescriptions\" graphql:\"prescriptions\""
		} "json:\"consultations\" graphql:\"consultations\""
		Prescriptions []*PrescriptionParts "json:\"prescriptions\" graphql:\"prescriptions\""
	} "json:\"getPatient\" graphql:\"getPatient\""
}
type GetPatients struct {
	GetPatients *struct {
		Items     []*PatientParts "json:\"items\" graphql:\"items\""
		NextToken *string         "json:\"next_token\" graphql:\"next_token\""
	} "json:\"getPatients\" graphql:\"getPatients\""
}
type CreatePatient struct {
	CreatePatient *PatientParts "json:\"createPatient\" graphql:\"createPatient\""
}

const GetPatientDocument = `query GetPatient ($id: ID!) {
	getPatient(id: $id) {
		... PatientParts
		consultations {
			id
			created_at
			assignee {
				id
				name {
					... NameParts
				}
			}
			prescriptions {
				... PrescriptionParts
			}
		}
		prescriptions {
			... PrescriptionParts
		}
	}
}
fragment PatientParts on Patient {
	id
	name {
		... NameParts
	}
	email
	phone
	date_of_birth {
		day
		month
		year
	}
	sex
	gender
	address {
		line1
		line2
		city
		postal_code
		country
	}
	stripe {
		id
	}
}
fragment NameParts on Name {
	title
	given_name
	family_name
}
fragment PrescriptionParts on Prescription {
	id
	created_at
	prescriber {
		id
		name {
			... NameParts
		}
	}
}
`

func (c *Client) GetPatient(id string, interceptors ...client.RequestInterceptor) (*GetPatient, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetPatient
	if err := c.Client.Post("GetPatient", GetPatientDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetPatientsDocument = `query GetPatients ($limit: Int, $next_token: String) {
	getPatients(limit: $limit, next_token: $next_token) {
		items {
			... PatientParts
		}
		next_token
	}
}
fragment NameParts on Name {
	title
	given_name
	family_name
}
fragment PatientParts on Patient {
	id
	name {
		... NameParts
	}
	email
	phone
	date_of_birth {
		day
		month
		year
	}
	sex
	gender
	address {
		line1
		line2
		city
		postal_code
		country
	}
	stripe {
		id
	}
}
`

func (c *Client) GetPatients(limit *int, nextToken *string, interceptors ...client.RequestInterceptor) (*GetPatients, error) {
	vars := map[string]interface{}{
		"limit":      limit,
		"next_token": nextToken,
	}

	var res GetPatients
	if err := c.Client.Post("GetPatients", GetPatientsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreatePatientDocument = `mutation CreatePatient ($input: CreatePatientInput!) {
	createPatient(input: $input) {
		... PatientParts
	}
}
fragment NameParts on Name {
	title
	given_name
	family_name
}
fragment PatientParts on Patient {
	id
	name {
		... NameParts
	}
	email
	phone
	date_of_birth {
		day
		month
		year
	}
	sex
	gender
	address {
		line1
		line2
		city
		postal_code
		country
	}
	stripe {
		id
	}
}
`

func (c *Client) CreatePatient(input CreatePatientInput, interceptors ...client.RequestInterceptor) (*CreatePatient, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreatePatient
	if err := c.Client.Post("CreatePatient", CreatePatientDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
