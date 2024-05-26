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

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
