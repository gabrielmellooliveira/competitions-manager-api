package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompetitionsResult_Success(t *testing.T) {
	data := []byte(`{
		"competitions": [
			{
				"id": 1,
				"name": "Competition 1",
				"currentSeason": {
					"startDate": "2025-01-01"
				}
			},
			{
				"id": 2,
				"name": "Competition 2",
				"currentSeason": {
					"startDate": "2024-01-01"
				}
			}
		]
	}`)

	competitions, err := CreateCompetitionsResult(data)

	assert.Nil(t, err)
	assert.Len(t, competitions, 2)
	assert.Equal(t, int32(1), competitions[0].Id)
	assert.Equal(t, "Competition 1", competitions[0].Name)
}

func TestCreateCompetitionsResult_Failure(t *testing.T) {
	data := []byte(`{
		"competitions: [
			{
				"id": 1,
				"name": "Competition 1",
				"currentSeason": {
					"startDate": "invalid-date"
				}
			}
		]
	}`)

	competitions, err := CreateCompetitionsResult(data)

	assert.NotNil(t, err)
	assert.Nil(t, competitions)
}

func TestConvertCompetitionResultToOutputDto_Success(t *testing.T) {
	competition := CompetitionResult{
		Id:   1,
		Name: "Competition 1",
		CurrentSeason: CompetitionCurrentSeasonResult{
			StartDate: "2025-01-01",
		},
	}

	outputDto, err := ConvertCompetitionResultToOutputDto(competition)

	assert.Nil(t, err)
	assert.Equal(t, int32(1), outputDto.Id)
	assert.Equal(t, "Competition 1", outputDto.Nome)
	assert.Equal(t, 2025, outputDto.Temporada)
}

func TestConvertCompetitionResultToOutputDto_Failure(t *testing.T) {
	competition := CompetitionResult{
		Id:   1,
		Name: "Competition 1",
		CurrentSeason: CompetitionCurrentSeasonResult{
			StartDate: "invalid-date",
		},
	}

	outputDto, err := ConvertCompetitionResultToOutputDto(competition)

	assert.NotNil(t, err)
	assert.Equal(t, CompetitionOutputDto{}, outputDto)
}
