package main

import (
	mysql "linkpulse/internal/repository/mysql"
	service "linkpulse/internal/service/auth"
	handler "linkpulse/internal/api/handler/auth"
	validator "linkpulse/internal/validator/auth"
	repository "linkpulse/internal/repository/mysql/auth"
	middleware "linkpulse/internal/api/middleware/auth"
)

func SetupAuthModule (config service.Config , mysqlRepo mysql.MySQLDB) (handler.Hander, middleware.Middleware) {
	repo := repository.New(mysqlRepo)
	val := validator.New()
	svc := service.New(config, repo)
	authMiddleware := middleware.New(repo , config.AccessTokenSecret)
	handler :=handler.New(svc , val)
	return handler, authMiddleware
}