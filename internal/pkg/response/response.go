package response

import "github.com/gin-gonic/gin"

type Body struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    any               `json:"data,omitempty"`
	Errors  map[string]any `json:"errors,omitempty"`
}

type Responder struct {
	c *gin.Context
}

func New(c *gin.Context) Responder {
	return Responder{
		c: c,
	}
}
