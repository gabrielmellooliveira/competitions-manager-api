package usecase_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	usecase "github.com/gabrielmellooliveira/competitions-manager-api/internal/usecase/broadcast"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestSendEmailSupportersUseCase_Execute_Success(t *testing.T) {
	mockQueue := new(mocks.MockQueue)
	mockRepo := new(mocks.MockNotificationRepository)
	useCase := usecase.NewSendEmailSupportersUseCase(mockRepo, mockQueue)

	notificationID := uuid.New()
	notification := &entity.Notification{Id: notificationID, Message: "O jogo iniciou", UpdatedAt: time.Time{}}

	mockQueue.On("Consume", "matches").Return(mockNotificationJSON(*notification))
	mockRepo.On("GetNotificationById", notificationID.String()).Return(notification, nil)
	mockRepo.On("UpdateNotification", notificationID.String(), mock.AnythingOfType("entity.Notification")).Return(nil)

	useCase.Execute()

	mockQueue.AssertCalled(t, "Consume", "matches")
	mockRepo.AssertCalled(t, "GetNotificationById", notificationID.String())
	mockRepo.AssertCalled(t, "UpdateNotification", notificationID.String(), mock.AnythingOfType("entity.Notification"))
}

func TestSendEmailSupportersUseCase_Execute_ErrorUpdatingNotification(t *testing.T) {
	mockQueue := new(mocks.MockQueue)
	mockRepo := new(mocks.MockNotificationRepository)
	useCase := usecase.NewSendEmailSupportersUseCase(mockRepo, mockQueue)

	notificationID := uuid.New()
	notification := &entity.Notification{Id: notificationID, Message: "O jogo iniciou", UpdatedAt: time.Time{}}

	mockQueue.On("Consume", "matches").Return(mockNotificationJSON(*notification))
	mockRepo.On("GetNotificationById", notificationID.String()).Return(notification, nil)
	mockRepo.On("UpdateNotification", notificationID.String(), mock.AnythingOfType("entity.Notification")).Return(errors.New("update error"))

	useCase.Execute()

	mockQueue.AssertCalled(t, "Consume", "matches")
	mockRepo.AssertCalled(t, "GetNotificationById", notificationID.String())
	mockRepo.AssertCalled(t, "UpdateNotification", notificationID.String(), mock.AnythingOfType("entity.Notification"))
}

func mockNotificationJSON(notification entity.Notification) []byte {
	byteData, err := json.Marshal(notification)
	if err != nil {
		return nil
	}

	return byteData
}
