package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var filePath = "file.txt"
var tokenLifetime = int64(3600)

func TestCreateMemoryTokenProvider(t *testing.T) {
	provider := CreateMemoryTokenProvider(tokenLifetime)

	assert.NotNil(t, provider)
}

func TestGenerate(t *testing.T) {
	provider := CreateMemoryTokenProvider(tokenLifetime)

	newToken := provider.Generate(filePath)

	assert := assert.New(t)
	assert.NotNil(newToken)

	for i := 1; i <= 10; i++ {
		assert.NotEqual(newToken, provider.Generate(filePath))
	}
}

func TestFindAndValidate(t *testing.T) {
	provider := CreateMemoryTokenProvider(tokenLifetime)
	validTokenValue := provider.Generate(filePath)
	invalidTokenValue := validTokenValue + "1"
    validToken, validTokenFound := provider.FindToken(validTokenValue)
    invalidToken, invalidTokenFound := provider.FindToken(invalidTokenValue)

	assert := assert.New(t)
    assert.True(validTokenFound)
	assert.True(provider.Validate(validToken))
    assert.Equal(validToken.GetValue(), validTokenValue)
	assert.Nil(invalidToken)
    assert.False(invalidTokenFound)
}

func TestCleanUp(t *testing.T) {
	var (
		minTokenLifetime = int64(3)
	)

	provider := CreateMemoryTokenProvider(minTokenLifetime)
	expiredTokenValue := provider.Generate(filePath)
	time.Sleep(time.Duration(minTokenLifetime+1) * time.Second)
	newTokenValue := provider.Generate(filePath)

	assert := assert.New(t)
    expiredToken, expiredTokenFound := provider.FindToken(expiredTokenValue)
    newToken, newTokenFound := provider.FindToken(newTokenValue)

    assert.Nil(expiredToken)
    assert.NotNil(newToken)
    assert.False(expiredTokenFound)
    assert.True(newTokenFound)
}
