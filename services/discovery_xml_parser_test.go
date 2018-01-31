package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateDiscoveryXmlParser(t *testing.T) {
    parser := CreateDiscoveryXmlParser()

    assert.NotNil(t, parser)
}
