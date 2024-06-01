package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"re-sep-user/internal/database"
	"re-sep-user/internal/system"
)

type Server struct {
	db   *database.Service
	port int
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(system.Config().HTTPPort)
	NewServer := &Server{
		port: port,

		db: database.NewDBService(),
	}

	// Register route handlers
	serverHandler := http.NewServeMux()
	serverHandler.Handle("/api/v1", NewServer.RegisterRoutes())

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      serverHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
