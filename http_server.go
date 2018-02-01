package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/msofficepreview/controllers"
	"github.com/ildarusmanov/msofficepreview/interfaces"
)

func CreateRouter(serviceLocator interfaces.ServiceLocator) *gin.Engine {
	wopiC := controllers.CreateWopiController(
		serviceLocator.Get("Storage").(interfaces.Storage),
		serviceLocator.Get("TokenProvider").(interfaces.TokenProvider),
	)

	previewsC := controllers.CreatePreviewsController(
		serviceLocator.Get("PreviewGenerator").(interfaces.PreviewGenerator),
	)

	statusC := controllers.CreateStatusController()

	r := gin.Default()
	wopi := r.Group("/wopi")
	{
		wopi.GET("/files/:fileId/contents", wopiC.CreateAction(wopiC.GetFile))
		wopi.GET("/files/:fileId", wopiC.CreateAction(wopiC.CheckFileInfo))
	}
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/previews", previewsC.Create)
		apiv1.GET("/status/check", statusC.Check)
	}

	return r
}
