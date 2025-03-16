package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	NOTIFICATION_TYPE_START string = "inicio"
	NOTIFICATION_TYPE_END   string = "fim"
)

type Notification struct {
	Id          uuid.UUID `gorm:"primaryKey"`
	SupporterId uuid.UUID `json:"supporterId"`
	Team        string    `json:"team"`
	Score       string    `json:"score"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewNotification(
	supporterId uuid.UUID,
	team string,
	score string,
	message string,
) *Notification {
	return &Notification{
		Id:          uuid.New(),
		SupporterId: supporterId,
		Team:        team,
		Score:       score,
		Message:     message,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
