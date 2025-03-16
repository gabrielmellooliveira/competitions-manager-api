package usecase

import (
	"errors"
	"io"

	auth "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/auth"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
	security "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/security"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type LoginUseCase struct {
	Repository     repository.UserRepository
	PasswordHasher security.PasswordHasher
	Authenticator  auth.Authenticator
}

func NewLoginUseCase(
	repository repository.UserRepository,
	passwordHasher security.PasswordHasher,
	authenticator auth.Authenticator,
) *LoginUseCase {
	return &LoginUseCase{
		Repository:     repository,
		PasswordHasher: passwordHasher,
		Authenticator:  authenticator,
	}
}

func (u *LoginUseCase) Execute(context webserver.Context) (any, error) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		return nil, err
	}

	input, err := CreateLoginInput(body)
	if err != nil {
		return nil, err
	}

	user, err := u.Repository.GetUserByName(input.Name)
	if err != nil {
		return nil, err
	}

	err = u.PasswordHasher.ComparePassword(user.Password, input.Password)
	if err != nil {
		return nil, errors.New("senha incorreta")
	}

	token, err := u.Authenticator.GenerateToken(user.Name)
	if err != nil {
		return nil, err
	}

	return LoginOutput{Token: token}, nil
}
