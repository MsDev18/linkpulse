package link

import (
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetOneByID(c *gin.Context) {
	const op = "link-handler.GetOneByID"

	ctx := c.Request.Context()

	value := c.Param("id")

	id, err := strconv.Atoi(value)
	if err != nil {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("invalid link id").
				SetKind(richerror.KindBadRequest).
				SetErr(err),
		)
		return
	}

	res, err := h.service.GetOneByID(ctx, uint(id))
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK, "OK ✅", res)
}
