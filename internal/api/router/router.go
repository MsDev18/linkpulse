package router

import (
	health "linkpulse/internal/api/handler/health"
	auth "linkpulse/internal/api/handler/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	e *gin.Engine
	health health.Handler
	auth auth.Hander
}

func New(e *gin.Engine , health health.Handler, auth auth.Hander) Router {
	return Router{
		e: e,
		health: health,
		auth: auth,
	}
}


func (r Router) RegisterRoute () {
	r.RegisterHealthRoute()
	r.RegisterAuthRoute()
}