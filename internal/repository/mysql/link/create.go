package link

import (
	"context"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
	"time"

	"github.com/go-sql-driver/mysql"
)

func (r Repository) Create(ctx context.Context, link entity.Link) (*entity.Link, error) {
	const op = "link-repository.Create"

	now := time.Now().UTC().Truncate(time.Second)
	link.CreatedAt = now
	link.UpdatedAt = now
	
	query := `
		INSERT INTO link (
			user_id,
			original_url,
			code,
			expires_at, 
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.conn.DB.ExecContext(
		ctx,
		query,
		link.UserID,
		link.OriginalURL,
		link.Code,
		link.ExpiresAt,
		link.CreatedAt,
		link.UpdatedAt,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, richerror.New(op).
				SetMessage("this code already exists").
				SetKind(richerror.KindConflict).
				SetErr(err)
		}

		return nil, richerror.New(op).
			SetMessage("failed to create link").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	id, _ := result.LastInsertId()

	link.ID = uint(id)

	return &link, nil
}
