package mocks

import (
	"github.com/stretchr/testify/mock"
)

type WopiDiscoveryMock struct {
	mock.Mock
}

func CreateWopiDiscoveryMock() *WopiDiscoveryMock {
	return new(WopiDiscoveryMock)
}

func (m *WopiDiscoveryMock) FindPreviewUrl(zone, ext string) (string, error) {
	args := m.Called(zone, ext)
	return args.Get(0).(string), args.Error(1)
}
