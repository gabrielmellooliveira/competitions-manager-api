package interfaces

import "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"

type SupporterRepository interface {
	CreateSupporter(supporter entity.Supporter) error
	GetSupportersByTeam(team string) ([]entity.Supporter, error)
}
