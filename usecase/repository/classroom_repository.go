package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type ClassroomRepository struct {
	Persistence repository.ClassroomPersistence
}

func NewClassroomRepository(persistence repository.ClassroomPersistence) *ClassroomRepository {
	return &ClassroomRepository{Persistence: persistence}
}

func (repo *ClassroomRepository) Create(input repository.ClassroomInput) (model.ClassroomInterface, error) {
	class := model.NewClassroom()
	class.Name = input.Name
	class.Level = input.Level
	class.Grade = input.Grade
	class.Shift = input.Shift
	class.Description = input.Description
	class.ANNE = input.ANNE
	class.Year = input.Level

	_, err := class.IsValid()
	if err != nil {
		return nil, err
	}

	err = repo.Persistence.Create(class)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (repo *ClassroomRepository) FindById(id string) (model.ClassroomInterface, error) {
	class, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Classroom{}, err
	}

	return class, nil
}

func (repo *ClassroomRepository) FindByName(name string) ([]model.ClassroomInterface, error) {
	classrooms, err := repo.Persistence.FindByName(name)
	if err != nil {
		return nil, err
	}

	return classrooms, nil
}

func (repo *ClassroomRepository) List(year string) ([]model.ClassroomInterface, error) {
	result, err := repo.Persistence.List(year)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *ClassroomRepository) Enable(id string) (model.ClassroomInterface, error) {
	class, err := repo.Persistence.Enable(id)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (repo *ClassroomRepository) Disable(id string) (model.ClassroomInterface, error) {
	class, err := repo.Persistence.Enable(id)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (repo *ClassroomRepository) ANNE(id, anne string) (model.ClassroomInterface, error) {
	result, err := repo.Persistence.ANNE(id, anne)
	if err != nil {
		return nil, err
	}

	return result, nil
}
