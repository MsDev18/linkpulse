package auth

import (
	"context"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) UpdateOtpByUserID(ctx context.Context, o entity.Otp) (error) {
	const op = "auth-repository.UpdateOtpByID"

	query := `UPDATE otp SET code = ?, expires_at = ? WHERE user_id = ?`

	res, err := r.conn.DB.ExecContext(ctx, query, o.Code, o.ExpiresAt, o.UserID)
	if err != nil {
		return richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return  richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	if count == 0 {
		return richerror.New(op).
			SetMessage("not found otp record").
			SetKind(richerror.KindNotFound)
	}
	return nil
}
