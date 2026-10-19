package auth

import (
	"context"
	"linkpulse/internal/pkg/richerror"
	"time"
)

func (r Repository) RevokeSessionByID(ctx context.Context, sessionID uint) error {
	const op = "auth-repository.RevokeSessionByID"

	revokeAt := time.Now()

	query := `UPDATE session SET revoke_at = ? WHERE id = ?`

	res, err := r.conn.DB.ExecContext(ctx, query, revokeAt, sessionID)
	if err != nil {
		return richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return richerror.New(op).
			SetMessage("unexpected err").
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	if count == 0 {
		return richerror.New(op).
			SetMessage("not found session").
			SetKind(richerror.KindNotFound).
			SetErr(err)
	}

	return nil
}
