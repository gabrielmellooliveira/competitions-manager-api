package usecase

import (
	"errors"
	"io"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
	security "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/security"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type SignUpUseCase struct {
	Repository     repository.UserRepository
	PasswordHasher security.PasswordHasher
}

func NewSignUpUseCase(
	repository repository.UserRepository,
	passwordHasher security.PasswordHasher,
) *SignUpUseCase {
	return &SignUpUseCase{
		Repository:     repository,
		PasswordHasher: passwordHasher,
	}
}

func (u *SignUpUseCase) Execute(context webserver.Context) (any, error) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		return nil, err
	}

	input, err := CreateSignUpInput(body)
	if err != nil {
		return nil, err
	}

	if input.Password != input.PasswordConfirm {
		return nil, errors.New("a confirmação da senha está diferente da senha")
	}

	hashedPassword, err := u.PasswordHasher.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(input.User, hashedPassword)

	err = u.Repository.CreateUser(*user)
	if err != nil {
		return nil, err
	}

	return SignUpOutput{
		User:    input.User,
		Message: "Cadastro realizado com sucesso",
	}, nil
}
