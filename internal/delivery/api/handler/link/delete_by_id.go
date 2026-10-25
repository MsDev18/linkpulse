package link

import (
	authmiddleware "linkpulse/internal/delivery/api/middleware/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) Delete(c *gin.Context) {
	const op = "link-handler.Delete"

	ctx := c.Request.Context()

	value, exists := c.Get(authmiddleware.USER_ID_KEY)
	userID, ok := value.(uint)
	if !exists || !ok {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("login required").
				SetKind(richerror.KindUnauthorize),
		)
		return
	}

	linkID, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if err != nil || linkID == 0 {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("invalid link id").
				SetKind(richerror.KindBadRequest),
		)
		return
	}

	if err := h.service.DeleteByID(ctx, userID, uint(linkID)); err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK, "delete successfully ✅", nil)
}
