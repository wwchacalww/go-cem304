package model_test

import (
	"testing"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestEvaluationIsValid_Valid(t *testing.T) {
	evaluation := model.NewEvaluation("1º Bimestre", "8,5", 121, 12)

	require.Equal(t, evaluation.GetID(), 121)
	require.Equal(t, evaluation.GetTerm(), "1º Bimestre")
	require.Equal(t, evaluation.GetAbsences(), 12)
	require.Equal(t, evaluation.GetNote(), "8,5")

	valid, err := evaluation.IsValid()
	require.True(t, valid)
	require.Nil(t, err)
}
