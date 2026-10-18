package auth

import (
	"context"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"

	"github.com/go-sql-driver/mysql"
)

func (r Repository) CreateUser(ctx context.Context, u entity.User) (*entity.User, error) {
	const op = "auth-repository.CreateUser"

	query := `INSERT INTO user(name, phone_number) VALUES(?,?)`

	res, err := r.conn.DB.ExecContext(ctx, query, u.Name, u.PhoneNumber)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, richerror.New(op).
				SetMessage("this phone number already exists").
				SetKind(richerror.KindConflict).
				SetErr(err)
		}
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)
	
	return &u, nil
}
