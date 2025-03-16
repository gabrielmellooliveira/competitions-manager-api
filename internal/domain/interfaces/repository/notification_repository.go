package interfaces

import "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"

type NotificationRepository interface {
	CreateNotification(notification entity.Notification) error
	UpdateNotification(notificationId any, notification entity.Notification) error
	GetNotificationById(notificationId any) (*entity.Notification, error)
}
