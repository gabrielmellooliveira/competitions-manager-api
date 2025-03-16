package usecase

import (
	"errors"
	"testing"

	interfaces "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListCompetitionsUseCase_Success(t *testing.T) {
	mockHttpAdapter := new(mocks.MockHttp)
	mockContext := interfaces.Context{}

	mockHttpAdapter.On("Get", "/competitions").Return(mockCompetitionsJSON(), nil)

	useCase := NewListCompetitionsUseCase(mockHttpAdapter)

	result, err := useCase.Execute(mockContext)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "Campeonato Brasileiro Série A", result.([]CompetitionOutputDto)[0].Nome)
	mockHttpAdapter.AssertExpectations(t)
}

func TestListCompetitionsUseCase_Failure(t *testing.T) {
	mockHttpAdapter := new(mocks.MockHttp)
	mockContext := interfaces.Context{}

	mockHttpAdapter.On("Get", "/competitions").Return(mockCompetitionsJSON(), errors.New("erro ao buscar competições"))

	useCase := NewListCompetitionsUseCase(mockHttpAdapter)
	result, err := useCase.Execute(mockContext)

	assert.Error(t, err)
	assert.Equal(t, result, []CompetitionOutputDto{})
	assert.Equal(t, "erro ao buscar competições", err.Error())
}

func mockCompetitionsJSON() []byte {
	return []byte(`{
    "competitions": [
			{
				"id": 2013,
				"name": "Campeonato Brasileiro Série A",
				"currentSeason": {
					"startDate": "2021-05-30"
				}
			},
			{
				"id": 2000,
				"name": "FIFA World Cup",
				"currentSeason": {
					"startDate": "2021-05-30"
				}
			}
    ]
	}`)
}
