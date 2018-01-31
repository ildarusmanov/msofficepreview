package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/interfaces"
	"net/http"
)

type WopiController struct {
	storage       interfaces.Storage
	tokenProvider interfaces.TokenProvider
}

func CreateWopiController(storage interfaces.Storage, tokenProvider interfaces.TokenProvider) *WopiController {
	return &WopiController{
		storage:       storage,
		tokenProvider: tokenProvider,
	}
}

func (c *WopiController) CheckFileInfo(ctx *gin.Context) {
	accessToken := ctx.Params.ByName("accessToken")
	fileId := ctx.Params.ByName("fileId")

	if !c.tokenProvider.Validate(accessToken) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	fileInfo, err := c.storage.GetFileInfo(fileId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"BaseFileName": fileInfo.GetFileName(),
		"OwnerId":      fileInfo.GetOwnerId(),
		"Size":         fileInfo.GetSize(),
		"UserId":       accessToken,
		"Version":      fileInfo.GetVersion(),
	})
}

func (c *WopiController) GetFile(ctx *gin.Context) {
	accessToken := ctx.Params.ByName("accessToken")
	fileId := ctx.Params.ByName("fileId")

	if !c.tokenProvider.Validate(accessToken) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	content, err := c.storage.GetContents(fileId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Data(http.StatusOK, "application/octet-stream", content)
}
