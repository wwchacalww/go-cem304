package model_test

import (
	"log"
	"testing"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestParentIsValid_Valid(t *testing.T) {
	parent := model.NewParent("Fulano de Tal", "Pai")
	parent.Email = "test@mail.com"
	parent.CPF = "934.233.343-34"

	if parent.Students == nil {
		log.Print("Nulo")
	}
	require.NotNil(t, parent.GetID())
	require.Equal(t, parent.Name, "Fulano de Tal")
	require.Equal(t, parent.Relationship, "Pai")
	require.Equal(t, parent.Email, "test@mail.com")
	require.False(t, parent.Responsible)

	valid, err := parent.IsValid()
	require.True(t, valid)
	require.Nil(t, err)
}

func TestParentIsValid_Invalid(t *testing.T) {
	parent := model.NewParent("Ful", "")
	parent.Email = "wrong-email"
	parent.CPF = "934.233.343-34"

	if parent.Students == nil {
		log.Print("Nulo")
	}
	require.NotNil(t, parent.GetID())
	valid, err := parent.IsValid()
	require.False(t, valid)
	log.Println(err.Error())
	require.Equal(t, err.Error(), "email: wrong-email does not validate as email;name: Ful does not validate as stringlength(5|20);relationship: non zero value required")
}
