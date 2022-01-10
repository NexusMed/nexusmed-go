package nexusmed

import (
	"encoding/json"

	"github.com/nexusmed/nexusmed-go/consult"
)

// AnswerInputList converts []consult.AnswerInput to []*consult.AnswerInput
func AnswerInputList(in []consult.AnswerInput) []*consult.AnswerInput {
	out := make([]*consult.AnswerInput, len(in))
	for i := range in {
		out[i] = &in[i]
	}
	return out
}

// QuestionnaireAnswers converts interface{} to *consult.QuestionnaireAnswers
func QuestionnaireAnswers(in interface{}) *consult.QuestionnaireAnswers {
	b, _ := json.Marshal(in)
	var qans consult.QuestionnaireAnswers
	json.Unmarshal(b, &qans)
	return &qans
}
