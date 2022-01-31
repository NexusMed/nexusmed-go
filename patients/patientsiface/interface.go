package patientsiface

import (
	"github.com/nexusmed/nexusmed-go/client"
	"github.com/nexusmed/nexusmed-go/patients"
)

type PatientsAPI interface {
	// Get patient by Id
	GetPatient(id string, interceptors ...client.RequestInterceptor) (*patients.GetPatient, error)

	// Get paginated patients
	GetPatients(limit *int, nextToken *string, interceptors ...client.RequestInterceptor) (*patients.GetPatients, error)

	// Create a new patient
	CreatePatient(patients.CreatePatientInput, ...client.RequestInterceptor) (*patients.CreatePatient, error)

	// Set the Authorization header
	SetApiKey(key string)
}
