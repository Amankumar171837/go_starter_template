package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go_starter_template/config"
	"go_starter_template/repository/postgres"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Cfg *config.Config
	Gin *gin.Engine
	DB  *postgres.DB
}

func NewServer(cfg *config.Config, engine *gin.Engine, db *postgres.DB) *Server {
	return &Server{
		Cfg: cfg,
		Gin: engine,
		DB:  db,
	}
}

func (s *Server) Run(addr string) error {
	log.Printf("Server starting on port %s", addr)
	return s.Gin.Run(":" + addr)
}

func (s *Server) Shutdown(ctx context.Context, srv *http.Server) error {
	// Gracefully shutdown the Gin HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error closing session of gin: %v", err)
	}

	// Close DB connection
	if s.DB != nil {
		s.DB.Close()
	} else {
		log.Println("Database connection is nil, skipping close.")
	}

	return nil
}
