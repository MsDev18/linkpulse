package main

import (
	"linkpulse/internal/api/server"
	"linkpulse/internal/pkg/migrator"
	"linkpulse/internal/repository/mysql"
)

func main () {
	// load configurations 
	config := LoadConfig()

	// migrator 
	m := migrator.New(config.MySQL)
	defer m.Close()
	m.Up()

	// new instance myswl repo 
	mysqlRepo := mysql.New(config.MySQL)

	// handlers
	health := SetupHealthModule()
	auth := SetupAuthModule(config.Auth, mysqlRepo)

	httpServer := server.New(config.Server , health , auth)
	httpServer.Run()
}