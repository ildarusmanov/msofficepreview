package mocks

import (
    "github.com/ildarusmanov/msofficepreview/interfaces"
	"github.com/stretchr/testify/mock"
)

type TokenProviderMock struct {
	mock.Mock
}

func CreateTokenProviderMock() *TokenProviderMock {
	return new(TokenProviderMock)
}

func (m *TokenProviderMock) Generate(filePath string) string {
	args := m.Called(filePath)
	return args.Get(0).(string)
}

func (m *TokenProviderMock) Validate(token interfaces.Token) bool {
	args := m.Called(token)
	return args.Get(0).(bool)
}

func (m *TokenProviderMock) FindToken(accessToken string) (interfaces.Token, bool) {
    args := m.Called(accessToken)
    return args.Get(0).(interfaces.Token), args.Get(1).(bool)
}
