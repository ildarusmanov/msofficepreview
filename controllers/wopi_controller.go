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

type ActionHandlerFunc func(*gin.Context, interfaces.Token)

func CreateWopiController(storage interfaces.Storage, tokenProvider interfaces.TokenProvider) *WopiController {
	return &WopiController{
		storage:       storage,
		tokenProvider: tokenProvider,
	}
}

func (c *WopiController) CreateAction(action ActionHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileId := ctx.Param("fileId")
		token, ok := c.tokenProvider.FindToken(fileId)

		if !ok || !c.tokenProvider.Validate(token) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		action(ctx, token)
	}
}

func (c *WopiController) CheckFileInfo(ctx *gin.Context, token interfaces.Token) {
	fileInfo, err := c.storage.GetFileInfo(token.GetFilePath())

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"BaseFileName": fileInfo.GetFileName(),
		"OwnerId":      fileInfo.GetOwnerId(),
		"Size":         fileInfo.GetSize(),
		"UserId":       token.GetValue(),
		"Version":      fileInfo.GetVersion(),
	})
}

func (c *WopiController) GetFile(ctx *gin.Context, token interfaces.Token) {
	content, err := c.storage.GetContents(token.GetFilePath())

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Data(http.StatusOK, "application/octet-stream", content)
}
