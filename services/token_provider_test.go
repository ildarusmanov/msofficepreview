package services

import (
  "time"
  "testing"
  "github.com/stretchr/testify/assert"
)

var tokenLifetime = int64(3600)

func TestCreateTokenProvider(t *testing.T) {
    provider := CreateTokenProvider(tokenLifetime)

    assert.NotNil(t, provider)
}

func TestGenerate(t *testing.T) {
    provider := CreateTokenProvider(tokenLifetime)

    newToken := provider.Generate()

    assert := assert.New(t)
    assert.NotNil(newToken)

    for i := 1; i <= 10; i++ {
        assert.NotEqual(newToken,provider.Generate())
    }
}

func TestValidate(t *testing.T) {
    provider := CreateTokenProvider(tokenLifetime)
    validToken := provider.Generate()
    invalidToken := validToken + "1"

    assert := assert.New(t)
    assert.True(provider.Validate(validToken))
    assert.False(provider.Validate(invalidToken))
}

func TestCleanUp(t *testing.T) {
    var (
        minTokenLifetime = int64(3)
    )

    provider := CreateTokenProvider(minTokenLifetime)
    expiredToken := provider.Generate()
    time.Sleep(time.Duration(minTokenLifetime+1) * time.Second)
    provider.CleanUp()
    newToken := provider.Generate()

    assert := assert.New(t)
    assert.False(provider.Validate(expiredToken))
    assert.True(provider.Validate(newToken))
}
