package controllers

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/gin-gonic/gin"
)

func TestCreateWopiController(t *testing.T) {
  controller := CreateWopiController(wopiStorage)

  assert.NotNil(t, controller)
}
