package auth

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	const op = "auth-repository.GetUserByPhoneNumber"

	query := `SELECT * FROM user WHERE phone_number = ?`

	row := r.conn.DB.QueryRowContext(ctx, query, phoneNumber)

	var u entity.User
	if err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return &u, nil
}
