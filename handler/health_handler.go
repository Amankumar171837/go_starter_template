package handler

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "running",
	})
}

func (h *HealthHandler) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"ping": "pong",
	})
}
