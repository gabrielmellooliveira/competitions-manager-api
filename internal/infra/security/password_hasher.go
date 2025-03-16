package security

import (
	security "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/security"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewBcryptHasher() security.PasswordHasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (b *BcryptHasher) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
