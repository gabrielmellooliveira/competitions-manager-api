package usecase

import (
	"encoding/json"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/common"
)

type NotifySupportersInput struct {
	Type    string `json:"tipo" validate:"required"`
	Team    string `json:"time" validate:"required"`
	Score   string `json:"placar"`
	Message string `json:"mensagem" validate:"required"`
}

func CreateNotifySupportersInput(data []byte) (NotifySupportersInput, error) {
	var input NotifySupportersInput
	if err := json.Unmarshal(data, &input); err != nil {
		return NotifySupportersInput{}, err
	}

	if err := common.Validate(input); err != nil {
		return NotifySupportersInput{}, err
	}

	return input, nil
}
