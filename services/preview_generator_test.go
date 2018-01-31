package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/ildarusmanov/msofficepreview/test/mocks"
)

var serverHost = "http://test.com"

func TestCreatePreviewGenerator(t *testing.T) {
  wopiDiscovery := mocks.CreateWopiDiscoveryMock()
  provider := mocks.CreateTokenProviderMock()
  storage := mocks.CreateStorageMock()

  generator := CreatePreviewGenerator(serverHost, wopiDiscovery, provider, storage)

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

  wopiDiscovery := mocks.CreateWopiDiscoveryMock()
  wopiDiscovery.On("FindPreviewUrl", "internal-https", "txt").Return("urlsrc", nil)

  provider := mocks.CreateTokenProviderMock()
  provider.On("Generate").Return(accessToken)

  fileInfo := mocks.CreateFileInfoMock()
  fileInfo.On("GetFileName").Return(fileName)
  fileInfo.On("GetSize").Return(fileSize)
  fileInfo.On("GetVersion").Return(fileVersion)
  fileInfo.On("GetOwnerId").Return(fileOwnerId)

  storage := mocks.CreateStorageMock()
  storage.On("GetFileInfo", fileName).Return(fileInfo, nil)

  generator := CreatePreviewGenerator(serverHost, wopiDiscovery, provider, storage)

  previewInfo, err := generator.GetPreviewLink(fileName)

  assert := assert.New(t)
  assert.Nil(err)
  assert.NotNil(previewInfo.GetSrc())
  assert.Equal(previewInfo.GetToken(), accessToken)
  assert.NotNil(previewInfo.GetTokenTtl())
}
