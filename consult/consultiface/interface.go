package consultiface

import (
	"github.com/nexusmed/nexusmed-go/client"
	"github.com/nexusmed/nexusmed-go/consult"
)

type ConsultAPI interface {
	AnswerQuestionnaire(consult.AnswerQuestionnaireInput) (*consult.AnswerQuestionnaire, error)
	
	New() *Client
	SetApiKey(key string)
}
