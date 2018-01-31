package mocks

import (
	"github.com/stretchr/testify/mock"
)

type FileInfoMock struct {
	mock.Mock
}

func CreateFileInfoMock() *FileInfoMock {
	return new(FileInfoMock)
}

func (m *FileInfoMock) GetFileName() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *FileInfoMock) GetVersion() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *FileInfoMock) GetSize() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *FileInfoMock) GetOwnerId() string {
	args := m.Called()
	return args.Get(0).(string)
}
