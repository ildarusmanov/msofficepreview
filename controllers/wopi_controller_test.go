package controllers

import (
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/gin-gonic/gin"
)

func TestCreateWopiController(t *testing.T) {
  controller := CreateWopiController(wopiStorage, tokenProvider)

  assert.NotNil(t, controller)
}

func TestCheckFileInfo(t *testing.T) {
  controller := CreateWopiController(storage, tokenProvider)

  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  controller.CheckFileInfo(ctx)

  assert := assert.New(t)
  assert.Equal(w.Result().Status, 200)
}

func TestGetFile(t *testing.T) {
  controller := CreateWopiController(storage, tokenProvider)

  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  controller.GetFile(ctx)

  assert := assert.New(t)
  assert.Equal(w.Result().Status, 200)
}
