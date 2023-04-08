package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type StudentRepository struct {
	Persistence repository.StudentPersistence
}

func NewStudentRepository(persistence repository.StudentPersistence) *StudentRepository {
	return &StudentRepository{Persistence: persistence}
}

func (repo *StudentRepository) Create(input repository.StudentInput) (model.StudentInterface, error) {
	study, err := repo.Persistence.Create(input)
	if err != nil {
		return nil, err
	}

	return study, nil
}

func (repo *StudentRepository) FindById(id string) (model.StudentInterface, error) {
	study, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Student{}, err
	}

	return study, nil
}

func (repo *StudentRepository) FindByEducar(educar int64) (model.StudentInterface, error) {
	study, err := repo.Persistence.FindByEducar(educar)
	if err != nil {
		return &model.Student{}, err
	}

	return study, nil
}

func (repo *StudentRepository) FindByName(name string) ([]model.StudentInterface, error) {
	Students, err := repo.Persistence.FindByName(name)
	if err != nil {
		return nil, err
	}

	return Students, nil
}

func (repo *StudentRepository) FindByParent(name string) ([]model.StudentInterface, error) {
	Students, err := repo.Persistence.FindByParent(name)
	if err != nil {
		return nil, err
	}

	return Students, nil
}

func (repo *StudentRepository) List(classroom_id string) ([]model.StudentInterface, error) {
	result, err := repo.Persistence.List(classroom_id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *StudentRepository) Enable(id string) (model.StudentInterface, error) {
	study, err := repo.Persistence.Enable(id)
	if err != nil {
		return nil, err
	}

	return study, nil
}

func (repo *StudentRepository) Disable(id string) (model.StudentInterface, error) {
	study, err := repo.Persistence.Disable(id)
	if err != nil {
		return nil, err
	}

	return study, nil
}

func (repo *StudentRepository) ANNE(id, anne string) (model.StudentInterface, error) {
	result, err := repo.Persistence.ANNE(id, anne)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *StudentRepository) AddMass(mass []repository.StudentInput) ([]model.StudentInterface, error) {
	result, err := repo.Persistence.AddMass(mass)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *StudentRepository) ChangeClassroom(id, classroom_id string) error {
	err := repo.Persistence.ChangeClassroom(id, classroom_id)
	if err != nil {
		return err
	}

	return err
}
