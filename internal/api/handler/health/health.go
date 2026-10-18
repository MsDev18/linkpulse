package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Health (c *gin.Context) {
	const op = "health-handler.Health"

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message" : "Health Check ✅" ,
	})
}