package router

import (
	"linkpulse/internal/api/handler/health"

	"github.com/gin-gonic/gin"
)

type Router struct {
	e *gin.Engine
	health health.Handler
}

func New(e *gin.Engine , health health.Handler) Router {
	return Router{
		e: e,
		health: health,
	}
}


func (r Router) RegisterRoute () {
	r.RegisterHealthRoute()
}