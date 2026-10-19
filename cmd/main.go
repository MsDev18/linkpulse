package main

import (
	"linkpulse/internal/api/server"
	"linkpulse/internal/migrator"
	"linkpulse/internal/repository/mysql"
	authmiddleware "linkpulse/internal/api/middleware/auth"
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

	// middlewares 
	authMiddleware := authmiddleware.New(config.Auth.AccessTokenSecret)
	

	httpServer := server.New(config.Server , health , auth, authMiddleware)
	httpServer.Run()
}