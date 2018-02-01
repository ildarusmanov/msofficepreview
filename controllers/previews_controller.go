package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/interfaces"
	"github.com/ildarusmanov/msofficepreview/models"
	"net/http"
)

type PreviewsController struct {
	generator interfaces.PreviewGenerator
}

func CreatePreviewsController(generator interfaces.PreviewGenerator) *PreviewsController {
	return &PreviewsController{generator}
}

func (c *PreviewsController) Create(ctx *gin.Context) {
	var json models.File

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	previewInfo, err := c.generator.GetPreviewLink(json.FilePath)

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
