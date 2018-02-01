package mocks

import (
    "github.com/stretchr/testify/mock"
)

type TokenMock struct {
    mock.Mock
}

func CreateTokenMock() *TokenMock {
    return new(TokenMock)
}

func (m *TokenMock) GetValue() string {
    args := m.Called()
    return args.Get(0).(string)
}

func (m *TokenMock) GetFilePath() string {
    args := m.Called()
    return args.Get(0).(string)
}

func (m *TokenMock) GetTtl() int64 {
    args := m.Called()
    return args.Get(0).(int64)
}
