package usecase

import (
	"encoding/json"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/common"
)

type RegisterSupporterInput struct {
	Name  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Team  string `json:"time" validate:"required"`
}

func CreateRegisterSupporterInput(data []byte) (RegisterSupporterInput, error) {
	var input RegisterSupporterInput
	if err := json.Unmarshal(data, &input); err != nil {
		return RegisterSupporterInput{}, err
	}

	if err := common.Validate(input); err != nil {
		return RegisterSupporterInput{}, err
	}

	if err := json.Unmarshal(data, &input); err != nil {
		return RegisterSupporterInput{}, err
	}

	return input, nil
}
