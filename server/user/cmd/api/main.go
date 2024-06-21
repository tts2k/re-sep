package main

import (
	"fmt"

	tokenDB "re-sep-user/internal/database/token"
	userDB "re-sep-user/internal/database/user"
	"re-sep-user/internal/server"
	config "re-sep-user/internal/system/config"
	logger "re-sep-user/internal/system/logger"
)

func main() {
	logger.InitLogger()

	//init DBs
	tokenDB.InitTokenDB()
	userDB.InitUserDB()

	server := server.NewServer()
	systemConfig := config.Config()

	fmt.Printf("Server listening on: %s:%s\n", systemConfig.BaseURL, systemConfig.HTTPPort)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
