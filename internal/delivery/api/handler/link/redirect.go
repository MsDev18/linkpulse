package link

import (
	"linkpulse/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Redirect(c *gin.Context) {
	c.Header("Cache-Control", "no-store")

	originalURL, err := h.service.Resolve(c.Request.Context(), c.Param("code"))
	if err != nil {
		response.New(c).Error(err)
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
