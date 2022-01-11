package nexusmed

import (
	"encoding/json"

	"github.com/nexusmed/nexusmed-go/consult"
	"github.com/nexusmed/nexusmed-go/patients"
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

// AsynchronousConsultation converts interface{} to *consult.AsynchronousConsultation
func AsynchronousConsultation(in interface{}) *consult.AsynchronousConsultation {
	b, _ := json.Marshal(in)
	var aclt consult.AsynchronousConsultation
	json.Unmarshal(b, &aclt)
	return &aclt
}

// Patient converts interface{} to *consult.Patient
func Patient(in interface{}) *patients.Patient {
	b, _ := json.Marshal(in)
	var pt patients.Patient
	json.Unmarshal(b, &pt)
	return &pt
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}
