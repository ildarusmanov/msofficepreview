package mocks

import (
  "github.com/ildarusmanov/msofficepreview/interfaces"
  "github.com/stretchr/testify/mock"
)

type PreviewsGeneratorMock struct {
  mock.Mock
}

func CreatePreviewGeneratorMock() *PreviewsGeneratorMock {
    return new(PreviewsGeneratorMock)
}

func (m *PreviewsGeneratorMock) GetPreviewLink(fileId string) (interfaces.PreviewInfo, error) {
  args := m.Called(fileId)
  return args.Get(0).(interfaces.PreviewInfo), args.Error(1)
}