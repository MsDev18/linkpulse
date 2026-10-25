package router

import (
	auth "linkpulse/internal/delivery/api/handler/auth"
	health "linkpulse/internal/delivery/api/handler/health"
	link "linkpulse/internal/delivery/api/handler/link"
	authmiddleware "linkpulse/internal/delivery/api/middleware/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	e              *gin.Engine
	health         health.Handler
	auth           auth.Hander
	link           link.Handler
	authMiddleware authmiddleware.Middleware
}

func New(e *gin.Engine, health health.Handler, auth auth.Hander, link link.Handler, authMiddleware authmiddleware.Middleware) Router {
	return Router{
		e:              e,
		health:         health,
		auth:           auth,
		link:           link,
		authMiddleware: authMiddleware,
	}
}

func (r Router) RegisterRoute() {
	r.RegisterHealthRoute()
	r.RegisterAuthRoute()
	r.RegisterLinkRoute()
}
