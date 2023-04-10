package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type StudentInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetBirthDay() time.Time
	GetGender() string
	GetANNE() string
	GetNote() string
	GetEducar() int64
	GetEducaDF() string
	GetClassroom() ClassroomInterface
	GetStatus() bool
	GetAddress() string
	GetCity() string
	GetCEP() string
	GetFones() string
	GetCPF() string
	GetParents() []ParentInterface
	AttachParent(parent Parent) error
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Enable() error
	Disable() error
}

type Student struct {
	ID        string             `valid:"uuidv4" json:"id"`
	Name      string             `valid:"required,stringlength(5|50)" json:"Name"`
	BirthDay  time.Time          `valid:"required" json:"birth_day"`
	Gender    string             `valid:"optional" json:"gender"`
	ANNE      string             `valid:"optional" json:"anne"`
	Note      string             `valid:"optional" json:"note"`
	Educar    int64              `valid:"required" json:"ieducar"`
	EducaDF   string             `valid:"optional" json:"educa_df"`
	Classroom ClassroomInterface `valid:"optional" json:"classroom"`
	Status    bool               `valid:"optional" json:"status"`
	Address   string             `valid:"optional" json:"address"`
	City      string             `valid:"optional" json:"city"`
	CEP       string             `valid:"optional" json:"cep"`
	Fones     string             `valid:"optional" json:"fones"`
	CPF       string             `valid:"optional" json:"cpf"`
	Parents   []ParentInterface
	CreatedAt time.Time `valid:"optional" json:"created_at"`
	UpdatedAt time.Time `valid:"optional" json:"updated_at"`
}

func NewStudent(name, bd string, educar int64) (*Student, error) {
	dt := strings.Split(bd, "/")
	if len(dt) != 3 {
		return &Student{}, fmt.Errorf("Birth day invalid")
	}
	y, err := strconv.Atoi(dt[2])
	if err != nil {
		return &Student{}, err
	}
	m, err := strconv.Atoi(dt[1])
	if err != nil {
		return &Student{}, err
	}
	d, err := strconv.Atoi(dt[0])

	if err != nil {
		return &Student{}, err
	}
	birth_day := time.Date(y, time.Month(m), d, 12, 15, 5, 5, time.Local)
	Student := Student{
		ID:        uuid.NewV4().String(),
		Name:      name,
		BirthDay:  birth_day,
		Educar:    educar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    true,
	}

	return &Student, nil
}

func (s *Student) IsValid() (bool, error) {
	if s.Status != false && s.Status != true {
		s.Status = false
	}

	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Student) GetID() string {
	return s.ID
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetBirthDay() time.Time {
	return s.BirthDay
}

func (s *Student) GetGender() string {
	return s.Gender
}

func (s *Student) GetANNE() string {
	return s.ANNE
}

func (s *Student) GetNote() string {
	return s.Note
}

func (s *Student) GetEducar() int64 {
	return s.Educar
}

func (s *Student) GetEducaDF() string {
	return s.EducaDF
}

func (s *Student) GetClassroom() ClassroomInterface {
	return s.Classroom
}

func (s *Student) GetStatus() bool {
	return s.Status
}

func (s *Student) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *Student) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *Student) Enable() error {
	s.Status = true
	return nil
}

func (s *Student) Disable() error {
	s.Status = false
	return nil
}

func (s *Student) GetAddress() string {
	return s.Address
}

func (s *Student) GetCity() string {
	return s.City
}

func (s *Student) GetCEP() string {
	return s.CEP
}

func (s *Student) GetFones() string {
	return s.Fones
}

func (s *Student) GetCPF() string {
	return s.CPF
}

func (s *Student) GetParents() []ParentInterface {
	return s.Parents
}

func (s *Student) AttachParent(parent Parent) error {
	if s.Parents == nil {
		s.Parents = []ParentInterface{&parent}
		return nil
	}
	s.Parents = append(s.Parents, &parent)
	return nil
}
