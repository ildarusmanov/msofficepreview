package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/test/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePreviewsController(t *testing.T) {
	generator := mocks.CreatePreviewGeneratorMock()

	controller := CreatePreviewsController(generator)

	assert.NotNil(t, controller)
}

func TestCreateMethod(t *testing.T) {
	var (
		fileId          = "fileId"
		previewSrc      = "previewUrl"
		previewToken    = "token"
		previewTokenTtl = int64(10)
	)

	previewInfo := mocks.CreatePreviewInfoMock()
	previewInfo.On("GetSrc").Return(previewSrc)
	previewInfo.On("GetToken").Return(previewToken)
	previewInfo.On("GetTokenTtl").Return(previewTokenTtl)

	generator := mocks.CreatePreviewGeneratorMock()
	generator.On("GetPreviewLink", fileId).Return(previewInfo, nil)

	controller := CreatePreviewsController(generator)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Params = gin.Params{
		gin.Param{Key: "fileId", Value: fileId},
	}

	controller.Create(ctx)

	assert := assert.New(t)
	assert.Equal(w.Result().StatusCode, http.StatusOK)
}
