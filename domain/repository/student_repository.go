package repository

import (
	"wwchacalww/go-cem304/domain/model"
)

type StudentInput struct {
	Name        string `json:"name"`
	BirthDay    string `json:"birth_day"`
	Gender      string `json:"gender"`
	ANNE        string `json:"anne"`
	Note        string `json:"note"`
	Educar      int64  `json:"ieducar"`
	EducaDF     string `json:"educa_df"`
	ClassroomID string `json:"classroom_id"`
	Status      bool   `json:"status"`
	Address     string `json:"address"`
	City        string `json:"city"`
	CEP         string `json:"cep"`
	Fones       string `json:"fones"`
	CPF         string `json:"cpf"`
	Father      string `json:"father"`
	Mother      string `json:"mother"`
	Responsible string `json:"responsible"`
}

type StudentRepositoryInterface interface {
	Create(input StudentInput) (model.StudentInterface, error)
	FindById(id string) (model.StudentInterface, error)
	FindByEducar(educar int64) (model.StudentInterface, error)
	FindByName(name string) ([]model.StudentInterface, error)
	FindByParent(name string) ([]model.StudentInterface, error)
	List(classroom_id string) ([]model.StudentInterface, error)
	Enable(id string) (model.StudentInterface, error)
	Disable(id string) (model.StudentInterface, error)
	ANNE(id, anne string) (model.StudentInterface, error)
	AddMass(mass []StudentInput) ([]model.StudentInterface, error)
	ChangeClassroom(id, classroom_id string) error
}

type StudentPersistence interface {
	Create(study StudentInput) (model.StudentInterface, error)
	FindById(id string) (model.StudentInterface, error)
	FindByEducar(educar int64) (model.StudentInterface, error)
	FindByName(name string) ([]model.StudentInterface, error)
	FindByParent(name string) ([]model.StudentInterface, error)
	List(classroom_id string) ([]model.StudentInterface, error)
	Enable(id string) (model.StudentInterface, error)
	Disable(id string) (model.StudentInterface, error)
	ANNE(id, anne string) (model.StudentInterface, error)
	AddMass(mass []StudentInput) ([]model.StudentInterface, error)
	ChangeClassroom(id, classroom_id string) error
}
