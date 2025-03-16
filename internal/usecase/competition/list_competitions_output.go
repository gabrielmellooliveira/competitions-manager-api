package usecase

import (
	"encoding/json"
	"time"
)

type ResultData struct {
	Competitions []CompetitionResult `json:"competitions"`
}

type CompetitionResult struct {
	Id            int32                          `json:"id"`
	Name          string                         `json:"name"`
	CurrentSeason CompetitionCurrentSeasonResult `json:"currentSeason"`
}

type CompetitionCurrentSeasonResult struct {
	StartDate string `json:"startDate"`
}

type CompetitionOutputDto struct {
	Id        int32  `json:"id"`
	Nome      string `json:"nome"`
	Temporada int    `json:"temporada"`
}

func CreateCompetitionsResult(data []byte) ([]CompetitionResult, error) {
	var result ResultData
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result.Competitions, nil
}

func ConvertCompetitionResultToOutputDto(competitionResult CompetitionResult) (CompetitionOutputDto, error) {
	startDate, err := time.Parse("2006-01-02", competitionResult.CurrentSeason.StartDate)
	if err != nil {
		return CompetitionOutputDto{}, err
	}

	return CompetitionOutputDto{
		Id:        competitionResult.Id,
		Nome:      competitionResult.Name,
		Temporada: startDate.Year(),
	}, nil
}
