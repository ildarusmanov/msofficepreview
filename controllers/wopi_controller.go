package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/ildarusmanov/msofficepreview/interfaces"
)

type WopiController struct {
    storage interfaces.Storage
    tokenProvider interfaces.TokenProvider
}

func CreateWopiController(storage interfaces.Storage, tokenProvider interfaces.TokenProvider) *WopiController {
    return &WopiController{
        storage: storage,
        tokenProvider: tokenProvider,
    }
}

func (c *WopiController) CheckFileInfo(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{})
}

func (c *WopiController) GetFile(ctx *gin.Context) {
    ctx.Status(http.StatusOK)
}
