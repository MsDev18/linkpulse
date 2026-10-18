package main

import (
	mysql "linkpulse/internal/repository/mysql"
	service "linkpulse/internal/service/auth"
	handler "linkpulse/internal/api/handler/auth"
	validator "linkpulse/internal/validator/auth"
	repository "linkpulse/internal/repository/mysql/auth"
)

func SetupAuthModule (config service.Config , mysqlRepo mysql.MySQLDB) handler.Hander {
	repo := repository.New(mysqlRepo)
	val := validator.New()
	svc := service.New(config, repo)
	return handler.New(svc , val)
}