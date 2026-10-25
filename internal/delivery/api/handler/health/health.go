package health

import (
	"linkpulse/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Health (c *gin.Context) {
	const op = "health-handler.Health"

	response.New(c).Send(http.StatusOK , "Health Check ✅" , nil)
}