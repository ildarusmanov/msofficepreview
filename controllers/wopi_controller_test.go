package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/test/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
    "time"
)

func TestCreateWopiController(t *testing.T) {
	storage := mocks.CreateStorageMock()
	provider := mocks.CreateTokenProviderMock()

	controller := CreateWopiController(storage, provider)

	assert.NotNil(t, controller)
}

func TestCheckFileInfo(t *testing.T) {
	var (
		filePath     = "/dir/file/oath/file.txt"
		fileSize    = int64(100)
		fileVersion = "111"
		fileName    = "file.txt"
		fileOwnerId = "123"
        tokenValue  = "token-value"
        tokenTtl    = time.Now().Unix()
	)

    token := mocks.CreateTokenMock()
    token.On("GetValue").Return(tokenValue)
    token.On("GetTtl").Return(tokenTtl)
    token.On("GetFilePath").Return(filePath)

	fileInfo := mocks.CreateFileInfoMock()
	fileInfo.On("GetFileName").Return(fileName)
	fileInfo.On("GetVersion").Return(fileVersion)
	fileInfo.On("GetSize").Return(fileSize)
	fileInfo.On("GetOwnerId").Return(fileOwnerId)

	storage := mocks.CreateStorageMock()
	storage.On("GetFileInfo", filePath).Return(fileInfo, nil)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Validate", token).Return(true)
    provider.On("Generate", filePath).Return(tokenValue)
    provider.On("FindToken", tokenValue).Return(token, true)

	controller := CreateWopiController(storage, provider)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "fileId", Value: tokenValue}}

	controller.CheckFileInfo(ctx)

	assert := assert.New(t)
	assert.Equal(w.Result().StatusCode, http.StatusOK)
}

func TestGetFile(t *testing.T) {
	var (
		tokenValue   = "accessToken"
        tokenTtl     = time.Now().Unix()
        filePath     = "dir/file/path/file.txt"
        fileId       = tokenValue
		fileContents = []byte("contents")
	)

	storage := mocks.CreateStorageMock()
	storage.On("GetContents", filePath).Return(fileContents, nil)

    token := mocks.CreateTokenMock()
    token.On("GetValue").Return(tokenValue)
    token.On("GetTtl").Return(tokenTtl)
    token.On("GetFilePath").Return(filePath)

	provider := mocks.CreateTokenProviderMock()
	provider.On("Validate", token).Return(true)
    provider.On("FindToken", tokenValue).Return(token, true)

	controller := CreateWopiController(storage, provider)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "fileId", Value: fileId}}

	controller.GetFile(ctx)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert := assert.New(t)
	assert.Equal(resp.StatusCode, http.StatusOK)
	assert.Equal(body, fileContents)
}
