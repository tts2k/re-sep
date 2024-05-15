package main

import (
	"fmt"
	"os"

	"re-sep-user/internal/server"
)

func main() {

	server := server.NewServer()

	fmt.Printf("Server listening on port: %s\n", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
