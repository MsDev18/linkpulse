package auth

import (
	"context"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) CreateSession(ctx context.Context, session entity.Session) (*entity.Session, error) {
	const op = "auth-service.CreateSession"

	query := `INSERT INTO session(user_id, expires_at , revoke_at) VALUES(?,?,?)`

	res, err := r.conn.DB.ExecContext(ctx,query, session.UserID, session.ExpiresAt, session.RevokeAt)
	if err != nil {
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	id, _ := res.LastInsertId()
	session.ID = uint(id)

	return &session, nil
}
