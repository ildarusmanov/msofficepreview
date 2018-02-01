package main

import (
	"github.com/ildarusmanov/msofficepreview/test/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterEndpoints(t *testing.T) {
	var (
		fileId          = "fileId"
		fileName        = "file.xls"
		fileContents    = []byte("content")
		fileSize        = int64(len(fileContents))
		fileOwnerId     = "ownerId"
		fileVersion     = "123"
		previewSrc      = "previewUrl"
		previewToken    = "token"
		previewTokenTtl = int64(10)
		accessToken     = "accessToken"
	)

	previewInfo := mocks.CreatePreviewInfoMock()
	previewInfo.On("GetSrc").Return(previewSrc)
	previewInfo.On("GetToken").Return(previewToken)
	previewInfo.On("GetTokenTtl").Return(previewTokenTtl)

	generator := mocks.CreatePreviewGeneratorMock()
	generator.On("GetPreviewLink", fileId).Return(previewInfo, nil)

	fileInfo := mocks.CreateFileInfoMock()
	fileInfo.On("GetFileName").Return(fileName)
	fileInfo.On("GetVersion").Return(fileVersion)
	fileInfo.On("GetSize").Return(fileSize)
	fileInfo.On("GetOwnerId").Return(fileOwnerId)

	storage := mocks.CreateStorageMock()
	storage.On("GetContents", fileId).Return(fileContents, nil)
	storage.On("GetFileInfo", fileId).Return(fileInfo, nil)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Validate", accessToken).Return(true)

	serviceLocator := mocks.CreateServiceLocatorMock()
	serviceLocator.On("Get", "TokenProvider").Return(provider, nil)
	serviceLocator.On("Get", "Storage").Return(storage, nil)
	serviceLocator.On("Get", "PreviewGenerator").Return(generator, nil)

	router := CreateRouter(serviceLocator)

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/wopi/files/"+fileId+"?accessToken="+accessToken, nil)
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/wopi/files/"+fileId+"/contents?accessToken="+accessToken, nil)
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", "/api/v1/previews/"+fileId, nil)
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
