package main

import (
	"log/slog"
	"os"

	"re-sep-content/internal/server"
)

func main() {
	if err := server.Serve(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
