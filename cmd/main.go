package main

import (
	"fwd/internal/server"
	"log"
)

func main() {
	if err := server.NewServer().Start(); err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
}
