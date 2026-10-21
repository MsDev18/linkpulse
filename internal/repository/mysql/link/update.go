package link

import (
	"context"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
	"time"

	"github.com/go-sql-driver/mysql"
)

func (r Repository) Update(ctx context.Context, userID, linkID uint, link entity.Link) (*entity.Link, error) {
	const op = "link-repository.Update"

	link.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	query := `
  UPDATE link
  SET original_url = ?, code = ?, expires_at = ?, updated_at = ?
  WHERE id = ? AND user_id = ?
 `

	result, err := r.conn.DB.ExecContext(
		ctx,
		query,
		link.OriginalURL,
		link.Code,
		link.ExpiresAt,
		link.UpdatedAt,
		linkID,
		userID,
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
			SetMessage("failed to update link").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return nil, richerror.New(op).
			SetMessage("failed to get affected rows").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	// MySQL also reports zero when all values are unchanged.
	if count == 0 {
		var exists bool
		err := r.conn.DB.QueryRowContext(ctx,
			`SELECT EXISTS(SELECT 1 FROM link WHERE id = ? AND user_id = ?)`,
			linkID, userID,
		).Scan(&exists)
		if err != nil {
			return nil, richerror.New(op).
				SetMessage("failed to check link").
				SetKind(richerror.KindUnexpected).
				SetErr(err)
		}
		if !exists {
			return nil, richerror.New(op).
				SetMessage("link not found").
				SetKind(richerror.KindNotFound)
		}
	}

	return &link, nil
}
