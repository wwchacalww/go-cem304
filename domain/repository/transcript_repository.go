package repository

import "wwchacalww/go-cem304/domain/model"

type TranscriptInput struct {
	EducaDF   string `json:"educa_df"`
	Level     string `json:"level"`
	Subject   string `json:"subject"`
	Result    string `json:"result"`
	Note      string `json:"note"`
	Absences  int    `json:"absences"`
	Workload  int    `json:"workload"`
	Formation string `json:"formation"`
	Year      int    `json:"year"`
}

type TranscriptRepositoryInterface interface {
	AddMass(mass []TranscriptInput) error
	ListStudentTranscript(student_id string) ([]model.TranscriptInterface, error)
}

type TranscriptPersistence interface {
	Save(input TranscriptInput) (model.TranscriptInterface, error)
	ListStudentTranscript(student_id string) ([]model.TranscriptInterface, error)
}
