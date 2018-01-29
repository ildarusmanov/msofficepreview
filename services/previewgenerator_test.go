package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

const fileId = "url"

func TestCreatePreviewGenerator(t *testing.T) {
  generator := CreatePreviewGenerator(tokenProvider, storage)

  assert.NotNil(t, generator)
}

func TestGetPreviewLink(t *testing.T) {
  accessToken := tokenProvider.Generate()

  generator := CreatePreviewGenerator(tokenProvider, storage)

  previewLink := generator.GetPreviewLink(accessToken, fileId)

  assert.NotNil(t, previewLink)
}
