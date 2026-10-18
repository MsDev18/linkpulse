package main

import (
	"linkpulse/internal/api/server"
	"linkpulse/internal/pkg/migrator"
)

func main () {
	// load configurations 
	config := LoadConfig()

	// migrator 
	m := migrator.New(config.MySQL)
	defer m.Close()
	m.Up()

	// handlers
	health := SetupHealthModule()

	httpServer := server.New(config.Server , health)
	httpServer.Run()
}