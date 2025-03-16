package repository

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	database "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/database"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
)

type NotificationRepository struct {
	Database database.Database
}

func NewNotificationRepository(database database.Database) repository.NotificationRepository {
	return &NotificationRepository{
		Database: database,
	}
}

func (r *NotificationRepository) CreateNotification(notification entity.Notification) error {
	return r.Database.Create(notification)
}

func (r *NotificationRepository) UpdateNotification(notificationId any, notification entity.Notification) error {
	return r.Database.Update(notification, notificationId)
}

func (r *NotificationRepository) GetNotificationById(notificationId any) (*entity.Notification, error) {
	notification := &entity.Notification{}

	err := r.Database.First(notification, "id", notificationId)
	if err != nil {
		return &entity.Notification{}, err
	}

	return notification, nil
}
