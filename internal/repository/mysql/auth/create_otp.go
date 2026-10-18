package auth

import (
	"context"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"

	"github.com/go-sql-driver/mysql"
)

func (r Repository) CreateOtp(ctx context.Context, otp entity.Otp) (*entity.Otp, error) {
	const op = "auth-repository.CreateOtp"

	query := `INSERT INTO otp(user_id,code,expires_at) VALUES(?,?,?)`

	res, err := r.conn.DB.ExecContext(ctx, query, otp.UserID, otp.Code, otp.ExpiresAt)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, richerror.New(op).
				SetMessage("otp with this user id already exists").
				SetKind(richerror.KindConflict).
				SetErr(err)
		}
		return nil, richerror.New(op).
			SetMessage("unexpected error").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	id, _ := res.LastInsertId()
	otp.ID = uint(id)

	return &otp, nil
}
