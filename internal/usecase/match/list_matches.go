package usecase

import (
	"errors"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type ListMatchesUseCase struct {
	HttpAdapter interfaces.Http
}

func NewListMatchesUseCase(httpAdapter interfaces.Http) *ListMatchesUseCase {
	return &ListMatchesUseCase{
		HttpAdapter: httpAdapter,
	}
}

func (u *ListMatchesUseCase) Execute(context webserver.Context) (any, error) {
	competitionId := context.GetParam("competitionId")
	team := context.QueryParams.Get("equipe")
	matchday := context.QueryParams.Get("rodada")

	url, err := createFilterUrl(competitionId, matchday)
	if err != nil {
		return []MatchOutputDto{}, err
	}

	matches, err := u.HttpAdapter.Get(url)
	if err != nil {
		return []MatchOutputDto{}, err
	}

	matchesResult, err := ConvertJsonToMatchesResult(matches)
	if err != nil {
		return []MatchOutputDto{}, err
	}

	matchesDetailsOutput := filterMatchesByTeam(team, matchesResult)

	if matchday == "" {
		return MatchOutputDto{Matches: matchesDetailsOutput}, nil
	} else {
		return MatchWithMatchDayOutputDto{MatchDay: matchday, Matches: matchesDetailsOutput}, nil
	}
}

func createFilterUrl(competitionId string, matchday string) (string, error) {
	if competitionId == "" {
		return "", errors.New("identificador do campeonato não identificado")
	}

	url := "/competitions/" + competitionId + "/matches"

	if matchday == "" {
		return url, nil
	} else {
		return url + "?matchday=" + matchday, nil
	}
}

func filterMatchesByTeam(team string, matchesResult []MatchResult) []MatchDetailsOutputDto {
	var matchesDetailsOutput []MatchDetailsOutputDto

	for _, matchResult := range matchesResult {
		if team == "" || (matchResult.HomeTeam.Name == team || matchResult.AwayTeam.Name == team) {
			matchesDetailsOutput = append(
				matchesDetailsOutput,
				ConvertMatchResultToOutputDto(matchResult),
			)
		}
	}

	return matchesDetailsOutput
}
