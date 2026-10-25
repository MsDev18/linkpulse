package auth

import (
	service "linkpulse/internal/service/auth"
	validator "linkpulse/internal/validator/auth"
)

type Hander struct {
	service   service.Service
	validator validator.Validator
}

func New(service service.Service, validator validator.Validator) Hander {
	return Hander{
		service:   service,
		validator: validator,
	}
}
