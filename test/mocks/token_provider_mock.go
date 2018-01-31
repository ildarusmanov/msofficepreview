package mocks

import (
	"github.com/stretchr/testify/mock"
)

type TokenProviderMock struct {
	mock.Mock
}

func CreateTokenProviderMock() *TokenProviderMock {
	return new(TokenProviderMock)
}

func (m *TokenProviderMock) Generate() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *TokenProviderMock) Validate(accessToken string) bool {
	args := m.Called(accessToken)
	return args.Get(0).(bool)
}
