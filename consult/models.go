// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package consult

import (
	"fmt"
	"io"
	"strconv"
)

type Consultation interface {
	IsConsultation()
}

type IConsultation interface {
	IsIConsultation()
}

type Answer struct {
	Index  int    `json:"index"`
	Value  string `json:"value"`
	Reject bool   `json:"reject"`
}

type AnswerInput struct {
	Choice []*int  `json:"choice,omitempty"`
	Text   *string `json:"text,omitempty"`
}

type AnswerQuestionnaireInput struct {
	Patient         *PatientInput  `json:"patient,omitempty"`
	QuestionnaireID string         `json:"questionnaire_id"`
	Answers         []*AnswerInput `json:"answers,omitempty"`
}

type AsynchronousConsultation struct {
	ID                   string                `json:"id"`
	Patient              *Patient              `json:"patient,omitempty"`
	Prescriber           *Prescriber           `json:"prescriber,omitempty"`
	Status               ConsultationStatus    `json:"status"`
	Rejected             *bool                 `json:"rejected,omitempty"`
	RejectReason         *string               `json:"reject_reason,omitempty"`
	Products             []*Product            `json:"products,omitempty"`
	QuestionnaireAnswers *QuestionnaireAnswers `json:"questionnaire_answers,omitempty"`
}

func (AsynchronousConsultation) IsIConsultation() {}
func (AsynchronousConsultation) IsConsultation()  {}

type CreateAnswer struct {
	Value  string `json:"value"`
	Reject *bool  `json:"reject,omitempty"`
}

type CreateConsultationInput struct {
	Type                 ConsultationInputType      `json:"type"`
	Patient              *PatientInput              `json:"patient,omitempty"`
	Products             []*ProductInput            `json:"products,omitempty"`
	QuestionnaireAnswers *QuestionnaireAnswersInput `json:"questionnaire_answers,omitempty"`
}

type CreateQuestionnaireInput struct {
	Title     *string          `json:"title,omitempty"`
	Questions []*QuestionInput `json:"questions,omitempty"`
}

type Dosage struct {
	Quantity *float64    `json:"quantity,omitempty"`
	Unit     *DosageUnit `json:"unit,omitempty"`
}

type Medication struct {
	Name     *string `json:"name,omitempty"`
	Dosage   *Dosage `json:"dosage,omitempty"`
	Quantity *int    `json:"quantity,omitempty"`
}

type Name struct {
	Title      *string `json:"title,omitempty"`
	GivenName  *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
}

type Patient struct {
	ID       string  `json:"id"`
	Name     *Name   `json:"name,omitempty"`
	StripeID *string `json:"stripe_id,omitempty"`
}

type PatientInput struct {
	ID string `json:"id"`
}

type Prescriber struct {
	ID       string    `json:"id"`
	Name     *Name     `json:"name,omitempty"`
	Register *Register `json:"register,omitempty"`
}

type Product struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Medication *Medication `json:"medication,omitempty"`
	Quantity   *int        `json:"quantity,omitempty"`
}

type ProductInput struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type Question struct {
	Index    int          `json:"index"`
	Text     string       `json:"text"`
	Info     []*string    `json:"info,omitempty"`
	Type     QuestionType `json:"type"`
	Answers  []*Answer    `json:"answers,omitempty"`
	Required *bool        `json:"required,omitempty"`
}

type QuestionInput struct {
	Type    QuestionType    `json:"type"`
	Text    string          `json:"text"`
	Info    []*string       `json:"info,omitempty"`
	Answers []*CreateAnswer `json:"answers,omitempty"`
}

type Questionnaire struct {
	ID        string      `json:"id"`
	Title     *string     `json:"title,omitempty"`
	Questions []*Question `json:"questions,omitempty"`
	CreatedAt *string     `json:"created_at,omitempty"`
}

type QuestionnaireAnswer struct {
	Question *Question `json:"question,omitempty"`
	Value    []string  `json:"value,omitempty"`
	Choice   []*int    `json:"choice,omitempty"`
	Text     *string   `json:"text,omitempty"`
}

type QuestionnaireAnswers struct {
	ID            string                 `json:"id"`
	Questionnaire *Questionnaire         `json:"questionnaire,omitempty"`
	Answers       []*QuestionnaireAnswer `json:"answers,omitempty"`
	CreatedAt     *string                `json:"created_at,omitempty"`
}

type QuestionnaireAnswersInput struct {
	ID string `json:"id"`
}

type Questionnaires struct {
	Items     []*Questionnaire `json:"items,omitempty"`
	NextToken *string          `json:"next_token,omitempty"`
}

type Register struct {
	Type  *RegisterType `json:"type,omitempty"`
	Value *string       `json:"value,omitempty"`
}

type ConsultationInputType string

const (
	ConsultationInputTypeAsynchronous ConsultationInputType = "asynchronous"
	ConsultationInputTypeSynchronous  ConsultationInputType = "synchronous"
	ConsultationInputTypeInPerson     ConsultationInputType = "in_person"
)

