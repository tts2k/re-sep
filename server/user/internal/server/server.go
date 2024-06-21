package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	config "re-sep-user/internal/system/config"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	config := config.Config()
	port, _ := strconv.Atoi(config.HTTPPort)
	NewServer := &Server{
		port: port,
	}

	// Register route handlers
	serverHandler := http.NewServeMux()
	serverHandler.Handle("/", NewServer.RegisterRoutes())

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.BaseURL, port),
		Handler:      serverHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
