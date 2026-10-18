package main

import "linkpulse/internal/api/handler/health"

func SetupHealthModule () health.Handler {
	handler := health.New()
	return handler
}