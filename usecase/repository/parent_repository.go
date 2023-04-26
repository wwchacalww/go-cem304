package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type ParentRepository struct {
	Persistence repository.ParentPersistence
}

func NewParentRepository(persistence repository.ParentPersistence) *ParentRepository {
	return &ParentRepository{Persistence: persistence}
}

func (repo *ParentRepository) Create(input repository.ParentInput) (model.ParentInterface, error) {
	parent, err := repo.Persistence.Create(input)
	if err != nil {
		return nil, err
	}

	return parent, nil
}

func (repo *ParentRepository) FindById(id string) (model.ParentInterface, error) {
	parent, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Parent{}, err
	}

	return parent, nil
}

func (repo *ParentRepository) FindByCPF(cpf string) (model.ParentInterface, error) {
	parent, err := repo.Persistence.FindByCPF(cpf)
	if err != nil {
		return &model.Parent{}, err
	}

	return parent, nil
}

func (repo *ParentRepository) FindByName(name string) ([]model.ParentInterface, error) {
	paretens, err := repo.Persistence.FindByName(name)
	if err != nil {
		return nil, err
	}

	return paretens, nil
}

func (repo *ParentRepository) ChangeFone(id, fones string) error {
	err := repo.Persistence.ChangeFone(id, fones)
	return err
}

func (repo *ParentRepository) ChangeEmail(id, email string) error {
	err := repo.Persistence.ChangeEmail(id, email)
	return err
}

func (repo *ParentRepository) ChangeCPF(id, cpf string) error {
	err := repo.Persistence.ChangeCPF(id, cpf)
	return err
}

func (repo *ParentRepository) ChangeResponsible(id, student_id string, status bool) error {
	err := repo.Persistence.ChangeResponsible(id, student_id, status)
	return err
}
