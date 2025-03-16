package usecase

import (
	"time"

	queue "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/queue"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
)

type SendEmailSupportersUseCase struct {
	NotificationRepository repository.NotificationRepository
	Queue                  queue.Queue
}

func NewSendEmailSupportersUseCase(
	notificationRepository repository.NotificationRepository,
	queue queue.Queue,
) *SendEmailSupportersUseCase {
	return &SendEmailSupportersUseCase{
		NotificationRepository: notificationRepository,
		Queue:                  queue,
	}
}

func (u *SendEmailSupportersUseCase) Execute() {
	const QUEUE_NAME = "matches"

	u.Queue.Consume(QUEUE_NAME, func(data []byte) error {
		notificationData, err := CreateNotification(data)
		if err != nil {
			return err
		}

		// Aqui ficaria o código responsável por enviar um e-mail para o torcedor
		// Normalmente algum serviço externo responsável por isso

		err = u.updateNotification(notificationData.Id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *SendEmailSupportersUseCase) updateNotification(notificationId string) error {
	notification, err := u.NotificationRepository.GetNotificationById(notificationId)
	if err != nil {
		return err
	}

	notification.UpdatedAt = time.Now()

	err = u.NotificationRepository.UpdateNotification(notificationId, *notification)
	if err != nil {
		return err
	}

	return nil
}
