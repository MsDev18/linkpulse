package auth

import (
	"context"
	"linkpulse/internal/pkg/richerror"
)

func (s Service) Logout(ctx context.Context, sessionID uint) error {
	const op = "auth-service.Logout"

	session, err := s.repository.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}

	if session == nil {
		return richerror.New(op).
			SetMessage("not found session").
			SetKind(richerror.KindUnauthorize)
	}

	err = s.repository.RevokeSessionByID(ctx , session.ID)
	if err != nil {
		return err
	}

	return nil
}
