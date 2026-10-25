package link

import (
	authmiddleware "linkpulse/internal/delivery/api/middleware/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAll(c *gin.Context) {
	const op = "link-handler.GetAll"

	ctx := c.Request.Context()

	value, exists := c.Get(authmiddleware.USER_ID_KEY)
	if !exists {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("login required").
				SetKind(richerror.KindUnauthorize),
		)
		return
	}

	userID, ok := value.(uint)
	if !ok {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("invalid user id ").
				SetKind(richerror.KindUnauthorize),
		)
		return
	}

	res, err := h.service.GetAll(ctx, userID)
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK, "OK ✅", res)
}
