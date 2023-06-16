package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type TeacherRepository struct {
	Persistence repository.TeacherPersistence
}

func NewTeacherRepository(persistence repository.TeacherPersistence) *TeacherRepository {
	return &TeacherRepository{Persistence: persistence}
}

func (repo *TeacherRepository) Create(input repository.TeacherInput) (model.TeacherInterface, error) {
	teacher, err := repo.Persistence.Create(input)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (repo *TeacherRepository) FindById(id string) (model.TeacherInterface, error) {
	Teacher, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Teacher{}, err
	}

	return Teacher, nil
}

func (repo *TeacherRepository) FindByCPF(cpf string) (model.TeacherInterface, error) {
	Teacher, err := repo.Persistence.FindByCPF(cpf)
	if err != nil {
		return &model.Teacher{}, err
	}

	return Teacher, nil
}

func (repo *TeacherRepository) FindByName(name string) ([]model.TeacherInterface, error) {
	paretens, err := repo.Persistence.FindByName(name)
	if err != nil {
		return nil, err
	}

	return paretens, nil
}

func (repo *TeacherRepository) Save(id string, input repository.TeacherInput) (model.TeacherInterface, error) {
	teacher, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Teacher{}, err
	}

	teacher.Save(input.Email, input.BirthDay, input.Gender, input.Fones, input.License, input.Note)
	err = repo.Persistence.Save(teacher)
	if err != nil {
		return &model.Teacher{}, err
	}
	return teacher, nil
}
