package mocks

import (
  "github.com/stretchr/testify/mock"
)

type PreviewsGeneratorMock struct {
  mock.Mock
}

func CreatePreviewGeneratorMock() *PreviewsGeneratorMock {
    return new(PreviewsGeneratorMock)
}

func (m *PreviewsGeneratorMock) GetPreviewLink(fileId string) (string, error) {
  args := m.Called(fileId)
  return args.Get(0).(string), args.Error(1)
}