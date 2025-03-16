package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertJsonToMatchesResult_Success(t *testing.T) {
	data := []byte(`{
		"matches": [
			{
				"matchDay": 1,
				"homeTeam": { "shortName": "Coritiba" },
				"awayTeam": { "shortName": "Parana" },
				"score": { "fullTime": { "home": 3, "away": 1 } }
			},
			{
				"matchDay": 2,
				"homeTeam": { "shortName": "Palmeiras" },
				"awayTeam": { "shortName": "Vasco" },
				"score": { "fullTime": { "home": 2, "away": 2 } }
			}
		]
	}`)

	matches, err := ConvertJsonToMatchesResult(data)

	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.Equal(t, int32(1), matches[0].MatchDay)
	assert.Equal(t, "Coritiba", matches[0].HomeTeam.Name)
	assert.Equal(t, 3, *matches[0].Score.FullTime.Home)
	assert.Equal(t, 1, *matches[0].Score.FullTime.Away)
}

func TestConvertJsonToMatchesResult_Error(t *testing.T) {
	data := []byte(`{
		"matches": [
			{
				"matchDay": 1,
				"homeTeam": { "shortName": "Coritiba" },
				"awayTeam": { "shortName": "Parana" },
				"score": { "fullTime": { "home": "invalid", "away": "invalid" } }
			}
		]
	}`)

	matches, err := ConvertJsonToMatchesResult(data)

	assert.NotNil(t, err)
	assert.Nil(t, matches)
}

func TestConvertMatchResultToOutputDto_Success(t *testing.T) {
	match := MatchResult{
		MatchDay: 1,
		HomeTeam: MatchTeamResult{Name: "Coritiba"},
		AwayTeam: MatchTeamResult{Name: "Parana"},
		Score: MatchScoreResult{
			FullTime: MatchScoreFullTimeResult{
				Home: new(int),
				Away: new(int),
			},
		},
	}
	*match.Score.FullTime.Home = 3
	*match.Score.FullTime.Away = 1

	outputDto := ConvertMatchResultToOutputDto(match)

	assert.Equal(t, "Coritiba", outputDto.HomeTeam)
	assert.Equal(t, "Parana", outputDto.AwayTeam)
	assert.Equal(t, "3-1", outputDto.Score)
}

func TestConvertMatchResultToOutputDto_NoScore(t *testing.T) {
	match := MatchResult{
		MatchDay: 1,
		HomeTeam: MatchTeamResult{Name: "Coritiba"},
		AwayTeam: MatchTeamResult{Name: "Parana"},
		Score: MatchScoreResult{
			FullTime: MatchScoreFullTimeResult{
				Home: nil,
				Away: nil,
			},
		},
	}

	outputDto := ConvertMatchResultToOutputDto(match)

	assert.Equal(t, "Coritiba", outputDto.HomeTeam)
	assert.Equal(t, "Parana", outputDto.AwayTeam)
	assert.Equal(t, "", outputDto.Score)
}
