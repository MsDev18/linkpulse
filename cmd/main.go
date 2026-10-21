package main

import (
	"linkpulse/internal/api/server"
	"linkpulse/internal/migrator"
	"linkpulse/internal/repository/mysql"
	"linkpulse/internal/repository/redis"
)

func main() {
	// load configurations
	config := LoadConfig()

	// migrator
	m := migrator.New(config.MySQL)
	defer m.Close()
	m.Up()

	// new instance myswl repo
	mysqlRepo := mysql.New(config.MySQL)
	redisRepo := redis.New(config.Redis)
	defer redisRepo.Conn.Close()

	// handlers
	health := SetupHealthModule()
	auth, authMiddleware := SetupAuthModule(config.Auth, mysqlRepo)
	link := SetupLinkModule(config.Link, mysqlRepo)

	httpServer := server.New(config.Server, health, auth, link, authMiddleware)
	httpServer.Run()
}
