package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ParentInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetRelationship() string
	GetCPF() string
	GetEmail() string
	GetFones() string
	GetResponsible() bool
	GetStudents() []StudentInterface
}

type Parent struct {
	ID           string             `json:"id" valid:"uuidv4"`
	Name         string             `json:"name" valid:"required,stringlength(5|20)"`
	Relationship string             `json:"relationship" valid:"alphanum,required"`
	CPF          string             `json:"cpf" valid:"optional"`
	Email        string             `json:"email" valid:"optional,email"`
	Fones        string             `json:"fones" valid:"optional"`
	Responsible  bool               `json:"responsible" valid:"optional"`
	Students     []StudentInterface `json:"students"`
}

func NewParent(name, relationship string) *Parent {
	parent := Parent{
		ID:           uuid.NewV4().String(),
		Name:         name,
		Relationship: relationship,
		Responsible:  false,
	}

	return &parent
}

func (p *Parent) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Parent) GetID() string {
	return p.ID
}

func (p *Parent) GetName() string {
	return p.Name
}

func (p *Parent) GetRelationship() string {
	return p.Relationship
}

func (p *Parent) GetEmail() string {
	return p.Email
}

func (p *Parent) GetCPF() string {
	return p.CPF
}

func (p *Parent) GetFones() string {
	return p.Fones
}

func (p *Parent) GetResponsible() bool {
	return p.Responsible
}

func (p *Parent) GetStudents() []StudentInterface {
	return p.Students
}
