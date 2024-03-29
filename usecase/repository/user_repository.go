package repository

import (
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"

	uuid "github.com/satori/go.uuid"
)

type UserRepository struct {
	Persistence repository.UserPersistenceInterface
}

func NewUserRepository(persistence repository.UserPersistenceInterface) *UserRepository {
	return &UserRepository{Persistence: persistence}
}

func (repo *UserRepository) Create(input repository.UserInput) (model.UserInterface, error) {
	hash, err := utils.HashPassord(input.Password)
	if err != nil {
		return nil, err
	}
	user := model.NewUser()
	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.Password = hash

	err = repo.Persistence.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) FindById(id string) (model.UserInterface, error) {
	user, err := repo.Persistence.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (model.UserInterface, error) {
	user, err := repo.Persistence.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) List() ([]model.UserInterface, error) {
	result, err := repo.Persistence.List()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *UserRepository) ChangePassword(id, pwd string) error {
	hash, err := utils.HashPassord(pwd)
	if err != nil {
		return err
	}
	err = repo.Persistence.ChangePassword(id, hash)
	return err
}

func (repo *UserRepository) ChangeRole(id, role string) error {
	err := repo.Persistence.ChangeRole(id, role)
	return err
}

func (repo *UserRepository) ChangeAvatarUrl(email string) (string, error) {
	avatar_url := uuid.NewV4().String() + ".jpg"
	err := repo.Persistence.ChangeAvatarUrl(email, avatar_url)
	if err != nil {
		return "", err
	}
	return avatar_url, nil
}
