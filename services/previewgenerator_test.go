package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/ildarusmanov/msofficepreview/test/mocks"
)

const fileId = "url"

func TestCreatePreviewGenerator(t *testing.T) {
  provider := mocks.CreateTokenProviderMock()
  storage := mocks.CreateStorageMock()

  generator := CreatePreviewGenerator(provider, storage)

  assert.NotNil(t, generator)
}

func TestGetPreviewLink(t *testing.T) {
  var (
    accessToken = "access token"
    fileName = "file.txt"
    fileSize = int64(3)
    fileOwnerId = "owner id"
    fileVersion = "ver1"
  )

  provider := mocks.CreateTokenProviderMock()
  provider.On("Validate", accessToken).Return(true)

  fileInfo := mocks.CreateFileInfoMock()
  fileInfo.On("GetFileName").Return(fileName)
  fileInfo.On("GetSize").Return(fileSize)
  fileInfo.On("GetVersion").Return(fileVersion)
  fileInfo.On("GetOwnerId").Return(fileOwnerId)

  storage := mocks.CreateStorageMock()
  storage.On("GetFileInfo").Return(fileInfo, nil)

  generator := CreatePreviewGenerator(provider, storage)

  previewLink, err := generator.GetPreviewLink(accessToken, fileId)

  assert := assert.New(t)
  assert.Nil(err)
  assert.NotNil(previewLink)
}
