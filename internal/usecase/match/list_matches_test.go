package usecase

import (
	"errors"
	"net/url"
	"testing"

	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListMatchesUseCase_Execute_Success(t *testing.T) {
	mockHttpAdapter := new(mocks.MockHttp)

	mockHttpAdapter.On("Get", "/competitions/123/matches").Return([]byte(`{
		"matches": [
			{
				"matchDay": 1,
				"homeTeam": { "shortName": "Coritiba" },
				"awayTeam": { "shortName": "Parana" },
				"score": { "fullTime": { "home": 3, "away": 1 } }
			}
		]
	}`), nil)

	useCase := NewListMatchesUseCase(mockHttpAdapter)

	context := webserver.Context{}
	context.QueryParams = url.Values{}
	context.GetParam = func(key string) string {
		return "123"
	}

	result, err := useCase.Execute(context)

	assert.Nil(t, err)
	assert.IsType(t, MatchOutputDto{}, result)
	mockHttpAdapter.AssertExpectations(t)
}

func TestListMatchesUseCase_Execute_Failure(t *testing.T) {
	mockHttpAdapter := new(mocks.MockHttp)

	mockHttpAdapter.On("Get", "/competitions/123/matches").Return([]byte(`{}`), errors.New("HTTP error"))

	useCase := NewListMatchesUseCase(mockHttpAdapter)

	context := webserver.Context{}
	context.QueryParams = url.Values{}
	context.GetParam = func(key string) string {
		return "123"
	}

	result, err := useCase.Execute(context)

	assert.NotNil(t, err)
	assert.Equal(t, []MatchOutputDto{}, result)
}
