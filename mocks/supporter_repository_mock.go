package mocks

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockSupporterRepository struct {
	mock.Mock
}

func (m *MockSupporterRepository) CreateSupporter(supporter entity.Supporter) error {
	args := m.Called(supporter)
	return args.Error(0)
}

func (m *MockSupporterRepository) GetSupportersByTeam(team string) ([]entity.Supporter, error) {
	args := m.Called(team)
	if args.Get(0) != nil {
		return args.Get(0).([]entity.Supporter), args.Error(1)
	}
	return nil, args.Error(1)
}
