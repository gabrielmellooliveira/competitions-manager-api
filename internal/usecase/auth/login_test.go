package usecase

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	interfaces "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase_Execute_Success(t *testing.T) {
	mockContext := interfaces.Context{}
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockAuth := new(mocks.MockAuthenticator)

	loginUseCase := NewLoginUseCase(mockRepo, mockHasher, mockAuth)

	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewBufferString(`{"usuario": "gabriel", "senha": "gabriel12345"}`)),
	}
	mockRepo.On("GetUserByName", "gabriel").Return(&entity.User{Name: "gabriel", Password: "hashedpass"}, nil)
	mockHasher.On("ComparePassword", "hashedpass", "gabriel12345").Return(nil)
	mockAuth.On("GenerateToken", "gabriel").Return("token123", nil)

	result, err := loginUseCase.Execute(mockContext)

	assert.NoError(t, err)
	assert.Equal(t, LoginOutput{Token: "token123"}, result)
}

func TestLoginUseCase_Execute_NotFound(t *testing.T) {
	mockContext := interfaces.Context{}
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockAuth := new(mocks.MockAuthenticator)

	loginUseCase := NewLoginUseCase(mockRepo, mockHasher, mockAuth)

	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewBufferString(`{"name":"unknownuser","password":"gabriel12345"}`)),
	}
	mockRepo.On("GetUserByName", "unknownuser").Return(entity.User{}, errors.New("user not found"))

	result, err := loginUseCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_IncorrectPassword(t *testing.T) {
	mockContext := interfaces.Context{}
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockAuth := new(mocks.MockAuthenticator)

	loginUseCase := NewLoginUseCase(mockRepo, mockHasher, mockAuth)

	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewBufferString(`{"name":"gabriel","password":"wrongpass"}`)),
	}
	mockRepo.On("GetUserByName", "gabriel").Return(entity.User{Name: "gabriel", Password: "hashedpass"}, nil)
	mockHasher.On("ComparePassword", "hashedpass", "wrongpass").Return(errors.New("senha incorreta"))

	result, err := loginUseCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_TokenGenerationFailure(t *testing.T) {
	mockContext := interfaces.Context{}
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockAuth := new(mocks.MockAuthenticator)

	loginUseCase := NewLoginUseCase(mockRepo, mockHasher, mockAuth)

	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewBufferString(`{"name":"gabriel","password":"gabriel12345"}`)),
	}
	mockRepo.On("GetUserByName", "gabriel").Return(entity.User{Name: "gabriel", Password: "hashedpass"}, nil)
	mockHasher.On("ComparePassword", "hashedpass", "gabriel12345").Return(nil)
	mockAuth.On("GenerateToken", "gabriel").Return("", errors.New("token generation failed"))

	result, err := loginUseCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
}
