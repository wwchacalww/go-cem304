package model_test

import (
	"testing"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestSubjectIsValid_Valid(t *testing.T) {
	subject := model.NewSubject("História", "História", "1º ano")
	subject.CH = 80
	subject.Semester = "Bloco A"

	require.NotNil(t, subject.GetID())
	require.Equal(t, subject.Name, "História")
	require.Equal(t, subject.Level, "1º ano")
	require.Equal(t, subject.Grade, "Ensino Médio")
	require.Equal(t, subject.License, "História")
	require.Equal(t, subject.Year, 2023)

	valid, err := subject.IsValid()
	require.True(t, valid)
	require.Nil(t, err)
}
