package usecase

import (
	"encoding/json"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/common"
)

type LoginInput struct {
	Name     string `json:"usuario" validate:"required"`
	Password string `json:"senha" validate:"required"`
}

func CreateLoginInput(data []byte) (LoginInput, error) {
	var input LoginInput
	if err := json.Unmarshal(data, &input); err != nil {
		return LoginInput{}, err
	}

	if err := common.Validate(input); err != nil {
		return LoginInput{}, err
	}

	if err := json.Unmarshal(data, &input); err != nil {
		return LoginInput{}, err
	}

	return input, nil
}
