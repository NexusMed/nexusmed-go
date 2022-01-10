package consultiface

import (
	"github.com/nexusmed/nexusmed-go/consult"
)

type ConsultAPI interface {
	// Get a questionnaire by Id
	GetQuestionnaire(id string) (*consult.GetQuestionnaire, error)

	// Create a new consultation
	CreateConsultation(consult.CreateConsultationInput) (*consult.CreateConsultation, error)

	// Create a new questionnaire
	CreateQuestionnaire(consult.CreateQuestionnaireInput) (*consult.CreateQuestionnaire, error)

	// Answer an existing questionnaire
	AnswerQuestionnaire(consult.AnswerQuestionnaireInput) (*consult.AnswerQuestionnaire, error)

	// Set the Authorization header
	SetApiKey(key string)
}
