package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestValidate_Success(t *testing.T) {
	input := TestStruct{
		Name:  "Gabriel",
		Email: "gabriel@gmail.com",
	}

	err := Validate(input)

	assert.NoError(t, err)
}

func TestValidateName_Failure(t *testing.T) {
	input := TestStruct{
		Name:  "",
		Email: "gabriel@gmail.com",
	}

	err := Validate(input)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "campo 'Name' inválido required")
}

func TestValidateEmail_Failure(t *testing.T) {
	input := TestStruct{
		Name:  "Gabriel",
		Email: "gabriel",
	}

	err := Validate(input)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "campo 'Email' inválido email")
}
