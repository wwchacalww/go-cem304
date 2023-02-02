package application

import (
	"wwchacalww/go-cem304/application/dto"
	"wwchacalww/go-cem304/application/utils"
)

type UserService struct {
	Persistence UserPersistenceInterface
}

func NewUserService(persistence UserPersistenceInterface) *UserService {
	return &UserService{Persistence: persistence}
}

func (s *UserService) Create(input dto.UserInput) (UserInterface, error) {
	hash, err := utils.HashPassord(input.Password)
	if err != nil {
		return nil, err
	}
	user := NewUser()
	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.Password = hash

	err = s.Persistence.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindById(id string) (UserInterface, error) {
	user, err := s.Persistence.FindById(id)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (s *UserService) FindByEmail(email string) (UserInterface, error) {
	user, err := s.Persistence.FindByEmail(email)
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (s *UserService) List() ([]UserInterface, error) {
	result, err := s.Persistence.List()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Authenticate(email, password string) (dto.AuthenticateOutput, error)
// CheckIsValidToken(token, refresh_token string) (dto.AuthenticateOutput, error)
