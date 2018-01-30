package controllers

import (
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/ildarusmanov/msofficepreview/test/mocks"
  "github.com/gin-gonic/gin"
)

func TestCreatePreviewsController(t *testing.T) {
  generator := mocks.CreatePreviewGeneratorMock()

  controller := CreatePreviewsController(generator)

  assert.NotNil(t, controller)
}

func TestCreateMethod(t *testing.T) {
  var (
    accessToken = "accessToken"
    fileId = "fileId"
    previewUrl = "previewUrl"
  )

  generator := mocks.CreatePreviewGeneratorMock()
  generator.On("GetPreviewLink", accessToken, fileId).Return(previewUrl, nil)

  controller := CreatePreviewsController(generator)
  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  ctx.Params = gin.Params{
    gin.Param{Key: "fileId", Value: fileId},
    gin.Param{Key: "accessToken", Value: accessToken},
  }

  controller.Create(ctx)

  assert := assert.New(t)
  assert.Equal(w.Result().Status, 200)
}


