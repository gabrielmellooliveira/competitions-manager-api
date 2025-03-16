package entity

import (
	"github.com/google/uuid"
)

type Supporter struct {
	Id    uuid.UUID `gorm:"primaryKey"`
	Name  string    `json:"nome"`
	Email string    `gorm:"uniqueIndex" json:"email"`
	Team  string    `json:"time"`
}

func NewSupporter(name string, email string, team string) *Supporter {
	return &Supporter{
		Id:    uuid.New(),
		Name:  name,
		Email: email,
		Team:  team,
	}
}
