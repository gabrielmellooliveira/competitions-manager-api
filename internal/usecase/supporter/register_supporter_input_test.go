package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRegisterSupporterInput_ValidInput(t *testing.T) {
	data := []byte(`{"nome": "gabriel", "email": "gabriel@gmail.com", "time": "Coritiba"}`)
	expected := RegisterSupporterInput{
		Name:  "gabriel",
		Email: "gabriel@gmail.com",
		Team:  "Coritiba",
	}

	input, err := CreateRegisterSupporterInput(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, input)
}

func TestCreateRegisterSupporterInput_InvalidJSON(t *testing.T) {
	data := []byte(`{"usuario": "gabriel", "senha": "gabriel12345"`)

	_, err := CreateRegisterSupporterInput(data)

	assert.Error(t, err)
}

func TestCreateRegisterSupporterInput_MissingFields(t *testing.T) {
	data := []byte(`{"nome": "gabriel", "time": "Coritiba"}`)

	_, err := CreateRegisterSupporterInput(data)

	assert.Error(t, err)
}

func TestCreateRegisterSupporterInput_InvalidFields(t *testing.T) {
	data := []byte(`{"nome": "gabriel", "email": "gabriel12345", "time": "gabriel12345"}`)

	_, err := CreateRegisterSupporterInput(data)

	assert.Error(t, err)
}

func TestCreateRegisterSupporterInput_EmptyInput(t *testing.T) {
	data := []byte(``)

	_, err := CreateRegisterSupporterInput(data)

	assert.Error(t, err)
}
