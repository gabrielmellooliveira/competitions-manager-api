package usecase

import (
	"encoding/json"

	"github.com/google/uuid"
)

type NotifySupportersMessage struct {
	Id          uuid.UUID `json:"id"`
	SupporterId uuid.UUID `json:"supporterId"`
	Team        string    `json:"team"`
	Score       string    `json:"score"`
	Message     string    `json:"message"`
}

func (m *NotifySupportersMessage) ConvertToByte() ([]byte, error) {
	byteData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}
