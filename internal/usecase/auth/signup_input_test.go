package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSignUpInput_Success(t *testing.T) {
	validJSON := []byte(`{"usuario":"gabriel","senha":"password123","confirmarSenha":"password123"}`)

	input, err := CreateSignUpInput(validJSON)

	assert.NoError(t, err)
	assert.Equal(t, "gabriel", input.User)
	assert.Equal(t, "password123", input.Password)
	assert.Equal(t, "password123", input.PasswordConfirm)
}

func TestCreateSignUpInput_InvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{"usuario":"gabriel", "senha":"password123", "confirmarSenha":}`)

	input, err := CreateSignUpInput(invalidJSON)

	assert.Error(t, err)
	assert.Empty(t, input)
}

func TestCreateSignUpInput_ValidationError(t *testing.T) {
	invalidInput := []byte(`{"usuario":"","senha":"password123","confirmarSenha":"password123"}`)

	input, err := CreateSignUpInput(invalidInput)

	assert.Error(t, err)
	assert.Empty(t, input)
}
