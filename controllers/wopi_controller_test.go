package controllers

import (
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/ildarusmanov/msofficepreview/test/mocks"
  "github.com/gin-gonic/gin"
)

func TestCreateWopiController(t *testing.T) {
  storage := mocks.CreateStorageMock()
  provider := mocks.CreateTokenProviderMock()

  controller := CreateWopiController(storage, provider)

  assert.NotNil(t, controller)
}

func TestCheckFileInfo(t *testing.T) {
  storage := mocks.CreateStorageMock()
  provider := mocks.CreateTokenProviderMock()

  controller := CreateWopiController(storage, provider)

  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  controller.CheckFileInfo(ctx)

  assert := assert.New(t)
  assert.Equal(w.Result().Status, 200)
}

func TestGetFile(t *testing.T) {
  storage := mocks.CreateStorageMock()
  provider := mocks.CreateTokenProviderMock()

  controller := CreateWopiController(storage, provider)

  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  controller.GetFile(ctx)

  assert := assert.New(t)
  assert.Equal(w.Result().Status, 200)
}
