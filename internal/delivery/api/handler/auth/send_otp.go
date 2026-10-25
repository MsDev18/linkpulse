package auth

import (
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Hander) SendOtp(c *gin.Context) {
	const op = "auth-handler.SendOtp"

	var req dto.SendOtpRequest
	if err := c.ShouldBind(&req); err != nil {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("can't bind request body").
				SetKind(richerror.KindBadRequest).
				SetErr(err),
		)
		return
	}

	if err := h.validator.SendOtp(c.Request.Context(), req); err != nil {
		response.New(c).Error(err)
		return
	}

	res, err := h.service.SendOtp(c.Request.Context(), req)
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK, "send otp successfully", res)
}
