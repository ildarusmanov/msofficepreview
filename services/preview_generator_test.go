package services

import (
	"github.com/ildarusmanov/msofficepreview/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
		filePath    = "dir/path/file.txt"
		fileName    = "file.txt"
		fileSize    = int64(3)
		fileOwnerId = "owner id"
		fileVersion = "ver1"
		tokenTtl    = time.Now().Unix()
	)

	wopiDiscovery := mocks.CreateWopiDiscoveryMock()
	wopiDiscovery.On("FindPreviewUrl", "internal-https", "txt").Return("urlsrc", nil)

	token := mocks.CreateTokenMock()
	token.On("GetValue").Return(accessToken)
	token.On("GetTtl").Return(tokenTtl)
	token.On("GetFilePath").Return(filePath)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Generate", fileName).Return(accessToken)
	provider.On("FindToken", accessToken).Return(token, true)

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
	assert.Equal(previewInfo.GetTokenTtl(), tokenTtl*1000)
}
