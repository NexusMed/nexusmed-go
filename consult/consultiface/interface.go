package consultiface

import (
	"github.com/nexusmed/nexusmed-go/client"
	"github.com/nexusmed/nexusmed-go/consult"
)

type ConsultAPI interface {
	// Get a questionnaire by Id
	GetQuestionnaire(id string, interceptors ...client.RequestInterceptor) (*consult.GetQuestionnaire, error)

	// Create a new consultation
	CreateConsultation(consult.CreateConsultationInput, ...client.RequestInterceptor) (*consult.CreateConsultation, error)

	// Create a new questionnaire
	CreateQuestionnaire(consult.CreateQuestionnaireInput, ...client.RequestInterceptor) (*consult.CreateQuestionnaire, error)

	// Answer an existing questionnaire
	AnswerQuestionnaire(consult.AnswerQuestionnaireInput, ...client.RequestInterceptor) (*consult.AnswerQuestionnaire, error)

	// Set the Authorization header
	SetApiKey(key string)
}
