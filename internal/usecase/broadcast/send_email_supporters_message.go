package usecase

import (
	"encoding/json"
)

type Notification struct {
	Id          string `json:"id"`
	SupporterId string `json:"supporterId"`
	Team        string `json:"team"`
	Score       string `json:"score"`
	Message     string `json:"message"`
}

func CreateNotification(data []byte) (*Notification, error) {
	var notification Notification
	err := json.Unmarshal(data, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
