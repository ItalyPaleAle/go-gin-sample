package main

import (
	"context"
	"log"

	"github.com/italypaleale/go-gin-sample/server"
)

func main() {
	// Create the Server object
	srv := server.Server{}
	err := srv.Init()
	if err != nil {
		log.Fatal("Error while initializing the server:", err)
		return
	}

	// Start the server in background and block until the server is shut down
	srv.Start(context.Background())
}
