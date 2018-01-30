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
    ctx.JSON(http.StatusOK, gin.H{})
}
