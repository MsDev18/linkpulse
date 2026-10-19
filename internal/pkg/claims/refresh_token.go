package claims

import (
	"linkpulse/internal/pkg/richerror"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenClaims struct {
	UserID uint `json:"user-id"`
	jwt.RegisteredClaims
}

func GenRefreshToken(secret string, userID uint, expiresAt time.Duration) (string, error) {
	const op = "claims.GenRefreshToken"

	claim := RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	jwtToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", richerror.New(op).
			SetMessage("unexpected err refresh token").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	return jwtToken, nil
}

func ParseRefreshToken(secret string, tokenStr string) (RefreshTokenClaims, error) {
	const op = "claims.ParseRefreshToken"

	claim := RefreshTokenClaims{}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claim,
		func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return RefreshTokenClaims{}, richerror.New(op).
			SetMessage("invalid token").
			SetKind(richerror.KindUnauthorize).
			SetErr(err)
	}

	if !token.Valid || claim.UserID == 0 {
		return RefreshTokenClaims{}, richerror.New(op).
			SetMessage("invalid access token").
			SetKind(richerror.KindUnauthorize)
	}

	return claim, nil
}
