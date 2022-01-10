package consultiface

import (
	"encoding/json"

	"github.com/nexusmed/nexusmed-go/consult"
)

// AnswerQuestionnaire converts consult.QuestionnaireAnswers to *consult.AnswerQuestionnaire
func AnswerQuestionnaire(in consult.QuestionnaireAnswers) *consult.AnswerQuestionnaire {
	b, _ := json.Marshal(in)
	var out consult.AnswerQuestionnaire
	json.Unmarshal(b, &out.AnswerQuestionnaire)
	return &out
}

// AnswerInputListValues converts []*consult.AnswerInput to []consult.AnswerInput
func AnswerInputListValues(in []*consult.AnswerInput) []consult.AnswerInput {
	out := make([]consult.AnswerInput, len(in))
	for i := range in {
		out[i] = *in[i]
	}
	return out
}
