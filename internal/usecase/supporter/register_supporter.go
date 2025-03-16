package usecase

import (
	"io"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type RegisterSupporterUseCase struct {
	Repository repository.SupporterRepository
}

func NewRegisterSupporterUseCase(
	repository repository.SupporterRepository,
) *RegisterSupporterUseCase {
	return &RegisterSupporterUseCase{
		Repository: repository,
	}
}

func (u *RegisterSupporterUseCase) Execute(context webserver.Context) (any, error) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		return nil, err
	}

	input, err := CreateRegisterSupporterInput(body)
	if err != nil {
		return nil, err
	}

	supporter := entity.NewSupporter(input.Name, input.Email, input.Team)

	err = u.Repository.CreateSupporter(*supporter)
	if err != nil {
		return nil, err
	}

	return RegisterSupporterOutput{
		Id:      supporter.Id.String(),
		Name:    input.Name,
		Email:   input.Email,
		Team:    input.Team,
		Message: "Cadastro realizado com sucesso",
	}, nil
}
