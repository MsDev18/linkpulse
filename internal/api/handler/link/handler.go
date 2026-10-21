package link

import (
	service "linkpulse/internal/service/link"
	validator "linkpulse/internal/validator/link"
)

type Handler struct {
	service service.Service
	validator validator.Validator
}

func New (service service.Service, validator validator.Validator) Handler {
	return Handler{
		service: service,
		validator: validator,
	}
}