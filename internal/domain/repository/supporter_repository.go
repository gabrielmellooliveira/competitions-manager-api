package repository

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	database "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/database"
	repository "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/repository"
)

type SupporterRepository struct {
	Database database.Database
}

func NewSupporterRepository(database database.Database) repository.SupporterRepository {
	return &SupporterRepository{
		Database: database,
	}
}

func (r *SupporterRepository) CreateSupporter(supporter entity.Supporter) error {
	return r.Database.Create(supporter)
}

func (r *SupporterRepository) GetSupportersByTeam(team string) ([]entity.Supporter, error) {
	db, err := r.Database.GetClient()
	if err != nil {
		return nil, err
	}

	query := `SELECT id, name, email, team FROM supporters WHERE team = $1`
	rows, err := db.Query(query, team)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var supporters []entity.Supporter

	for rows.Next() {
		var supporter entity.Supporter
		if err := rows.Scan(&supporter.Id, &supporter.Name, &supporter.Email, &supporter.Team); err != nil {
			return nil, err
		}
		supporters = append(supporters, supporter)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return supporters, nil
}
