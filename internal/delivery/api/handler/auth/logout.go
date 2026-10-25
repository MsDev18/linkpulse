package auth

import (
	"linkpulse/internal/delivery/api/middleware/auth"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Hander) Logout(c *gin.Context) {
	const op = "auth-handler.Logout"

	ctx := c.Request.Context()

	value, exists := c.Get(auth.SESSION_ID_KEY)
	if !exists {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("login required").
				SetKind(richerror.KindUnauthorize),
		)
		return
	}

	sessionID, ok := value.(uint)
	if !ok {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("invalid session id").
				SetKind(richerror.KindUnauthorize),
		)
		return
	}

	err := h.service.Logout(ctx , sessionID)
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusOK , "logout successfully ❌", nil)
}
