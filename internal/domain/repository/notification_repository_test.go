package repository

import (
	"errors"
	"testing"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNotification_Success(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewNotificationRepository(mockDB)

	notification := entity.Notification{Id: uuid.New(), Message: "O jogo está prestes a começar!"}

	mockDB.On("Create", notification).Return(nil)

	err := repo.CreateNotification(notification)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestCreateNotification_Failure(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewNotificationRepository(mockDB)

	notification := entity.Notification{Id: uuid.New(), Message: "O jogo está prestes a começar!"}

	mockDB.On("Create", notification).Return(errors.New("failed to create"))

	err := repo.CreateNotification(notification)

	assert.Error(t, err)
	assert.Equal(t, "failed to create", err.Error())
	mockDB.AssertExpectations(t)
}

func TestUpdateNotification_Success(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewNotificationRepository(mockDB)

	notification := entity.Notification{Id: uuid.New(), Message: "O jogo está prestes a começar!"}

	mockDB.On("Update", notification, 1).Return(nil)

	err := repo.UpdateNotification(1, notification)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateNotification_Failure(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewNotificationRepository(mockDB)

	notification := entity.Notification{Id: uuid.New(), Message: "O jogo está prestes a começar!"}

	mockDB.On("Update", notification, 1).Return(errors.New("update failed"))

	err := repo.UpdateNotification(1, notification)

	assert.Error(t, err)
	assert.Equal(t, "update failed", err.Error())
	mockDB.AssertExpectations(t)
}

func TestGetNotificationById_Success(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewNotificationRepository(mockDB)

	identifier := uuid.New()

	expectedNotification := entity.Notification{Id: identifier, Message: "O jogo está prestes a começar!"}

	mockDB.On("First", mock.Anything, "id", identifier).Run(func(args mock.Arguments) {
		notification := args.Get(0).(*entity.Notification)
		*notification = expectedNotification
	}).Return(nil)

	result, err := repo.GetNotificationById(identifier)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedNotification.Message, result.Message)
	mockDB.AssertExpectations(t)
}

func TestGetNotificationById_NotFound(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewNotificationRepository(mockDB)

	mockDB.On("First", mock.Anything, "id", 2).Return(errors.New("not found"))

	result, err := repo.GetNotificationById(2)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	assert.NotNil(t, result)
	mockDB.AssertExpectations(t)
}
