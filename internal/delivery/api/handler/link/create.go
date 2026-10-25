package link

import (
	authmiddleware "linkpulse/internal/delivery/api/middleware/auth"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(c *gin.Context) {
	const op = "link-handler.Create"

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

	var req dto.CreateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.New(c).Error(
			richerror.New(op).
				SetMessage("can't bind request").
				SetKind(richerror.KindUnexpected).
				SetErr(err),
		)
		return
	}

	if err := h.validator.CreateLink(ctx, req) ; err != nil {
		response.New(c).Error(err)
		return
	}

	res , err := h.service.Create(ctx , userID , req)
	if err != nil {
		response.New(c).Error(err)
		return
	}

	response.New(c).Send(http.StatusCreated, "create successfully" , res)
}
