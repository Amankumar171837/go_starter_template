package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go_starter_template/internal/config"
	"go_starter_template/internal/repository/postgres"
	"go_starter_template/internal/middleware"
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

	// Server setup

	gin.SetMode(gin.ReleaseMode) // Production mode (less verbose logging)
	// Create new router (no default middleware)
	r := gin.New()

	// Middleware

	r.Use(gin.Recovery())                             // Catches panics and returns 500 instead of crashing
	r.Use(middleware.Logger())                        // Logs each request (method, path, status, duration)

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "running",
		})
	})
	api := r.Group("/api/v2")
	{
		pubic := api.Group("/public")
		{
			pubic.GET("/ping", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"ping": "pong",
				})
			})
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server: ", err)
	}
}
