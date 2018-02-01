package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusController struct{}

func CreateStatusController() *StatusController {
	return &StatusController{}
}

func (c *StatusController) Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
