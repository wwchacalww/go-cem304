package model_test

import (
	"log"
	"testing"
	"time"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestClassroom(t *testing.T) {
	class := model.NewClassroom()
	class.Name = "1º ano A - Matutino"
	class.Grade = "Ensino Médio"
	class.Shift = "Matutino"
	class.Description = "Descrição da turma"
	class.ANNE = "TDG DPA(C)"
	class.Year = "2022"

	log.Println("UpdateAt: ", class.UpdatedAt)
	log.Println("CreatedAt: ", class.CreatedAt)

	require.NotNil(t, class.GetID())
	require.Equal(t, class.GetName(), "1º ano A - Matutino")

	invalid, err := class.IsValid()
	require.False(t, invalid)
	require.NotNil(t, err)

	class.Level = "1º ano"
	valid, err := class.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	class.Disable()
	require.False(t, class.GetStatus())
	class.Enable()
	require.True(t, class.GetStatus())

	location, err := time.LoadLocation("America/Sao_Paulo")
	data := time.Date(2023, time.February, 4, 11, 49, 0, 0, location)
	class.UpdatedAt = data
	require.NotEqual(t, class.CreatedAt, class.UpdatedAt)
	log.Println("UpdateAt: ", class.UpdatedAt)

}
