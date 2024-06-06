package main

import (
	"fmt"

	"re-sep-user/internal/server"
	config "re-sep-user/internal/system/config"
	logger "re-sep-user/internal/system/logger"
)

func main() {
	logger.InitLogger()

	server := server.NewServer()

	fmt.Printf("Server listening on port: %s\n", config.Config().HTTPPort)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
