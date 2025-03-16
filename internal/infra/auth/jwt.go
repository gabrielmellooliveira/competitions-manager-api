package auth

import (
	"errors"
	"time"

	interfaces "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/auth"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Value string `json:"value"`
	jwt.RegisteredClaims
}

type JwtAuthenticator struct {
	SecretKey string
}

func NewJwtAuthenticator(
	secretKey string,
) interfaces.Authenticator {
	return &JwtAuthenticator{
		SecretKey: secretKey,
	}
}

func (j *JwtAuthenticator) GenerateToken(value string) (string, error) {
	claims := JwtClaims{
		Value: value,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JwtAuthenticator) ValidateToken(token string) (string, error) {
	tokenResult, err := jwt.ParseWithClaims(
		token,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil || !tokenResult.Valid {
		return "", errors.New("token inválido")
	}

	claims, ok := tokenResult.Claims.(*JwtClaims)
	if !ok {
		return "", errors.New("token inválido")
	}

	return claims.Value, nil
}
