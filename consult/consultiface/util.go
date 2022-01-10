package consultiface

import (
	"encoding/json"

	"github.com/nexusmed/nexusmed-go/consult"
)

// AnswerQuestionnaire converts consult.QuestionnaireAnswers to *consult.AnswerQuestionnaire
func AnswerQuestionnaire(in consult.QuestionnaireAnswers) *consult.AnswerQuestionnaire {
	b, _ := json.Marshal(in)
	var out consult.AnswerQuestionnaire
	json.Unmarshal(b, &out)
	return &out
}
