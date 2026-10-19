package router

import (
	health "linkpulse/internal/api/handler/health"
	auth "linkpulse/internal/api/handler/auth"
	authmiddleware "linkpulse/internal/api/middleware/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	e *gin.Engine
	health health.Handler
	auth auth.Hander
	authMiddleware authmiddleware.Middleware
}

func New(e *gin.Engine , health health.Handler, auth auth.Hander, authMiddleware authmiddleware.Middleware) Router {
	return Router{
		e: e,
		health: health,
		auth: auth,
		authMiddleware: authMiddleware,
	}
}


func (r Router) RegisterRoute () {
	r.RegisterHealthRoute()
	r.RegisterAuthRoute()
}