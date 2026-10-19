package auth

import (
	authmiddleware "linkpulse/internal/api/middleware/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Hander) GetMe(c *gin.Context) {
	const op = "auth-handler.GetMe"

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
				SetMessage("invalid user id").
				SetKind(richerror.KindUnauthorize),
		)
		return
	}

	res, err := h.service.GetMe(ctx , userID)
	if err != nil {
		response.New(c).Error(err)
		return 
	}

	response.New(c).Send(http.StatusOK , "OK" , res)
}
