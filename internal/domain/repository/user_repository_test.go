package repository

import (
	"testing"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	mockDatabase := new(mocks.MockDatabase)
	userRepo := NewUserRepository(mockDatabase)

	user := entity.User{
		Id:   uuid.New(),
		Name: "Gabriel",
	}

	mockDatabase.On("Create", user).Return(nil)

	err := userRepo.CreateUser(user)
	assert.NoError(t, err)
	mockDatabase.AssertExpectations(t)
}

func TestCreateUser_Failure(t *testing.T) {
	mockDatabase := new(mocks.MockDatabase)
	userRepo := NewUserRepository(mockDatabase)

	user := entity.User{
		Id:   uuid.New(),
		Name: "Gabriel",
	}

	mockDatabase.On("Create", user).Return(assert.AnError)

	err := userRepo.CreateUser(user)
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockDatabase.AssertExpectations(t)
}

func TestGetUserByName_Success(t *testing.T) {
	mockDatabase := new(mocks.MockDatabase)
	userRepo := NewUserRepository(mockDatabase)

	userName := "Gabriel"
	user := &entity.User{
		Id:   uuid.New(),
		Name: userName,
	}

	mockDatabase.On("First", mock.AnythingOfType("*entity.User"), "name", userName).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.User)
		*arg = *user
	})

	result, err := userRepo.GetUserByName(userName)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockDatabase.AssertExpectations(t)
}

func TestGetUserByName_Failure(t *testing.T) {
	mockDatabase := new(mocks.MockDatabase)
	userRepo := NewUserRepository(mockDatabase)

	userName := "NonExistentUser"

	mockDatabase.On("First", mock.Anything, "name", userName).Return(assert.AnError)

	_, err := userRepo.GetUserByName(userName)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockDatabase.AssertExpectations(t)
}
