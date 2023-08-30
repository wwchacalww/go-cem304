package model

import (
	"github.com/asaskevich/govalidator"
)

type EvaluationInterface interface {
	IsValid() (bool, error)
	GetID() int
	GetSubject() SubjectInterface
	GetStudent() StudentInterface
	GetTerm() string
	GetNote() string
	GetAbsences() int
}

type Evaluation struct {
	ID       int              `json:"id" valid:"optional"`
	Subject  SubjectInterface `json:"subject_id" valid:"required,stringlength(5|30)"`
	Student  StudentInterface `json:"student_id" valid:"required,stringlength(5|30)"`
	Term     string           `json:"term" valid:"required,stringlength(5|30)"`
	Note     string           `json:"note" valid:"required"`
	Absences int              `json:"absences" valid:"required"`
}

func NewEvaluation(term string, note string, id, absences int) *Evaluation {
	Evaluation := Evaluation{
		ID:       id,
		Term:     term,
		Note:     note,
		Absences: absences,
	}

	return &Evaluation
}

func (e *Evaluation) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(e)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *Evaluation) GetID() int {
	return e.ID
}

func (e *Evaluation) GetSubject() SubjectInterface {
	return e.Subject
}

func (e *Evaluation) GetStudent() StudentInterface {
	return e.Student
}

func (e *Evaluation) GetTerm() string {
	return e.Term
}

func (e *Evaluation) GetNote() string {
	return e.Note
}

func (e *Evaluation) GetAbsences() int {
	return e.Absences
}
