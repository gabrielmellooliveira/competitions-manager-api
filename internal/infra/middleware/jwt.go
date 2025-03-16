package middleware

import (
	"errors"

	auth "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/auth"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type JwtMiddleware struct {
	Authenticator auth.Authenticator
}

func NewJwtMiddleware(authenticator auth.Authenticator) webserver.Middleware {
	return &JwtMiddleware{
		Authenticator: authenticator,
	}
}

func (m *JwtMiddleware) Execute(context webserver.Context) error {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		return errors.New("token não encontrado")
	}

	token = removeBearerFromToken(token)

	value, err := m.Authenticator.ValidateToken(token)
	if err != nil {
		return err
	}

	context.Set("username", value)

	return nil
}

func removeBearerFromToken(token string) string {
	return token[7:]
}
