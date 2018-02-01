package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type StatusController struct {}

func CreateStatusController() *StatusController {
    return &StatusController{}
}

func (c *StatusController) Check(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
