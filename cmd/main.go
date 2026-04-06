package main

import (
	"fwd/internal/config"
	"fwd/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	if err := server.NewServer(cfg.Server).Start(); err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
}
