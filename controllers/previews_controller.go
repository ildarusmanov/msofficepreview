package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/interfaces"
	"net/http"
)

type PreviewsController struct {
	generator interfaces.PreviewGenerator
}

func CreatePreviewsController(generator interfaces.PreviewGenerator) *PreviewsController {
	return &PreviewsController{generator}
}

func (c *PreviewsController) Create(ctx *gin.Context) {
	fileId := ctx.Params.ByName("fileId")
	previewInfo, err := c.generator.GetPreviewLink(fileId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Src":      previewInfo.GetSrc(),
		"Token":    previewInfo.GetToken(),
		"TokenTtl": previewInfo.GetTokenTtl(),
	})
}
