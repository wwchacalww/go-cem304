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
}

type StudentRepositoryInterface interface {
	Create(input StudentInput) (model.StudentInterface, error)
	FindById(id string) (model.StudentInterface, error)
	FindByName(name string) ([]model.StudentInterface, error)
	List(classroom_id string) ([]model.StudentInterface, error)
	Enable(id string) (model.StudentInterface, error)
	Disable(id string) (model.StudentInterface, error)
	ANNE(id, anne string) (model.StudentInterface, error)
	AddMass(mass []StudentInput) ([]model.StudentInterface, error)
}

type StudentPersistence interface {
	Create(study StudentInput) (model.StudentInterface, error)
	FindById(id string) (model.StudentInterface, error)
	FindByName(name string) ([]model.StudentInterface, error)
	List(classroom_id string) ([]model.StudentInterface, error)
	Enable(id string) (model.StudentInterface, error)
	Disable(id string) (model.StudentInterface, error)
	ANNE(id, anne string) (model.StudentInterface, error)
	AddMass(mass []StudentInput) ([]model.StudentInterface, error)
}