var AllConsultationInputType = []ConsultationInputType{
	ConsultationInputTypeAsynchronous,
	ConsultationInputTypeSynchronous,
	ConsultationInputTypeInPerson,
}

func (e ConsultationInputType) IsValid() bool {
	switch e {
	case ConsultationInputTypeAsynchronous, ConsultationInputTypeSynchronous, ConsultationInputTypeInPerson:
		return true
	}
	return false
}

func (e ConsultationInputType) String() string {
	return string(e)
}

func (e *ConsultationInputType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConsultationInputType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConsultationInputType", str)
	}
	return nil
}

func (e ConsultationInputType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConsultationStatus string

const (
	ConsultationStatusPending    ConsultationStatus = "pending"
	ConsultationStatusConfirmed  ConsultationStatus = "confirmed"
	ConsultationStatusAssigned   ConsultationStatus = "assigned"
	ConsultationStatusUnassigned ConsultationStatus = "unassigned"
	ConsultationStatusAccepted   ConsultationStatus = "accepted"
	ConsultationStatusRejected   ConsultationStatus = "rejected"
	ConsultationStatusStarted    ConsultationStatus = "started"
	ConsultationStatusAbandoned  ConsultationStatus = "abandoned"
	ConsultationStatusCompleted  ConsultationStatus = "completed"
	ConsultationStatusExpired    ConsultationStatus = "expired"
	ConsultationStatusCancelled  ConsultationStatus = "cancelled"
)

var AllConsultationStatus = []ConsultationStatus{
	ConsultationStatusPending,
	ConsultationStatusConfirmed,
	ConsultationStatusAssigned,
	ConsultationStatusUnassigned,
	ConsultationStatusAccepted,
	ConsultationStatusRejected,
	ConsultationStatusStarted,
	ConsultationStatusAbandoned,
	ConsultationStatusCompleted,
	ConsultationStatusExpired,
	ConsultationStatusCancelled,
}

func (e ConsultationStatus) IsValid() bool {
	switch e {
	case ConsultationStatusPending, ConsultationStatusConfirmed, ConsultationStatusAssigned, ConsultationStatusUnassigned, ConsultationStatusAccepted, ConsultationStatusRejected, ConsultationStatusStarted, ConsultationStatusAbandoned, ConsultationStatusCompleted, ConsultationStatusExpired, ConsultationStatusCancelled:
		return true
	}
	return false
}

func (e ConsultationStatus) String() string {
	return string(e)
}

func (e *ConsultationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConsultationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConsultationStatus", str)
	}
	return nil
}

func (e ConsultationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DosageUnit string

const (
	DosageUnitMg  DosageUnit = "mg"
	DosageUnitMcg DosageUnit = "mcg"
	DosageUnitMl  DosageUnit = "ml"
	DosageUnitIu  DosageUnit = "iu"
)

var AllDosageUnit = []DosageUnit{
	DosageUnitMg,
	DosageUnitMcg,
	DosageUnitMl,
	DosageUnitIu,
}

func (e DosageUnit) IsValid() bool {
	switch e {
	case DosageUnitMg, DosageUnitMcg, DosageUnitMl, DosageUnitIu:
		return true
	}
	return false
}

func (e DosageUnit) String() string {
	return string(e)
}

func (e *DosageUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DosageUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DosageUnit", str)
	}
	return nil
}

func (e DosageUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type QuestionType string

const (
	QuestionTypeFreeText   QuestionType = "free_text"
	QuestionTypeOneOfMany  QuestionType = "one_of_many"
	QuestionTypeManyOfMany QuestionType = "many_of_many"
)

var AllQuestionType = []QuestionType{
	QuestionTypeFreeText,
	QuestionTypeOneOfMany,
	QuestionTypeManyOfMany,
}

func (e QuestionType) IsValid() bool {
	switch e {
	case QuestionTypeFreeText, QuestionTypeOneOfMany, QuestionTypeManyOfMany:
		return true
	}
	return false
}

func (e QuestionType) String() string {
	return string(e)
}

func (e *QuestionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = QuestionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QuestionType", str)
	}
	return nil
}

func (e QuestionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RegisterType string

const (
	RegisterTypeGmc  RegisterType = "gmc"
	RegisterTypeGphc RegisterType = "gphc"
	RegisterTypeNmc  RegisterType = "nmc"
)

var AllRegisterType = []RegisterType{
	RegisterTypeGmc,
	RegisterTypeGphc,
	RegisterTypeNmc,
}

func (e RegisterType) IsValid() bool {
	switch e {
	case RegisterTypeGmc, RegisterTypeGphc, RegisterTypeNmc:
		return true
	}
	return false
}

func (e RegisterType) String() string {
	return string(e)
}

func (e *RegisterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RegisterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RegisterType", str)
	}
	return nil
}

func (e RegisterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
