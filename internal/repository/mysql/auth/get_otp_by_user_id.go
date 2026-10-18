package auth

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetOtpByUserID(ctx context.Context, userID uint) (*entity.Otp, error) {
	const op = "auth-repository.GetOtpByUserID"

	query := `SELECT * FROM otp WHERE user_id = ?`

	row := r.conn.DB.QueryRowContext(ctx, query, userID)

	var o entity.Otp
	if err := row.Scan(&o.ID, &o.UserID, &o.Code, &o.ExpiresAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return &o, nil
}
