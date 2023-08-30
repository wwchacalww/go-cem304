package repository

import (
	"wwchacalww/go-cem304/domain/dtos"
	"wwchacalww/go-cem304/domain/model"
)

type EvaluationInput struct {
	Student  string `valid:"required,stringlength(5|150)" json:"student_id"`
	Subject  int    `valid:"required,stringlength(5|50)" json:"subject_id"`
	Term     string `valid:"required" json:"term"`
	Note     string `valid:"required" json:"note"`
	Absences int    `valid:"required" json:"absences"`
}

type EvaluationRepositoryInterface interface {
	Create(input EvaluationInput) (model.EvaluationInterface, error)
	Update(note string, id, absences int) (model.EvaluationInterface, error)
	ImportCSVEvaluation(input dtos.InputImportCard) error
	// FindById(id int) (model.EvaluationInterface, error)
	// FindByEducar(educar int) ([]model.EvaluationInterface, error)
	// FindBySubject(subject_id int) ([]model.EvaluationInterface, error)
}

type EvaluationPersistence interface {
	Create(input EvaluationInput) error
	Update(note string, id, absences int) error
	FindEvaluation(student_id, term string, subject_id int) (model.EvaluationInterface, error)
	FindById(id int) (model.EvaluationInterface, error)
	CheckEvaluationExists(student_id, term string, subject int) (int, error)
	GetAndCheckStudent(classroom_id, educadf string) (string, error)
	GetSubjectId(classroom_id, name string) (int, error)
	// FindByEducar(educar int) ([]model.EvaluationInterface, error)
	// FindBySubject(subject_id int) ([]model.EvaluationInterface, error)
}
