package main

import "linkpulse/internal/delivery/api/handler/health"

func SetupHealthModule () health.Handler {
	handler := health.New()
	return handler
}