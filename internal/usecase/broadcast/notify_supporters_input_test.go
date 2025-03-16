package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotifySupportersInput_Success(t *testing.T) {
	validJSON := []byte(`{"tipo":"fim","time":"Coritiba","placar":"2-2","mensagem":"Jogo finalizado!"}`)

	input, err := CreateNotifySupportersInput(validJSON)

	assert.NoError(t, err)
	assert.Equal(t, "fim", input.Type)
	assert.Equal(t, "Coritiba", input.Team)
	assert.Equal(t, "2-2", input.Score)
	assert.Equal(t, "Jogo finalizado!", input.Message)
}

func TestCreateNotifySupportersInput_InvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{"tipo":"fim","time":"Coritiba", "mensagem":}`)

	input, err := CreateNotifySupportersInput(invalidJSON)

	assert.Error(t, err)
	assert.Empty(t, input)
}

func TestCreateNotifySupportersInput_ValidationError(t *testing.T) {
	invalidInput := []byte(`{"tipo":"","time":"","placar":"","mensagem":""}`)

	input, err := CreateNotifySupportersInput(invalidInput)

	assert.Error(t, err)
	assert.Empty(t, input)
}
