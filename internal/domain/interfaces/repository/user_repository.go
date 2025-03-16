package interfaces

import "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByName(userName string) (*entity.User, error)
}
