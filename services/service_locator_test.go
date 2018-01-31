package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateServiceLocator(t *testing.T) {
	locator := CreateServiceLocator()

	assert.NotNil(t, locator)
}

func TestServiceLocatorSetGet(t *testing.T) {
	var (
		key   = "str-key"
		value = "key-value"
	)

	locator := CreateServiceLocator()
	locator.Set(key, value)

	assert := assert.New(t)
	assert.NotNil(locator)
	assert.Equal(value, locator.Get(key).(string))
}
