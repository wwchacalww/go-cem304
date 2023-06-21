package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type SubjectRepository struct {
	Persistence repository.SubjectPersistence
}

func NewSubjectRepository(persistence repository.SubjectPersistence) *SubjectRepository {
	return &SubjectRepository{Persistence: persistence}
}

func (repo *SubjectRepository) Create(input repository.SubjectInput) (model.SubjectInterface, error) {
	subject, err := repo.Persistence.Create(input)
	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (repo *SubjectRepository) FindById(id string) (model.SubjectInterface, error) {
	subject, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Subject{}, err
	}

	return subject, nil
}

func (repo *SubjectRepository) FindByLicense(license string) ([]model.SubjectInterface, error) {
	subjects, err := repo.Persistence.FindByLicense(license)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (repo *SubjectRepository) FindByName(name string) ([]model.SubjectInterface, error) {
	subjects, err := repo.Persistence.FindByName(name)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
