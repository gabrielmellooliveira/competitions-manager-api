package mocks

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) GetNotificationById(notificationId any) (*entity.Notification, error) {
	args := m.Called(notificationId)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Notification), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockNotificationRepository) UpdateNotification(notificationId any, notification entity.Notification) error {
	args := m.Called(notificationId, notification)
	return args.Error(0)
}

func (m *MockNotificationRepository) CreateNotification(notification entity.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}
