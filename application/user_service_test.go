package application_test

import (
	"testing"
	"wwchacalww/go-cem304/application"
	"wwchacalww/go-cem304/application/dto"
	mock_application "wwchacalww/go-cem304/application/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_application.NewMockUserPersistenceInterface(ctrl)
	persistence.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()

	service := application.UserService{
		Persistence: persistence,
	}

	var input dto.UserInput
	input.Name = "Fulando"
	input.Email = "fulano@email.com"
	input.Password = "123456"
	input.Role = "Paciente"

	result, err := service.Create(input)
	require.Nil(t, err)
	require.Equal(t, result.GetName(), input.Name)
	require.Equal(t, result.GetEmail(), input.Email)
	require.Equal(t, result.GetRole(), input.Role)
	require.Equal(t, result.GetStatus(), true)
}

func TestService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := mock_application.NewMockUserInterface(ctrl)
	persistence := mock_application.NewMockUserPersistenceInterface(ctrl)
	persistence.EXPECT().FindById(gomock.Any()).Return(user, nil).AnyTimes()

	service := application.UserService{
		Persistence: persistence,
	}

	result, err := service.FindById("fake-id")
	require.Nil(t, err)
	require.Equal(t, result, user)
}

func TestService_FindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := mock_application.NewMockUserInterface(ctrl)
	persistence := mock_application.NewMockUserPersistenceInterface(ctrl)
	persistence.EXPECT().FindById(gomock.Any()).Return(user, nil).AnyTimes()

	service := application.UserService{
		Persistence: persistence,
	}

	result, err := service.FindById("fake-email@email.com")
	require.Nil(t, err)
	require.Equal(t, result, user)
}
