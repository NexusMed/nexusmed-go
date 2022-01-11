package patientsiface

import "github.com/nexusmed/nexusmed-go/patients"

type PatientsAPI interface {
	// Get patient by Id
	GetPatient(id string) (*patients.GetPatient, error)

	// Get paginated patients
	GetPatients(limit *int, nextToken *string) (*patients.GetPatients, error)

	// Create a new patient
	CreatePatient(patients.CreatePatientInput) (*patients.CreatePatient, error)
}
