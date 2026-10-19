package auth

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetUserByID(ctx context.Context, userID uint) (*entity.User, error) {
	const op = "auth-repository.GetUserByID"

	query := `SELECT * FROM user WHERE id = ?`

	row := r.conn.DB.QueryRowContext(ctx, query, userID)

	u := &entity.User{}
	if err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return u, nil
}
