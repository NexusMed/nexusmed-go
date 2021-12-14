// Code generated by Nexus Automation, DO NOT EDIT.

package generated

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
	CreatePatient       *Patient       "json:\"createPatient\" graphql:\"createPatient\""
	UpdatePatient       *Patient       "json:\"updatePatient\" graphql:\"updatePatient\""
	CreateQuestionnaire *Questionnaire "json:\"createQuestionnaire\" graphql:\"createQuestionnaire\""
	AnswerQuestionnaire *Answers       "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
	CreateConsultation  *Consultation  "json:\"createConsultation\" graphql:\"createConsultation\""
	CreatePayment       *Payment       "json:\"createPayment\" graphql:\"createPayment\""
}
type GetQuestionnaire struct {
	GetQuestionnaire *struct {
		ID        string  "json:\"id\" graphql:\"id\""
		Title     *string "json:\"title\" graphql:\"title\""
		Questions []*struct {
			Index    int          "json:\"index\" graphql:\"index\""
			Type     QuestionType "json:\"type\" graphql:\"type\""
			Text     string       "json:\"text\" graphql:\"text\""
			Required bool         "json:\"required\" graphql:\"required\""
			Info     *string      "json:\"info\" graphql:\"info\""
			Answers  []*struct {
				Index int    "json:\"index\" graphql:\"index\""
				Value string "json:\"value\" graphql:\"value\""
			} "json:\"answers\" graphql:\"answers\""
		} "json:\"questions\" graphql:\"questions\""
	} "json:\"getQuestionnaire\" graphql:\"getQuestionnaire\""
}
type CreatePatient struct {
	CreatePatient *struct {
		ID string "json:\"id\" graphql:\"id\""
	} "json:\"createPatient\" graphql:\"createPatient\""
}
type CreateQuestionnaire struct {
	CreateQuestionnaire *struct {
		ID string "json:\"id\" graphql:\"id\""
	} "json:\"createQuestionnaire\" graphql:\"createQuestionnaire\""
}
type AnswerQuestionnaire struct {
	AnswerQuestionnaire *struct {
		ID string "json:\"id\" graphql:\"id\""
	} "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
}

const GetQuestionnaireDocument = `query GetQuestionnaire ($id: ID!) {
	getQuestionnaire(id: $id) {
		id
		title
		questions {
			index
			type
			text
			required
			info
			answers {
				index
				value
			}
		}
	}
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
	}
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
	}
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
