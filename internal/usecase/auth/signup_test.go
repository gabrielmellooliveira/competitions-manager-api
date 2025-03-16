package usecase

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	interfaces "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUpUseCase_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockContext := interfaces.Context{}

	inputJSON := `{"usuario": "gabriel", "senha": "password123", "confirmarSenha": "password123"}`
	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(inputJSON))),
	}

	mockHasher.On("HashPassword", "password123").Return("hashedpassword", nil)
	mockRepo.On("CreateUser", mock.AnythingOfType("entity.User")).Return(nil)

	useCase := NewSignUpUseCase(mockRepo, mockHasher)
	result, err := useCase.Execute(mockContext)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "gabriel", result.(SignUpOutput).User)
	assert.Equal(t, "Cadastro realizado com sucesso", result.(SignUpOutput).Message)
	mockRepo.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
}

func TestSignUpUseCase_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockContext := interfaces.Context{}

	invalidJSON := `{"usuario": "", "senha": "password", "confirmarSenha": "password"}`
	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(invalidJSON))),
	}

	useCase := NewSignUpUseCase(mockRepo, mockHasher)
	result, err := useCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSignUpUseCase_PasswordMismatch(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockContext := interfaces.Context{}

	inputJSON := `{"usuario": "gabriel", "senha": "password123", "confirmarSenha": "password456"}`
	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(inputJSON))),
	}

	useCase := NewSignUpUseCase(mockRepo, mockHasher)
	result, err := useCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "a confirmação da senha está diferente da senha", err.Error())
}

func TestSignUpUseCase_HashPasswordError(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockContext := interfaces.Context{}

	inputJSON := `{"usuario": "gabriel", "senha": "password123", "confirmarSenha": "password123"}`
	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(inputJSON))),
	}

	mockHasher.On("HashPassword", "password123").Return("", errors.New("erro ao hashear senha"))

	useCase := NewSignUpUseCase(mockRepo, mockHasher)
	result, err := useCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "erro ao hashear senha", err.Error())
	mockHasher.AssertExpectations(t)
}

func TestSignUpUseCase_CreateUserError(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockHasher := new(mocks.MockPasswordHasher)
	mockContext := interfaces.Context{}

	inputJSON := `{"usuario": "gabriel", "senha": "password123", "confirmarSenha": "password123"}`
	mockContext.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(inputJSON))),
	}

	mockHasher.On("HashPassword", "password123").Return("hashedpassword", nil)
	mockRepo.On("CreateUser", mock.AnythingOfType("entity.User")).Return(errors.New("erro ao criar usuário"))

	useCase := NewSignUpUseCase(mockRepo, mockHasher)
	result, err := useCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "erro ao criar usuário", err.Error())
	mockRepo.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
}
