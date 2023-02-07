package repository

import (
	"wwchacalww/go-cem304/domain/model"
)

type ClassroomInput struct {
	Name        string `json:"name"`
	Level       string `json:"level"`
	Grade       string `json:"grade"`
	Shift       string `json:"shift"`
	Description string `json:"description"`
	ANNE        string `json:"anne"`
	Year        string `json:"year"`
}

type ClassroomRepositoryInterface interface {
	Create(input ClassroomInput) (model.ClassroomInterface, error)
	FindById(id string) (model.ClassroomInterface, error)
	FindByName(name string) ([]model.ClassroomInterface, error)
	List(year string) ([]model.ClassroomInterface, error)
	Enable(id string) (model.ClassroomInterface, error)
	Disable(id string) (model.ClassroomInterface, error)
	ANNE(id, anne string) (model.ClassroomInterface, error)
	AddMass(mass []ClassroomInput) ([]model.ClassroomInterface, error)
}

type ClassroomPersistence interface {
	Create(class model.ClassroomInterface) error
	FindById(id string) (model.ClassroomInterface, error)
	FindByName(name string) ([]model.ClassroomInterface, error)
	List(year string) ([]model.ClassroomInterface, error)
	Enable(id string) (model.ClassroomInterface, error)
	Disable(id string) (model.ClassroomInterface, error)
	ANNE(id, anne string) (model.ClassroomInterface, error)
	AddMass(mass []model.ClassroomInterface) ([]model.ClassroomInterface, error)
}
