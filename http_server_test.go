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
		fileContents    = []byte("content")
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

	storage := mocks.CreateStorageMock()
	storage.On("GetContents", fileId).Return(fileContents, nil)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Validate", accessToken).Return(true)

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
	req3, _ := http.NewRequest("POST", "/api/v1/previews/"+fileId, nil)

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, w1.Code)
	assert.Equal(http.StatusOK, w2.Code)
	assert.Equal(http.StatusOK, w3.Code)
}
