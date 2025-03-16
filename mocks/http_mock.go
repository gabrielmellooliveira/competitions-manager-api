package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockHttp struct {
	mock.Mock
}

func (m *MockHttp) AddHeader(key string, value string) {
	m.Called(key, value)
}

func (m *MockHttp) Get(url string) ([]byte, error) {
	args := m.Called(url)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockHttp) Post(url string, body io.Reader) ([]byte, error) {
	args := m.Called(url, body)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockHttp) Put(url string, body io.Reader) ([]byte, error) {
	args := m.Called(url, body)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockHttp) Delete(url string) ([]byte, error) {
	args := m.Called(url)
	return args.Get(0).([]byte), args.Error(1)
}
