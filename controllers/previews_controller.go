package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/ildarusmanov/msofficepreview/interfaces"
)

type PreviewsController struct {
    generator interfaces.PreviewGenerator
}

func CreatePreviewsController(generator interfaces.PreviewGenerator) *PreviewsController {
    return &PreviewsController{generator}
}

func (c *PreviewsController) Create(ctx *gin.Context) {
    fileId := ctx.Params.ByName("fileId")
    previewLink, err := c.generator.GetPreviewLink(fileId)

    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"Url": previewLink})
}
