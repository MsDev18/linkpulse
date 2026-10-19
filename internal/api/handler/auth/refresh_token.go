package auth

import (
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Hander) RefreshToken (c *gin.Context) {
	const op = "auth-handler.RefreshToken"

	ctx := c.Request.Context()

	var req dto.RefreshTokenRequest
	if err := c.ShouldBind(&req) ; err != nil {
		response.New(c).Error(
			richerror.New(op).
			SetMessage("can't bind data").
			SetKind(richerror.KindBadRequest).
			SetErr(err) ,
		)
		return 
	}

	res, err := h.service.RefreshToken(ctx , req)
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK , "OK" , res)
}