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

	statusController := controllers.CreateStatusController()

	r := gin.Default()
	wopi := r.Group("/wopi")
	{
		wopi.GET("/files/:fileId/contents", wopiController.GetFile)
		wopi.GET("/files/:fileId", wopiController.CheckFileInfo)
	}
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/previews", previewsController.Create)
		apiv1.GET("/status/check", statusController.Check)
	}

	return r
}
