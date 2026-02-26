package main

import (
	"log"

	"go_starter_template/internal/config"
	"go_starter_template/internal/handler"
	"go_starter_template/internal/middleware"
	"go_starter_template/internal/repository/postgres"
	"go_starter_template/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database connection

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config", err)
	}

	db, err := postgres.NewDB((cfg.Database.URL))
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	log.Println("connected to database")

	// Dependencies
	userRepo := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	postRepo := postgres.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)

	// Server setup

	gin.SetMode(gin.ReleaseMode) // Production mode (less verbose logging)
	// Create new router (no default middleware)
	r := gin.New()

	// Middleware

	r.Use(gin.Recovery())      // Catches panics and returns 500 instead of crashing
	r.Use(middleware.Logger()) // Logs each request (method, path, status, duration)

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "running",
		})
	})
	api := r.Group("/api/v2")
	{
		public := api.Group("/public")
		{
			public.GET("/ping", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"ping": "pong",
				})
			})

			public.GET("/users",userHandler.GetUsers)
		}

		identity := api.Group(("/identity"))
		{
			identity.POST("/users/new", authHandler.Register)
			identity.POST("/sessions", authHandler.Login)
			identity.POST("/sessions/refresh", authHandler.Refresh)
		}

		// Private Routes (Protected by JWT)
		private := api.Group("/resource")
		private.Use(middleware.AuthMiddleware(cfg))
		{
			private.POST("/posts", postHandler.CreatePost)
			private.GET("/posts", postHandler.GetPosts)
			private.GET("/posts/:id", postHandler.GetPost)
			private.PUT("/posts/:id", postHandler.UpdatePost)
			private.DELETE("/posts/:id", postHandler.DeletePost)
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server: ", err)
	}
}
