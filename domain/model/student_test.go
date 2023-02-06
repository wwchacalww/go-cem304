package model_test

import (
	"log"
	"testing"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestStudent(t *testing.T) {
	study, err := model.NewStudent("Fulano de Tal", "29/11/2016", 234123)
	require.Nil(t, err)
	require.NotNil(t, study.GetID())
	require.Equal(t, study.GetName(), "Fulano de Tal")

	// log.Println("UpdateAt: ", study.UpdatedAt)
	// log.Println("CreatedAt: ", study.CreatedAt)

	// require.NotNil(t, study.GetID())
	// require.Equal(t, study.GetName(), "1º ano A - Matutino")

	study.Name = "io"
	invalid, err := study.IsValid()
	require.False(t, invalid)
	require.NotNil(t, err)
	log.Println(invalid)
	study.Name = "Beltrano"
	valid, err := study.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	study.Disable()
	require.False(t, study.GetStatus())
	study.Enable()
	require.True(t, study.GetStatus())

	// location, err := time.LoadLocation("America/Sao_Paulo")
	// data := time.Date(2023, time.February, 4, 11, 49, 0, 0, location)
	// study.UpdatedAt = data
	// require.NotEqual(t, study.CreatedAt, study.UpdatedAt)
	// log.Println("UpdateAt: ", study.UpdatedAt)

}
