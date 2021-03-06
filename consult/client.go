// Code generated by Nexus Automation, DO NOT EDIT.

package consult

import (
	"github.com/nexusmed/nexusmed-go/client"
)

type Client struct {
	*client.Client
}

func New(interceptors ...client.RequestInterceptor) *Client {
	return &Client{client.New("/consult/graphql", interceptors...)}
}

type Query struct {
	GetQuestionnaire        *Questionnaire        "json:\"getQuestionnaire\" graphql:\"getQuestionnaire\""
	GetQuestionnaires       *Questionnaires       "json:\"getQuestionnaires\" graphql:\"getQuestionnaires\""
	GetQuestionnaireAnswers *QuestionnaireAnswers "json:\"getQuestionnaireAnswers\" graphql:\"getQuestionnaireAnswers\""
	GetConsultation         Consultation          "json:\"getConsultation\" graphql:\"getConsultation\""
	GetConsultsCount        *int                  "json:\"getConsultsCount\" graphql:\"getConsultsCount\""
}
type Mutation struct {
	CreateQuestionnaire *Questionnaire        "json:\"createQuestionnaire\" graphql:\"createQuestionnaire\""
	AnswerQuestionnaire *QuestionnaireAnswers "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
	CreateConsultation  Consultation          "json:\"createConsultation\" graphql:\"createConsultation\""
}
type AsynchronousConsultationParts struct {
	ID      string "json:\"id\" graphql:\"id\""
	Patient struct {
		ID   string "json:\"id\" graphql:\"id\""
		Name *struct {
			GivenName  *string "json:\"given_name\" graphql:\"given_name\""
			FamilyName *string "json:\"family_name\" graphql:\"family_name\""
		} "json:\"name\" graphql:\"name\""
	} "json:\"patient\" graphql:\"patient\""
	Status   ConsultationStatus "json:\"status\" graphql:\"status\""
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
	QuestionnaireAnswers *struct {
		ID string "json:\"id\" graphql:\"id\""
	} "json:\"questionnaire_answers\" graphql:\"questionnaire_answers\""
}
type QuestionParts struct {
	Index   int            "json:\"index\" graphql:\"index\""
	Type    QuestionType   "json:\"type\" graphql:\"type\""
	Text    string         "json:\"text\" graphql:\"text\""
	Info    []*string      "json:\"info\" graphql:\"info\""
	Answers []*AnswerParts "json:\"answers\" graphql:\"answers\""
}
type AnswerParts struct {
	Index  int    "json:\"index\" graphql:\"index\""
	Value  string "json:\"value\" graphql:\"value\""
	Reject bool   "json:\"reject\" graphql:\"reject\""
}
type GetConsultation struct {
	GetConsultation *struct {
		ID      string "json:\"id\" graphql:\"id\""
		Patient struct {
			ID   string "json:\"id\" graphql:\"id\""
			Name *struct {
				GivenName  *string "json:\"given_name\" graphql:\"given_name\""
				FamilyName *string "json:\"family_name\" graphql:\"family_name\""
			} "json:\"name\" graphql:\"name\""
		} "json:\"patient\" graphql:\"patient\""
		Status   ConsultationStatus "json:\"status\" graphql:\"status\""
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
		QuestionnaireAnswers *struct {
			ID string "json:\"id\" graphql:\"id\""
		} "json:\"questionnaire_answers\" graphql:\"questionnaire_answers\""
	} "json:\"getConsultation\" graphql:\"getConsultation\""
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
				ID   string "json:\"id\" graphql:\"id\""
				Name *struct {
					GivenName  *string "json:\"given_name\" graphql:\"given_name\""
					FamilyName *string "json:\"family_name\" graphql:\"family_name\""
				} "json:\"name\" graphql:\"name\""
			} "json:\"patient\" graphql:\"patient\""
			Status   ConsultationStatus "json:\"status\" graphql:\"status\""
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
			QuestionnaireAnswers *struct {
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
		CreatedAt     *string "json:\"created_at\" graphql:\"created_at\""
		Questionnaire struct {
			ID string "json:\"id\" graphql:\"id\""
		} "json:\"questionnaire\" graphql:\"questionnaire\""
	} "json:\"answerQuestionnaire\" graphql:\"answerQuestionnaire\""
}

const GetConsultationDocument = `query GetConsultation ($id: ID!) {
	getConsultation(id: $id) {
		... on AsynchronousConsultation {
			... AsynchronousConsultationParts
		}
	}
}
fragment AsynchronousConsultationParts on AsynchronousConsultation {
	id
	patient {
		id
		name {
			given_name
			family_name
		}
	}
	status
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
	questionnaire_answers {
		id
	}
}
`

func (c *Client) GetConsultation(id string, interceptors ...client.RequestInterceptor) (*GetConsultation, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetConsultation
	if err := c.Client.Post("GetConsultation", GetConsultationDocument, &res, vars, interceptors...); err != nil {
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
fragment AnswerParts on Answer {
	index
	value
	reject
}
fragment QuestionParts on Question {
	index
	type
	text
	info
	answers {
		... AnswerParts
	}
}
`

func (c *Client) GetQuestionnaire(id string, interceptors ...client.RequestInterceptor) (*GetQuestionnaire, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetQuestionnaire
	if err := c.Client.Post("GetQuestionnaire", GetQuestionnaireDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const CreateConsultationDocument = `mutation CreateConsultation ($input: CreateConsultationInput!) {
	createConsultation(input: $input) {
		__typename
		... on AsynchronousConsultation {
			... AsynchronousConsultationParts
		}
	}
}
fragment AsynchronousConsultationParts on AsynchronousConsultation {
	id
	patient {
		id
		name {
			given_name
			family_name
		}
	}
	status
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
	questionnaire_answers {
		id
	}
}
`

func (c *Client) CreateConsultation(input CreateConsultationInput, interceptors ...client.RequestInterceptor) (*CreateConsultation, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateConsultation
	if err := c.Client.Post("CreateConsultation", CreateConsultationDocument, &res, vars, interceptors...); err != nil {
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
	info
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

func (c *Client) CreateQuestionnaire(input CreateQuestionnaireInput, interceptors ...client.RequestInterceptor) (*CreateQuestionnaire, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res CreateQuestionnaire
	if err := c.Client.Post("CreateQuestionnaire", CreateQuestionnaireDocument, &res, vars, interceptors...); err != nil {
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
		created_at
		questionnaire {
			id
		}
	}
}
fragment QuestionParts on Question {
	index
	type
	text
	info
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

func (c *Client) AnswerQuestionnaire(input AnswerQuestionnaireInput, interceptors ...client.RequestInterceptor) (*AnswerQuestionnaire, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res AnswerQuestionnaire
	if err := c.Client.Post("AnswerQuestionnaire", AnswerQuestionnaireDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
