package repository

import "wwchacalww/go-cem304/domain/model"

type ParentInput struct {
	Name         string `json:"name" valid:"required,stringlength(5|20)"`
	Relationship string `json:"relationship" valid:"alphanum,required"`
	CPF          string `json:"cpf" valid:"optional"`
	Email        string `json:"email" valid:"optional,email"`
	Fones        string `json:"fones" valid:"optional"`
	Responsible  bool   `json:"responsible" valid:"optional"`
	Student_id   string `json:"student_id" valid:"required"`
}

type ParentRepositoryInterface interface {
	Create(input ParentInput) (model.ParentInterface, error)
	FindById(id string) (model.ParentInterface, error)
	FindByCPF(cpf string) (model.ParentInterface, error)
	FindByName(name string) ([]model.ParentInterface, error)
	ChangeFone(id, fones string) error
	ChangeEmail(id, email string) error
	ChangeCPF(id, cpf string) error
	ChangeResponsible(id, student_id string, status bool) error
}

type ParentPersistence interface {
	Create(input ParentInput) (model.ParentInterface, error)
	FindById(id string) (model.ParentInterface, error)
	FindByName(name string) ([]model.ParentInterface, error)
	FindByCPF(cpf string) (model.ParentInterface, error)
	ChangeFone(id, fones string) error
	ChangeEmail(id, email string) error
	ChangeCPF(id, cpf string) error
	ChangeResponsible(id, student_id string, status bool) error
}
