package mocks

import (
  "github.com/stretchr/testify/mock"
)

type ServiceLocatorMock struct {
    mock.Mock
}

func CreateServiceLocatorMock() *ServiceLocatorMock {
    return new(ServiceLocatorMock)
}

func (m *ServiceLocatorMock) Set(key string, value interface{}) {
   m.Called(key, value)
}

func (m *ServiceLocatorMock) Get(key string) interface{} {
   args := m.Called(key)
   return args.Get(0)
}
