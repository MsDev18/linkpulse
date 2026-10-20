package main

import (
	"linkpulse/internal/api/server"
	"linkpulse/internal/migrator"
	"linkpulse/internal/repository/mysql"
	"linkpulse/internal/repository/redis"
	"linkpulse/internal/repository/redis/ratelimiter"
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
	redisRepo := redis.New(config.Redis)
	defer redisRepo.Conn.Close()

	_ = ratelimiter.New(redisRepo)


	// handlers
	health := SetupHealthModule()
	auth , authMiddleware:= SetupAuthModule(config.Auth, mysqlRepo)
	
	httpServer := server.New(config.Server , health , auth, authMiddleware)
	httpServer.Run()
}