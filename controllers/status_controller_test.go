package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateStatusController(t *testing.T) {
	controller := CreateStatusController()
	assert.NotNil(t, controller)
}

func TestCheckStatusController(t *testing.T) {
	controller := CreateStatusController()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	controller.Check(ctx)

	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
}
