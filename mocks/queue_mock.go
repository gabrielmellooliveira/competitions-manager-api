package mocks

import (
	queue "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/queue"
	"github.com/stretchr/testify/mock"
)

type MockQueue struct {
	mock.Mock
}

func (m *MockQueue) Publish(queueName string, message []byte) error {
	args := m.Called(queueName, message)
	return args.Error(0)
}

func (m *MockQueue) Consume(queueName string, handler queue.Handler) {
	args := m.Called(queueName)
	data := args.Get(0).([]byte)
	handler(data)
}
