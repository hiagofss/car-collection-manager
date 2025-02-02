package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Healthcheck resource
// @Description Get health status
// @Tags app
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "healthy",
		"status":  "UP",
	})
}

// @Summary Ping Pong resource
// @Description Get pong response
// @Tags app
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingPongHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
