package usecase

import (
	"errors"
	"io"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	queue "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/queue"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type NotifySupportersUseCase struct {
	Repository             repository.SupporterRepository
	NotificationRepository repository.NotificationRepository
	Queue                  queue.Queue
}

func NewNotifySupportersUseCase(
	repository repository.SupporterRepository,
	notificationRepository repository.NotificationRepository,
	queue queue.Queue,
) *NotifySupportersUseCase {
	return &NotifySupportersUseCase{
		Repository:             repository,
		NotificationRepository: notificationRepository,
		Queue:                  queue,
	}
}

func (u *NotifySupportersUseCase) Execute(context webserver.Context) (any, error) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		return nil, err
	}

	input, err := CreateNotifySupportersInput(body)
	if err != nil {
		return nil, err
	}

	if input.Type == entity.NOTIFICATION_TYPE_END && input.Score == "" {
		return nil, errors.New("quando for o fim da partida, precisa ter um placar")
	}

	supporters, err := u.Repository.GetSupportersByTeam(input.Team)
	if err != nil {
		return nil, err
	}

	const QUEUE_NAME = "matches"

	for _, supporter := range supporters {
		notification := entity.NewNotification(supporter.Id, supporter.Team, input.Score, input.Message)

		err := u.NotificationRepository.CreateNotification(*notification)
		if err != nil {
			return nil, err
		}

		message := &NotifySupportersMessage{
			Id:          notification.Id,
			SupporterId: supporter.Id,
			Team:        supporter.Team,
			Score:       input.Score,
			Message:     input.Message,
		}

		byteData, err := message.ConvertToByte()
		if err != nil {
			return nil, err
		}

		u.Queue.Publish(QUEUE_NAME, byteData)
	}

	return NotifySupportersOutput{
		Message: "Notificações para os torcedores enviadas com sucesso",
	}, nil
}
