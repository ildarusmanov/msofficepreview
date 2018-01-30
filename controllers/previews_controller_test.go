package controllers

import (
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "github.com/gin-gonic/gin"
)

// PrviewsGeneratorMock
type previewsGeneratorMock struct {
  mock.Mock
}

// mock get preview link method
func (m *previewsGeneratorMock) GetPreviewLink(accessToken, fileId string) (string, error) {
  args := m.Called(accessToken, fileId)
  return args.Get(0).(string), args.Error(1)
}

func TestCreatePreviewsController(t *testing.T) {
  generator := new(previewsGeneratorMock)

  controller := CreatePreviewsController(generator)

  assert.NotNil(t, controller)
}

func TestCreateMethod(t *testing.T) {
  var (
    accessToken = "accessToken"
    fileId = "fileId"
    previewUrl = "previewUrl"
  )

  generator := new(previewsGeneratorMock)
  generator.On("GetPreviewLink", accessToken, fileId).Return(previewUrl, nil)

  controller := CreatePreviewsController(generator)
  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  controller.Create(ctx)

  assert := assert.New(t)
  assert.Equal(w.Result().Status, 200)
}


