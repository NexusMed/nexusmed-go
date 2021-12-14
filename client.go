// Code generated by Nexus Automation, DO NOT EDIT.

package nexusmed

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) *Client {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	GetPatient       *Patient       "json:\"getPatient\" graphql:\"getPatient\""
	GetQuestionnaire *Questionnaire "json:\"getQuestionnaire\" graphql:\"getQuestionnaire\""
	GetConsultation  *Consultation  "json:\"getConsultation\" graphql:\"getConsultation\""
	GetProduct       *Product       "json:\"getProduct\" graphql:\"getProduct\""
	GetProducts      *Products      "json:\"getProducts\" graphql:\"getProducts\""
}
type Mutation struct {
	CreatePatient       *Patient              "json:\"createPatient\" graphql:\"createPatient\""
	UpdatePatient       *Patient              "json:\"updatePatient\" graphql:\"updatePatient\""
	CreateQuestionnaire *Questionnaire        "json:\"createQuestionnaire\" graphql:\"createQuestionnaire\""
	AnswerQuestionnaire *QuestionnaireAnswers "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
	CreateConsultation  *Consultation         "json:\"createConsultation\" graphql:\"createConsultation\""
	CreatePayment       *Payment              "json:\"createPayment\" graphql:\"createPayment\""
}
type QuestionParts struct {
	Index   int            "json:\"index\" graphql:\"index\""
	Type    QuestionType   "json:\"type\" graphql:\"type\""
	Text    string         "json:\"text\" graphql:\"text\""
	Answers []*AnswerParts "json:\"answers\" graphql:\"answers\""
}
type AnswerParts struct {
	Index  int    "json:\"index\" graphql:\"index\""
	Value  string "json:\"value\" graphql:\"value\""
	Reject bool   "json:\"reject\" graphql:\"reject\""
}
type GetPatient struct {
	GetPatient *struct {
		ID   string "json:\"id\" graphql:\"id\""
		Name struct {
			Title      *string "json:\"title\" graphql:\"title\""
			GivenName  string  "json:\"given_name\" graphql:\"given_name\""
			FamilyName string  "json:\"family_name\" graphql:\"family_name\""
		} "json:\"name\" graphql:\"name\""
		Phone   *string "json:\"phone\" graphql:\"phone\""
		Email   *string "json:\"email\" graphql:\"email\""
		Address *struct {
			Line1      string  "json:\"line1\" graphql:\"line1\""
			Line2      *string "json:\"line2\" graphql:\"line2\""
			City       *string "json:\"city\" graphql:\"city\""
			PostalCode string  "json:\"postal_code\" graphql:\"postal_code\""
		} "json:\"address\" graphql:\"address\""
	} "json:\"getPatient\" graphql:\"getPatient\""
}
type GetQuestionnaire struct {
	GetQuestionnaire *struct {
		ID        string           "json:\"id\" graphql:\"id\""
		Title     *string          "json:\"title\" graphql:\"title\""
		Questions []*QuestionParts "json:\"questions\" graphql:\"questions\""
	} "json:\"getQuestionnaire\" graphql:\"getQuestionnaire\""
}
type CreateConsultation struct {
	CreateConsultation *struct {
		ID string "json:\"id\" graphql:\"id\""
	} "json:\"createConsultation\" graphql:\"createConsultation\""
}
type CreatePatient struct {
	CreatePatient *struct {
		ID string "json:\"id\" graphql:\"id\""
	} "json:\"createPatient\" graphql:\"createPatient\""
}
type CreateQuestionnaire struct {
	CreateQuestionnaire *struct {
		ID        string           "json:\"id\" graphql:\"id\""
		Title     *string          "json:\"title\" graphql:\"title\""
		Questions []*QuestionParts "json:\"questions\" graphql:\"questions\""
	} "json:\"createQuestionnaire\" graphql:\"createQuestionnaire\""
}
type AnswerQuestionnaire struct {
	AnswerQuestionnaire *struct {
		ID      string "json:\"id\" graphql:\"id\""
		Answers []*struct {
			Question QuestionParts "json:\"question\" graphql:\"question\""
			Value    []string      "json:\"value\" graphql:\"value\""
		} "json:\"answers\" graphql:\"answers\""
	} "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
}

const GetPatientDocument = `query GetPatient ($id: ID!) {
	getPatient(id: $id) {
		id
		name {
			title
			given_name
			family_name
		}
		phone
		email
		address {
			line1
			line2
			city
			postal_code
		}
	}
}
`

func (c *Client) GetPatient(ctx context.Context, id string, httpRequestOptions ...client.HTTPRequestOption) (*GetPatient, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetPatient
	if err := c.Client.Post(ctx, "GetPatient", GetPatientDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetQuestionnaireDocument = `query GetQuestionnaire ($id: ID!) {
	getQuestionnaire(id: $id) {
		id
		title
		questions {
			... QuestionParts
		}
	}
}
fragment QuestionParts on Question {
	index
	type
	text
	answers {
		... AnswerParts
	}
}
fragment AnswerParts on Answer {
	index
	value
	reject
}
`

func (c *Client) GetQuestionnaire(ctx context.Context, id string, httpRequestOptions ...client.HTTPRequestOption) (*GetQuestionnaire, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetQuestionnaire
	if err := c.Client.Post(ctx, "GetQuestionnaire", GetQuestionnaireDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateConsultationDocument = `mutation CreateConsultation ($input: CreateConsultationInput!) {
	createConsultation(input: $input) {
		id
	}
}
`

func (c *Client) CreateConsultation(ctx context.Context, input CreateConsultationInput, httpRequestOptions ...client.HTTPRequestOption) (*CreateConsultation, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateConsultation
	if err := c.Client.Post(ctx, "CreateConsultation", CreateConsultationDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreatePatientDocument = `mutation CreatePatient ($input: CreatePatientInput!) {
	createPatient(input: $input) {
		id
	}
}
`

func (c *Client) CreatePatient(ctx context.Context, input CreatePatientInput, httpRequestOptions ...client.HTTPRequestOption) (*CreatePatient, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreatePatient
	if err := c.Client.Post(ctx, "CreatePatient", CreatePatientDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateQuestionnaireDocument = `mutation CreateQuestionnaire ($input: CreateQuestionnaireInput!) {
	createQuestionnaire(input: $input) {
		id
		title
		questions {
			... QuestionParts
		}
	}
}
fragment QuestionParts on Question {
	index
	type
	text
	answers {
		... AnswerParts
	}
}
fragment AnswerParts on Answer {
	index
	value
	reject
}
`

func (c *Client) CreateQuestionnaire(ctx context.Context, input CreateQuestionnaireInput, httpRequestOptions ...client.HTTPRequestOption) (*CreateQuestionnaire, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateQuestionnaire
	if err := c.Client.Post(ctx, "CreateQuestionnaire", CreateQuestionnaireDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const AnswerQuestionnaireDocument = `mutation AnswerQuestionnaire ($input: AnswerQuestionnaireInput!) {
	answerQuestionnaire(input: $input) {
		id
		answers {
			question {
				... QuestionParts
			}
			value
		}
	}
}
fragment QuestionParts on Question {
	index
	type
	text
	answers {
		... AnswerParts
	}
}
fragment AnswerParts on Answer {
	index
	value
	reject
}
`

func (c *Client) AnswerQuestionnaire(ctx context.Context, input AnswerQuestionnaireInput, httpRequestOptions ...client.HTTPRequestOption) (*AnswerQuestionnaire, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res AnswerQuestionnaire
	if err := c.Client.Post(ctx, "AnswerQuestionnaire", AnswerQuestionnaireDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}
