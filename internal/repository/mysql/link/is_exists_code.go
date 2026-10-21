package link

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) IsExistsCode(ctx context.Context, code string) (bool, error) {
	const op = "link-repository.IsExistsCode"

	query := `SELECT * FROM link WHERE code = ?`

	row := r.conn.DB.QueryRowContext(ctx, query, code)

	var l entity.Link
	err := row.Scan(&l.ID, &l.UserID, &l.OriginalURL, &l.Code, &l.ExpiresAt, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return true, nil
}
