package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/controllers"
	"github.com/ildarusmanov/msofficepreview/interfaces"
)

func CreateRouter(serviceLocator interfaces.ServiceLocator) *gin.Engine {
	wopiController := controllers.CreateWopiController(
		serviceLocator.Get("Storage").(interfaces.Storage),
		serviceLocator.Get("TokenProvider").(interfaces.TokenProvider),
	)

	previewsController := controllers.CreatePreviewsController(
		serviceLocator.Get("PreviewGenerator").(interfaces.PreviewGenerator),
	)

	r := gin.Default()
	wopi := r.Group("/wopi")
	{
		wopi.GET("/files/:fileId", wopiController.CheckFileInfo)
		wopi.GET("/files/:fileId/contents", wopiController.GetFile)
	}
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/api/v1/:fileId", previewsController.Create)
	}

	return r
}
