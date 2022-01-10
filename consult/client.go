// Code generated by Nexus Automation, DO NOT EDIT.

package consult

import (
	"net/http"

	"github.com/nexusmed/nexusmed-go/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client) *Client {
	return &Client{Client: client.NewClient(cli, "https://api.nexusmed.io/consult")}
}

func (c *Client) SetApiKey(key string) {
	c.Client.SetApiKey(key)
}

type Query struct {
	GetQuestionnaire        *Questionnaire        "json:\"getQuestionnaire\" graphql:\"getQuestionnaire\""
	GetQuestionnaires       *Questionnaires       "json:\"getQuestionnaires\" graphql:\"getQuestionnaires\""
	GetQuestionnaireAnswers *QuestionnaireAnswers "json:\"getQuestionnaireAnswers\" graphql:\"getQuestionnaireAnswers\""
	GetConsultation         Consultation          "json:\"getConsultation\" graphql:\"getConsultation\""
}
type Mutation struct {
	CreateQuestionnaire *Questionnaire        "json:\"createQuestionnaire\" graphql:\"createQuestionnaire\""
	AnswerQuestionnaire *QuestionnaireAnswers "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
	CreateConsultation  Consultation          "json:\"createConsultation\" graphql:\"createConsultation\""
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
type GetQuestionnaire struct {
	GetQuestionnaire *struct {
		ID        string           "json:\"id\" graphql:\"id\""
		Title     *string          "json:\"title\" graphql:\"title\""
		Questions []*QuestionParts "json:\"questions\" graphql:\"questions\""
	} "json:\"getQuestionnaire\" graphql:\"getQuestionnaire\""
}
type CreateConsultation struct {
	CreateConsultation *struct {
		Typename                 *string "json:\"__typename\" graphql:\"__typename\""
		AsynchronousConsultation struct {
			ID      string "json:\"id\" graphql:\"id\""
			Patient struct {
				ID string "json:\"id\" graphql:\"id\""
			} "json:\"patient\" graphql:\"patient\""
			Status   ConsultationStatus "json:\"status\" graphql:\"status\""
			Products []*struct {
				ID string "json:\"id\" graphql:\"id\""
			} "json:\"products\" graphql:\"products\""
			QuestionnaireAnswers struct {
				ID string "json:\"id\" graphql:\"id\""
			} "json:\"questionnaire_answers\" graphql:\"questionnaire_answers\""
		} "graphql:\"... on AsynchronousConsultation\""
	} "json:\"createConsultation\" graphql:\"createConsultation\""
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

func (c *Client) GetQuestionnaire(id string) (*GetQuestionnaire, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetQuestionnaire
	if err := c.Client.Post("GetQuestionnaire", GetQuestionnaireDocument, &res, vars); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateConsultationDocument = `mutation CreateConsultation ($input: CreateConsultationInput!) {
	createConsultation(input: $input) {
		__typename
		... on AsynchronousConsultation {
			id
			patient {
				id
			}
			status
			products {
				id
			}
			questionnaire_answers {
				id
			}
		}
	}
}
`

func (c *Client) CreateConsultation(input CreateConsultationInput) (*CreateConsultation, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateConsultation
	if err := c.Client.Post("CreateConsultation", CreateConsultationDocument, &res, vars); err != nil {
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
fragment AnswerParts on Answer {
	index
	value
	reject
}
fragment QuestionParts on Question {
	index
	type
	text
	answers {
		... AnswerParts
	}
}
`

func (c *Client) CreateQuestionnaire(input CreateQuestionnaireInput) (*CreateQuestionnaire, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateQuestionnaire
	if err := c.Client.Post("CreateQuestionnaire", CreateQuestionnaireDocument, &res, vars); err != nil {
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

func (c *Client) AnswerQuestionnaire(input AnswerQuestionnaireInput) (*AnswerQuestionnaire, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res AnswerQuestionnaire
	if err := c.Client.Post("AnswerQuestionnaire", AnswerQuestionnaireDocument, &res, vars); err != nil {
		return nil, err
	}

	return &res, nil
}
