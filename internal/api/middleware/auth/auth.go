package auth

import (
	"linkpulse/internal/pkg/claims"
	"linkpulse/internal/pkg/response"
	"linkpulse/internal/pkg/richerror"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m Middleware) Auth() gin.HandlerFunc {
	const op = "middleware.Auth"
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		parts := strings.SplitN(auth, " ", -1)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.New(c).Error(
				richerror.New(op).
					SetMessage("invalid token").
					SetKind(richerror.KindUnauthorize),
			)
			c.Abort()
			return
		}

		token := parts[1]

		claim , err := claims.ParseAccessToken(m.config.accessTokenSecret , token)
		if err != nil {
			response.New(c).Error(err)
			c.Abort()
			return 
		}

		c.Set(USER_ID_KEY , claim.UserID)
	}
}
