package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLoginInput_ValidInput(t *testing.T) {
	data := []byte(`{"usuario": "gabriel", "senha": "gabriel12345"}`)
	expected := LoginInput{Name: "gabriel", Password: "gabriel12345"}

	input, err := CreateLoginInput(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, input)
}

func TestCreateLoginInput_InvalidJSON(t *testing.T) {
	data := []byte(`{"usuario": "gabriel", "senha": "gabriel12345"`)

	_, err := CreateLoginInput(data)

	assert.Error(t, err)
}

func TestCreateLoginInput_MissingFields(t *testing.T) {
	data := []byte(`{"usuario": "gabriel"}`)

	_, err := CreateLoginInput(data)

	assert.Error(t, err)
}

func TestCreateLoginInput_InvalidFields(t *testing.T) {
	data := []byte(`{"usuario": "", "senha": ""}`)

	_, err := CreateLoginInput(data)

	assert.Error(t, err)
}

func TestCreateLoginInput_EmptyInput(t *testing.T) {
	data := []byte(``)

	_, err := CreateLoginInput(data)

	assert.Error(t, err)
}
