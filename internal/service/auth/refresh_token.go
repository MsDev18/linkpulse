package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/claims"
	"linkpulse/internal/pkg/richerror"
	"time"
)

func (s Service) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	const op = "auth-service.RefreshToken"

	claim, err := claims.ParseRefreshToken(s.config.RefreshTokenSecret, req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	session, err := s.repository.GetSessionByID(ctx, claim.SessionID)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	if session == nil {
		return dto.RefreshTokenResponse{}, richerror.New(op).
			SetMessage("not found session").
			SetKind(richerror.KindUnauthorize).
			SetErr(err)
	}

	if claim.UserID == 0 ||
		claim.SessionID == 0 ||
		session.UserID != claim.UserID ||
		session.RevokeAt != nil ||
		!time.Now().Before(session.ExpiresAt) {
		return dto.RefreshTokenResponse{}, richerror.New(op).
			SetMessage("invalid or expired session").
			SetKind(richerror.KindUnauthorize)
	}

	accessToken, err := claims.GenAccessToken(s.config.AccessTokenSecret, claim.UserID, session.ID, s.config.AccessTokenDuration)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	return dto.RefreshTokenResponse{
		Token: dto.Token{
			AccessToken: accessToken,
		},
	}, nil
}
