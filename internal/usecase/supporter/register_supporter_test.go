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

func TestRegisterSupporterUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.MockSupporterRepository)

	mockRepo.On("CreateSupporter", mock.Anything).Return(nil)

	context := interfaces.Context{}
	inputJSON := `{"nome":"Gabriel","email":"gabriel@example.com","time":"Coritiba"}`
	context.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(inputJSON))),
	}

	useCase := NewRegisterSupporterUseCase(mockRepo)

	result, err := useCase.Execute(context)

	var output RegisterSupporterOutput = result.(RegisterSupporterOutput)

	assert.Nil(t, err)
	assert.Equal(t, RegisterSupporterOutput{
		Id:      output.Id,
		Name:    "Gabriel",
		Email:   "gabriel@example.com",
		Team:    "Coritiba",
		Message: "Cadastro realizado com sucesso",
	}, result)

	mockRepo.AssertExpectations(t)
}

func TestRegisterSupporterUseCase_Execute_ErrorCreatingSupporter(t *testing.T) {
	mockRepo := new(mocks.MockSupporterRepository)

	mockRepo.On("CreateSupporter", mock.Anything).Return(errors.New("repository error"))

	context := interfaces.Context{}
	inputJSON := `{"name":"Gabriel","email":"gabriel@example.com","team":"Coritiba"}`
	context.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(inputJSON))),
	}

	useCase := NewRegisterSupporterUseCase(mockRepo)

	result, err := useCase.Execute(context)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestRegisterSupporterUseCase_Execute_InvalidInput(t *testing.T) {
	context := interfaces.Context{}
	invalidBody := `{"name":"Gabriel","email":"invalid-email","team":"Coritiba"}`
	context.Request = http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(invalidBody))),
	}

	useCase := NewRegisterSupporterUseCase(nil)

	result, err := useCase.Execute(context)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
