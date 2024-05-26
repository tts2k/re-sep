package main

import (
	"fmt"

	"re-sep-user/internal/server"
	"re-sep-user/internal/system"
)

func main() {

	server := server.NewServer()

	fmt.Printf("Server listening on port: %s\n", system.Config().HTTPPort)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
