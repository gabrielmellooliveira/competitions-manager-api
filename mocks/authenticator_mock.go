package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockAuthenticator struct {
	mock.Mock
}

func (m *MockAuthenticator) GenerateToken(value string) (string, error) {
	args := m.Called(value)
	return args.String(0), args.Error(1)
}

func (m *MockAuthenticator) ValidateToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}
