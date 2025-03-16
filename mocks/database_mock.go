package mocks

import (
	"database/sql"

	entity "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/database"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabase) MigrateEntity(entity entity.Entity) {
	m.Called(entity)
}

func (m *MockDatabase) Create(entity entity.Entity) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockDatabase) Update(entity entity.Entity, id any) error {
	args := m.Called(entity, id)
	return args.Error(0)
}

func (m *MockDatabase) First(entity entity.Entity, key string, value any) error {
	args := m.Called(entity, key, value)
	return args.Error(0)
}

func (m *MockDatabase) Find(entities any, key string, value any) error {
	args := m.Called(entities, key, value)
	return args.Error(0)
}

func (m *MockDatabase) GetClient() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}
