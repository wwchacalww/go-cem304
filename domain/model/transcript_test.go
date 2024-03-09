package model_test

import (
	"testing"
	"wwchacalww/go-cem304/domain/model"

	"github.com/stretchr/testify/require"
)

func TestTranscriptIsValid_Valid(t *testing.T) {
	transcript := model.NewTranscript(
		"História",
		"Aprovado",
		"8.5",
		"1ª Série",
		"FGB",
		10,
		40,
		2024,
	)

	require.Equal(t, transcript.GetSubject(), "História")

	valid, err := transcript.IsValid()
	require.True(t, valid)
	require.Nil(t, err)
}
