package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO add tests
func TestBuildServiceLocator(t *testing.T) {
	config := CreateNewConfig()
	locator, err := BuildServiceLocator(config)

	assert.Nil(t, err)
	assert.NotNil(t, locator)
}
