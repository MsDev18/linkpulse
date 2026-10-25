package auth

import (
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Hander) CheckOtp(c *gin.Context) {
	const op = "auth-handler.CheckOtp"

	ctx := c.Request.Context()

	var req dto.CheckOtpRequest
	if err := c.ShouldBind(&req); err != nil {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("can't bind request body").
				SetKind(richerror.KindBadRequest).
				SetErr(err),
		)
		return
	}

	if err := h.validator.CheckOtp(ctx, req); err != nil {
		response.New(c).Error(err)
		return
	}

	res, err := h.service.CheckOtp(ctx, req)
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK, "Login Successfully", res)
}
