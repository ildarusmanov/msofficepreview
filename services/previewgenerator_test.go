package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

const accessToken = "example-token"
const exampleUrl = "url"

func TestCreatePreviewGenerator(t *testing.T) {
  generator := CreatePreviewGenerator(accessToken)

  assert.NotNil(t, generator)
  assert.Equal(t,  generator.GetAccessToken(), accessToken)
}

func TestGetPreviewLink(t *testing.T) {
  generator := CreatePreviewGenerator(accessToken)

  previewLink := generator.GetPreviewLink(exampleUrl)

  assert.NotNil(t, previewLink)
}

