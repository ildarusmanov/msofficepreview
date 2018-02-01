package main

import (
	"github.com/ildarusmanov/msofficepreview/test/mocks"
    "github.com/ildarusmanov/msofficepreview/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
    "bytes"
    "encoding/json"
    "time"
)

func TestRouterEndpoints(t *testing.T) {
	var (
		filePath        = "dir/file/path/file.xls"
		fileName        = "file.xls"
		fileContents    = []byte("content")
		fileSize        = int64(len(fileContents))
		fileOwnerId     = "ownerId"
		fileVersion     = "123"
		previewSrc      = "previewUrl"
		tokenValue      = "accessToken"
        tokenTtl        = time.Now().Unix()
        fileId          = tokenValue
	)

	previewInfo := mocks.CreatePreviewInfoMock()
	previewInfo.On("GetSrc").Return(previewSrc)
	previewInfo.On("GetToken").Return(tokenValue)
	previewInfo.On("GetTokenTtl").Return(tokenTtl)

	generator := mocks.CreatePreviewGeneratorMock()
	generator.On("GetPreviewLink", filePath).Return(previewInfo, nil)

	fileInfo := mocks.CreateFileInfoMock()
	fileInfo.On("GetFileName").Return(fileName)
	fileInfo.On("GetVersion").Return(fileVersion)
	fileInfo.On("GetSize").Return(fileSize)
	fileInfo.On("GetOwnerId").Return(fileOwnerId)

	storage := mocks.CreateStorageMock()
	storage.On("GetContents", filePath).Return(fileContents, nil)
	storage.On("GetFileInfo", filePath).Return(fileInfo, nil)

    token := mocks.CreateTokenMock()
    token.On("GetValue").Return(tokenValue)
    token.On("GetTtl").Return(tokenTtl)
    token.On("GetFilePath").Return(filePath)

	provider := mocks.CreateTokenProviderMock()
    provider.On("FindToken", tokenValue).Return(token, true)
	provider.On("Validate", token).Return(true)

	serviceLocator := mocks.CreateServiceLocatorMock()
	serviceLocator.On("Get", "TokenProvider").Return(provider, nil)
	serviceLocator.On("Get", "Storage").Return(storage, nil)
	serviceLocator.On("Get", "PreviewGenerator").Return(generator, nil)

	router := CreateRouter(serviceLocator)

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/wopi/files/"+fileId, nil)
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/wopi/files/"+fileId+"/contents", nil)
	w3 := httptest.NewRecorder()

    file := models.CreateFile(filePath)
    jsonValue, _ := json.Marshal(file)
	req3, _ := http.NewRequest("POST", "/api/v1/previews", bytes.NewBuffer(jsonValue))

    w4 := httptest.NewRecorder()
    req4, _ := http.NewRequest("GET", "/api/v1/status/check", nil)

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
    router.ServeHTTP(w4, req4)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, w1.Code)
	assert.Equal(http.StatusOK, w2.Code)
	assert.Equal(http.StatusOK, w3.Code)
    assert.Equal(http.StatusOK, w4.Code)
}
