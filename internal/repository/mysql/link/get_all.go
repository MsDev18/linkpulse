package link

import (
	"context"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) GetAll(ctx context.Context, userID uint) ([]*entity.Link, error) {
	const op = "link-repository.GetAll"

	query := `SELECT * FROM link WHERE user_id = ?`

	rows, err := r.conn.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, richerror.New(op).
			SetMessage("unexpected err").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	defer rows.Close()

	var list = make([]*entity.Link, 0)
	for rows.Next() {
		l := &entity.Link{}
		if err := rows.Scan(&l.ID, &l.UserID, &l.OriginalURL, &l.Code, &l.ExpiresAt, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, richerror.New(op).
				SetMessage("can't scan data").
				SetKind(richerror.KindUnexpected).
				SetErr(err)
		}
		list = append(list, l)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).
			SetMessage("failed to read links").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}
	return list, nil
}
