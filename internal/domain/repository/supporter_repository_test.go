package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	"github.com/gabrielmellooliveira/competitions-manager-api/mocks"
)

func TestCreateSupporter_Success(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewSupporterRepository(mockDB)

	supporter := entity.Supporter{
		Id:    uuid.New(),
		Name:  "Gabriel",
		Email: "gabriel@example.com",
		Team:  "Coritiba",
	}

	mockDB.On("Create", supporter).Return(nil)

	err := repo.CreateSupporter(supporter)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestCreateSupporter_Failure(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewSupporterRepository(mockDB)

	supporter := entity.Supporter{
		Id:    uuid.New(),
		Name:  "Gabriel",
		Email: "gabriel@example.com",
		Team:  "Coritiba",
	}

	mockDB.On("Create", supporter).Return(errors.New("database error"))

	err := repo.CreateSupporter(supporter)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockDB.AssertExpectations(t)
}

func TestGetSupportersByTeam_Failure(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	repo := NewSupporterRepository(mockDB)

	mockDB.On("GetClient").Return((*sql.DB)(nil), errors.New("failed to connect to database"))

	supporters, err := repo.GetSupportersByTeam("Coritiba")

	assert.Error(t, err)
	assert.Nil(t, supporters)
	assert.Equal(t, "failed to connect to database", err.Error())
	mockDB.AssertExpectations(t)
}
