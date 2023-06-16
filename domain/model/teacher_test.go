package model_test

import (
	"log"
	"testing"
	"time"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestTeacherIsValid_Valid(t *testing.T) {
	teacher := model.NewTeacher("Fulano de Tal", "Fulano", "555.555.555-55", "Língua Portuguesa")
	teacher.Email = "test@mail.com"
	teacher.Fones = "(61) 999.999.999"
	teacher.Gender = "Masculino"
	teacher.BirthDay = time.Now()

	require.NotNil(t, teacher.GetID())
	require.Equal(t, teacher.Name, "Fulano de Tal")
	require.Equal(t, teacher.Nick, "Fulano")
	require.Equal(t, teacher.Email, "test@mail.com")
	require.Equal(t, teacher.License, "Língua Portuguesa")

	valid, err := teacher.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	teacher.Save("", "07/06/2003", "Masculino", "(61) 988.888.888", "História", "")
	require.Equal(t, teacher.Gender, "Masculino")
	require.Equal(t, teacher.License, "História")
	require.Equal(t, teacher.Fones, "(61) 988.888.888")
	log.Println(teacher.BirthDay, teacher.Note)
}
