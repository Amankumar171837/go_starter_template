package identity

import (
	"go_starter_template/handler"

	"github.com/gin-gonic/gin"
)

func RegisterIdentityRoutes(r *gin.RouterGroup, authHandler *handler.AuthHandler) {
	r.POST("/users/new", authHandler.Register)
	r.POST("/sessions", authHandler.Login)
	r.POST("/sessions/refresh", authHandler.Refresh)
}
