package mocks

import (
  "github.com/stretchr/testify/mock"
)

type PreviewInfoMock struct {
    mock.Mock
}

func CreatePreviewInfoMock() *PreviewInfoMock {
    return new(PreviewInfoMock)
}

func (m *PreviewInfoMock) GetSrc() string {
   args := m.Called()
   return args.Get(0).(string)
}

func (m *PreviewInfoMock) GetToken() string {
   args := m.Called()
   return args.Get(0).(string)
}

func (m *PreviewInfoMock) GetTokenTtl() int64 {
   args := m.Called()
   return args.Get(0).(int64)
}
