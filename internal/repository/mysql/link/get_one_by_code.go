package link

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetOneByCode(ctx context.Context, code string) (*entity.Link, error) {
	const op = "link-repository.GetOneByCode"

	query := `SELECT id, user_id, original_url, code, expires_at, created_at, updated_at FROM link WHERE code = ?`
	row := r.conn.DB.QueryRowContext(ctx, query, code)

	var link entity.Link
	if err := row.Scan(&link.ID, &link.UserID, &link.OriginalURL, &link.Code, &link.ExpiresAt, &link.CreatedAt, &link.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, richerror.New(op).
			SetMessage("failed to get link by code").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	return &link, nil
}
