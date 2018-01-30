package mocks

import (
  "github.com/stretchr/testify/mock"
)

type StorageMock struct {
  mock.Mock
}

func CreateStorageMock() *storageMock {
    return new(StorageMock)
}

func (m *StorageMock) GetFileInfo(fileId string) (FileInfoMock, error) {
  args := m.Called(fileId)
  return args.Get(0).(FileInfoMock), args.Error(1)
}

func (m *StorageMock) GetContents(fileId string) ([]byte, error) {
  args := m.Called(fileId)
  return args.Get(0).([]byte), args.Error(1)
}
