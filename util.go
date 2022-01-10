package nexusmed

import "github.com/nexusmed/nexusmed-go/consult"

// Converts []consult.AnswerInput to []*consult.AnswerInput
func AnswersInput(in []consult.AnswerInput) []*consult.AnswerInput {
	out := make([]*consult.AnswerInput, len(in))
	for i := range in {
		out[i] = &in[i]
	}
	return out
}
