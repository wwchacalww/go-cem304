package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type TranscriptReporitory struct {
	Persistence repository.TranscriptPersistence
}

func NewTranscriptRepository(persistence repository.TranscriptPersistence) *TranscriptReporitory {
	return &TranscriptReporitory{Persistence: persistence}
}

func (repo *TranscriptReporitory) AddMass(mass []repository.TranscriptInput) error {
	for _, trc := range mass {
		_, err := repo.Persistence.Save(trc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *TranscriptReporitory) ListStudentTranscript(student_id string) ([]model.TranscriptInterface, error) {
	transcricpts, err := repo.Persistence.ListStudentTranscript((student_id))
	if err != nil {
		return nil, err
	}
	return transcricpts, nil
}
