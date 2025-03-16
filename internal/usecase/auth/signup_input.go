package usecase

import (
	"encoding/json"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/common"
)

type SignUpInput struct {
	User            string `json:"usuario" validate:"required"`
	Password        string `json:"senha" validate:"required"`
	PasswordConfirm string `json:"confirmarSenha" validate:"required"`
}

func CreateSignUpInput(data []byte) (SignUpInput, error) {
	var input SignUpInput
	if err := json.Unmarshal(data, &input); err != nil {
		return SignUpInput{}, err
	}

	if err := common.Validate(input); err != nil {
		return SignUpInput{}, err
	}

	if err := json.Unmarshal(data, &input); err != nil {
		return SignUpInput{}, err
	}

	return input, nil
}
