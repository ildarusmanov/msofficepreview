package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/test/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWopiController(t *testing.T) {
	storage := mocks.CreateStorageMock()
	provider := mocks.CreateTokenProviderMock()

	controller := CreateWopiController(storage, provider)

	assert.NotNil(t, controller)
}

func TestCheckFileInfo(t *testing.T) {
	var (
		fileId      = "fileId"
		accessToken = "accessToken"
		fileSize    = int64(100)
		fileVersion = "111"
		fileName    = "file.txt"
		fileOwnerId = "123"
	)

	fileInfo := mocks.CreateFileInfoMock()
	fileInfo.On("GetFileName").Return(fileName)
	fileInfo.On("GetVersion").Return(fileVersion)
	fileInfo.On("GetSize").Return(fileSize)
	fileInfo.On("GetOwnerId").Return(fileOwnerId)

	storage := mocks.CreateStorageMock()
	storage.On("GetFileInfo", fileId).Return(fileInfo, nil)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Validate", accessToken).Return(true)

	controller := CreateWopiController(storage, provider)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Params = gin.Params{
		gin.Param{Key: "fileId", Value: fileId},
		gin.Param{Key: "accessToken", Value: accessToken},
	}

	controller.CheckFileInfo(ctx)

	assert := assert.New(t)
	assert.Equal(w.Result().StatusCode, http.StatusOK)
}

func TestGetFile(t *testing.T) {
	var (
		fileId       = "fileId"
		accessToken  = "accessToken"
		fileContents = []byte("contents")
	)

	storage := mocks.CreateStorageMock()
	storage.On("GetContents", fileId).Return(fileContents, nil)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Validate", accessToken).Return(true)

	controller := CreateWopiController(storage, provider)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Params = gin.Params{
		gin.Param{Key: "fileId", Value: fileId},
		gin.Param{Key: "accessToken", Value: accessToken},
	}

	controller.GetFile(ctx)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert := assert.New(t)
	assert.Equal(resp.StatusCode, http.StatusOK)
	assert.Equal(body, fileContents)
}
