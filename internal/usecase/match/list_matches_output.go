package usecase

import (
	"encoding/json"
	"strconv"
)

type ResultData struct {
	Matches []MatchResult `json:"matches"`
}

type MatchResult struct {
	MatchDay int32            `json:"matchDay"`
	HomeTeam MatchTeamResult  `json:"homeTeam"`
	AwayTeam MatchTeamResult  `json:"awayTeam"`
	Score    MatchScoreResult `json:"score"`
}

type MatchTeamResult struct {
	Name string `json:"shortName"`
}

type MatchScoreResult struct {
	FullTime MatchScoreFullTimeResult `json:"fullTime"`
}

type MatchScoreFullTimeResult struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type MatchOutputDto struct {
	Matches []MatchDetailsOutputDto `json:"partidas"`
}

type MatchWithMatchDayOutputDto struct {
	MatchDay string                  `json:"rodada"`
	Matches  []MatchDetailsOutputDto `json:"partidas"`
}

type MatchDetailsOutputDto struct {
	HomeTeam string `json:"time_casa"`
	AwayTeam string `json:"time_fora"`
	Score    string `json:"placar"`
}

func ConvertJsonToMatchesResult(data []byte) ([]MatchResult, error) {
	var result ResultData
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result.Matches, nil
}

func ConvertMatchResultToOutputDto(matchResult MatchResult) MatchDetailsOutputDto {
	return MatchDetailsOutputDto{
		HomeTeam: matchResult.HomeTeam.Name,
		AwayTeam: matchResult.AwayTeam.Name,
		Score:    buildScore(matchResult.Score.FullTime.Home, matchResult.Score.FullTime.Away),
	}
}

func buildScore(home *int, away *int) string {
	score := ""

	if home != nil && away != nil {
		score = strconv.Itoa(*home) + "-" + strconv.Itoa(*away)
	}

	return score
}
