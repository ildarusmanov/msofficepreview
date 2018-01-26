package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

const accessToken = "example-token"

func TestCreatePreviewGenerator(t *testing.T) {
  generator := CreatePreviewGenerator(accessToken)

  assert.NotNil(t, generator)
  assert.Equal(t,  generator.GetAccessToken(), accessToken)
}

