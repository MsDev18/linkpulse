package link

import (
	"context"
	"database/sql"
	"errors"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetOneByID(ctx context.Context, id uint) (*entity.Link, error) {
	const op = "link-repository.GetOneByID"

	query := `SELECT * FROM link WHERE id = ?`

	row := r.conn.DB.QueryRowContext(ctx, query, id)

	var l entity.Link
	if err := row.Scan(&l.ID, &l.UserID, &l.OriginalURL, &l.Code, &l.ExpiresAt, &l.CreatedAt, &l.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return &l, nil
}
