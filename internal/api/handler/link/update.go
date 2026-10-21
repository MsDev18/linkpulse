package link

import (
	authmiddleware "linkpulse/internal/api/middleware/auth"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) Update(c *gin.Context) {
	const op = "link-handler.Update"

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
	id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if err != nil || id == 0 {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("invalid link id").
				SetKind(richerror.KindBadRequest),
		)
		return
	}
	var req dto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("can't bind request").
				SetKind(richerror.KindBadRequest).
				SetErr(err),
		)
		return
	}
	if err := h.validator.UpdateLink(ctx, req); err != nil {
		response.New(c).Error(err)
		return
	}
	res, err := h.service.Update(ctx, userID, uint(id), req)
	if err != nil {
		response.New(c).Error(err)
		return
	}
	response.New(c).Send(http.StatusOK, "link updated successfully", res)
}
