package repository

import (
	"wwchacalww/go-cem304/domain/model"
)

type ClassroomInput struct {
	Name        string `json:"Name"`
	Level       string `json:"level"`
	Grade       string `json:"grade"`
	Shift       string `json:"shift"`
	Description string `json:"description"`
	ANNE        string `json:"ANNE"`
	Year        string `json:"year"`
}

type ClassroomRepositoryInterface interface {
	Create(input ClassroomInput) (model.ClassroomInterface, error)
	FindById(id string) (model.ClassroomInterface, error)
	FindByName(name string) ([]model.ClassroomInterface, error)
	list(year string) ([]model.ClassroomInterface, error)
	Enable(id string) (model.ClassroomInterface, error)
	Disable(id string) (model.ClassroomInterface, error)
	ANNE(id, anne string) (model.ClassroomInterface, error)
}

type ClassroomPersistence interface {
	Create(class model.ClassroomInterface) error
	FindById(id string) (model.ClassroomInterface, error)
	FindByName(name string) ([]model.ClassroomInterface, error)
	List(year string) ([]model.ClassroomInterface, error)
	Enable(id string) (model.ClassroomInterface, error)
	Disable(id string) (model.ClassroomInterface, error)
	ANNE(id, anne string) (model.ClassroomInterface, error)
}
