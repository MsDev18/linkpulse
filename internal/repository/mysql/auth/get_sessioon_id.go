package auth

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetSessionByID(ctx context.Context, id uint) (*entity.Session, error) {
	const op = "auth-repository.GetSessionByUserID"

	query := `SELECT * FROM session WHERE id = ? `

	row := r.conn.DB.QueryRowContext(ctx, query, id)

	s := &entity.Session{}
	if err := row.Scan(&s.ID, &s.UserID, &s.ExpiresAt, &s.RevokeAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return s, nil
}
