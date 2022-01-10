package consultiface

import (
	"github.com/nexusmed/nexusmed-go/consult"
)

type ConsultAPI interface {
	AnswerQuestionnaire(consult.AnswerQuestionnaireInput) (*consult.AnswerQuestionnaire, error)

	New() *consult.Client
	SetApiKey(key string)
}
