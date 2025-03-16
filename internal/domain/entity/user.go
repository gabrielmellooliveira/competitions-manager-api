package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"uniqueIndex" json:"name"`
	Password string    `json:"password"`
}

func NewUser(name string, password string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		Password: password,
	}
}
