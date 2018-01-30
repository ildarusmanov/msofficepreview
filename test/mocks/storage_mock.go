package mocks

import (
  "github.com/ildarusmanov/msofficepreview/interfaces"
  "github.com/stretchr/testify/mock"
)

type StorageMock struct {
  mock.Mock
}

func CreateStorageMock() *StorageMock {
    return new(StorageMock)
}

func (m *StorageMock) GetFileInfo(fileId string) (interfaces.FileInfo, error) {
  args := m.Called(fileId)
  return args.Get(0).(interfaces.FileInfo), args.Error(1)
}

func (m *StorageMock) GetContents(fileId string) ([]byte, error) {
  args := m.Called(fileId)
  return args.Get(0).([]byte), args.Error(1)
}
