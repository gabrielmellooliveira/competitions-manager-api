package repository

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	database "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/database"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
)

type UserRepository struct {
	Database database.Database
}

func NewUserRepository(database database.Database) repository.UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (r *UserRepository) CreateUser(user entity.User) error {
	return r.Database.Create(user)
}

func (r *UserRepository) GetUserByName(userName string) (*entity.User, error) {
	user := &entity.User{}

	err := r.Database.First(user, "name", userName)
	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}
