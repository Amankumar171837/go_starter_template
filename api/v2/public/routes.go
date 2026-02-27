package public

import (
	"go_starter_template/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup, healthHandler *handler.HealthHandler, userHandler *handler.UserHandler) {
	r.GET("/ping", healthHandler.Ping)
	r.GET("/users", userHandler.GetUsers)
}

func RegisterHealthRoute(r *gin.RouterGroup, healthHandler *handler.HealthHandler) {
	r.GET("/health", healthHandler.Health)
}
