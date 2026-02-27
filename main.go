package main

import (
	"log"

	"go_starter_template/config"
	"go_starter_template/repository/postgres"
	"go_starter_template/server"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database connection

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config", err)
	}

	dbWrapper, err := postgres.NewDB((cfg.Database.URL))
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer dbWrapper.Close()

	log.Println("connected to database")

	// Server setup

	gin.SetMode(gin.ReleaseMode) // Production mode (less verbose logging)

	// Create new server instance
	srv := server.NewServer(cfg, gin.New(), dbWrapper)

	// Configure Routes and Middlewares
	server.ConfigureRoutes(srv)

	// Run server
	if err := srv.Run("8080"); err != nil {
		log.Fatal("failed to run server: ", err)
	}
}
