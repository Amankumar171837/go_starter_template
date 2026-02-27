package resource

import (
	"go_starter_template/config"
	"go_starter_template/handler"
	"go_starter_template/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterResourceRoutes(r *gin.RouterGroup, postHandler *handler.PostHandler, cfg *config.Config) {
	r.Use(middleware.AuthMiddleware(cfg))
	{
		r.POST("/posts", postHandler.CreatePost)
		r.GET("/posts", postHandler.GetPosts)
		r.GET("/posts/:id", postHandler.GetPost)
		r.PUT("/posts/:id", postHandler.UpdatePost)
		r.DELETE("/posts/:id", postHandler.DeletePost)
	}
}
