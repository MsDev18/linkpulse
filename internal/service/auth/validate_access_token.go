package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/claims"
	"linkpulse/internal/pkg/richerror"
	"time"
)

func (s Service) ValidateAccessToken(ctx context.Context, token string) (dto.ValidateAccessTokenResponse, error) {
	const op = "auth-service.ValidateAccessToken"

	if err := ctx.Err(); err != nil {
		return dto.ValidateAccessTokenResponse{}, richerror.New(op).
			SetMessage("context err").
			SetKind(richerror.KindTimeout).
			SetErr(err)
	}

	claim, err := claims.ParseAccessToken(
		s.config.AccessTokenSecret,
		token,
	)

	if err != nil {
		return dto.ValidateAccessTokenResponse{}, err
	}

	if claim.UserID == 0 || claim.SessionID == 0 {
		return dto.ValidateAccessTokenResponse{}, richerror.New(op).
			SetMessage("invalid access token").
			SetKind(richerror.KindUnauthorize)
	}

	session, err := s.repository.GetSessionByID(ctx, claim.SessionID)
	if err != nil {
		return dto.ValidateAccessTokenResponse{}, err
	}

	if session == nil || session.UserID != claim.UserID || session.RevokeAt != nil || !time.Now().Before(session.ExpiresAt) {
		return dto.ValidateAccessTokenResponse{}, richerror.New(op).
			SetMessage("invalid or expire session").
			SetKind(richerror.KindUnauthorize)
	}

	return dto.ValidateAccessTokenResponse{
		UserID:    session.UserID,
		SessionID: session.ID,
	}, nil
}
