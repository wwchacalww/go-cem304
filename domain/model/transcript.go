package model

import (
	"github.com/asaskevich/govalidator"
)

type TranscriptInterface interface {
	IsValid() (bool, error)
	GetID() int
	GetStudent() StudentInterface
	GetSubject() string
	GetResult() string
	GetNote() string
	GetAbsences() int
	GetWorkload() int
	GetLevel() string
	GetFormation() string
	GetYear() int
}

type Transcript struct {
	ID        int              `json:"id" valid:"optional"`
	Student   StudentInterface `json:"student_id" valid:"required,stringlength(5|30)"`
	Subject   string           `json:"subject" valid:"required,stringlength(3|40)"`
	Result    string           `json:"result" valid:"required,stringlength(3|20)"`
	Note      string           `json:"note" valid:"required"`
	Absences  int              `json:"absences" valid:"required"`
	Workload  int              `json:"workload" valid:"required"`
	Level     string           `json:"level" valid:"required"`
	Formation string           `json:"formation" valid:"required"`
	Year      int              `json:"year" valid:"required"`
}

func NewTranscript(subject, result, note, level, formation string, absences, workload, year int) *Transcript {
	transcript := Transcript{
		Subject:   subject,
		Result:    result,
		Note:      note,
		Absences:  absences,
		Workload:  workload,
		Level:     level,
		Formation: formation,
		Year:      year,
	}
	return &transcript
}

func (t *Transcript) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *Transcript) GetID() int {
	return t.ID
}

func (t *Transcript) GetStudent() StudentInterface {
	return t.Student
}

func (t *Transcript) GetSubject() string {
	return t.Subject
}

func (t *Transcript) GetResult() string {
	return t.Result
}

func (t *Transcript) GetNote() string {
	return t.Note
}

func (t *Transcript) GetAbsences() int {
	return t.Absences
}

func (t *Transcript) GetWorkload() int {
	return t.Workload
}

func (t *Transcript) GetLevel() string {
	return t.Level
}

func (t *Transcript) GetFormation() string {
	return t.Formation
}

func (t *Transcript) GetYear() int {
	return t.Year
}
