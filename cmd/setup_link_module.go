package main

import (
	handler "linkpulse/internal/api/handler/link"
	"linkpulse/internal/repository/mysql"
	repository "linkpulse/internal/repository/mysql/link"
	service "linkpulse/internal/service/link"
	validator "linkpulse/internal/validator/link"
)

func SetupLinkModule(config service.Config, mysqlRepo mysql.MySQLDB) handler.Handler {
	repo := repository.New(mysqlRepo)
	val := validator.New()
	svc := service.New(config, repo)
	return handler.New(svc, val)
}
