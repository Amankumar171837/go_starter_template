package server

import (
	"go_starter_template/api/v2/identity"
	"go_starter_template/api/v2/public"
	"go_starter_template/api/v2/resource"
	"go_starter_template/handler"
	"go_starter_template/middleware"
	"go_starter_template/repository/postgres"
	"go_starter_template/service"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(server *Server) {
	// Use custom middlewares
	server.Gin.Use(HandleOption)
	server.Gin.Use(gin.Recovery())
	server.Gin.Use(middleware.Logger())

	// Initialize Repositories
	userRepo := postgres.NewUserRepository(server.DB)
	postRepo := postgres.NewPostRepository(server.DB)

	// Initialize Services
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)
	authService := service.NewAuthService(userRepo, server.Cfg)

	// Initialize Handlers
	healthHandler := handler.NewHealthHandler()
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	authHandler := handler.NewAuthHandler(authService)

	// Global Health Route
	public.RegisterHealthRoute(server.Gin.Group(""), healthHandler)

	// API v2 Routes
	apiV2 := server.Gin.Group("/api/v2")
	{
		public.RegisterPublicRoutes(apiV2.Group("/public"), healthHandler, userHandler)
		identity.RegisterIdentityRoutes(apiV2.Group("/identity"), authHandler)
		resource.RegisterResourceRoutes(apiV2.Group("/resource"), postHandler, server.Cfg)
	}
}

// HandleOption sets security headers and CORS options
func HandleOption(c *gin.Context) {
	// Add CORS headers
	SetCORSHeaders(c)
	allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := make(map[string]bool)

	if allowedOriginsStr != "" {
		for _, origin := range strings.Split(allowedOriginsStr, ",") {
			allowedOrigins[strings.TrimSpace(origin)] = true
		}
	}

	origin := c.Request.Header.Get("Origin")
	if allowedOrigins[origin] {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	// Add Cache-Control headers
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func SetCORSHeaders(c *gin.Context) {
	allowCredentials := os.Getenv("CORS_ALLOW_CREDENTIALS")
	allowHeaders := os.Getenv("CORS_ALLOW_HEADERS")
	allowMethods := os.Getenv("CORS_ALLOW_METHODS")

	if !(allowCredentials == "" || allowHeaders == "" || allowMethods == "") {
		c.Header("Access-Control-Allow-Credentials", allowCredentials)
		c.Header("Access-Control-Allow-Headers", allowHeaders)
		c.Header("Access-Control-Allow-Methods", allowMethods)
	}
}
