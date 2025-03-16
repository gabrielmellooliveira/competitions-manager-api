package usecase

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type ListCompetitionsUseCase struct {
	HttpAdapter interfaces.Http
}

func NewListCompetitionsUseCase(httpAdapter interfaces.Http) *ListCompetitionsUseCase {
	return &ListCompetitionsUseCase{
		HttpAdapter: httpAdapter,
	}
}

func (u *ListCompetitionsUseCase) Execute(context webserver.Context) (any, error) {
	competitions, err := u.HttpAdapter.Get("/competitions")
	if err != nil {
		return []CompetitionOutputDto{}, err
	}

	competitionsResult, err := CreateCompetitionsResult(competitions)
	if err != nil {
		return []CompetitionOutputDto{}, err
	}

	var competitionsOutput []CompetitionOutputDto
	for _, competitionResult := range competitionsResult {
		output, err := ConvertCompetitionResultToOutputDto(competitionResult)
		if err != nil {
			return nil, err
		}

		competitionsOutput = append(competitionsOutput, output)
	}

	return competitionsOutput, nil
}
