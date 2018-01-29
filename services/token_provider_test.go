package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateTokenProvider(t *testing.T) {
    provider := CreateTokenProvider()

    assert.NotNil(t, provider)
}

func TestGenerate(t *testing.T) {
    provider := CreateTokenProvider()

    newToken := provider.Generate()

    assert := assert.New(t)
    assert.NotNil(newToken)

    for i := 1; i <= 10; i++ {
        assert.NotEqual(newToken,provider.Generate())
    }
}

func TestValidate(t *testing.T) {
    provider := CreateTokenProvider()
    validToken := provider.Generate()
    invalidToken := validToken + "1"

    assert.True(provider.Validate(validToken))
    assert.False(provider.Validate(invalidToken))
}
