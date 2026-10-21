package link

import (
	"context"
	"linkpulse/internal/pkg/richerror"
)

func (r Repository) DeleteByID(
	ctx context.Context,
	userID, linkID uint,
) error {
	const op = "link-repository.Delete"

	query := `DELETE FROM link WHERE id = ? AND user_id = ?`

	result, err := r.conn.DB.ExecContext(ctx, query, linkID, userID)
	if err != nil {
		return richerror.New(op).
			SetMessage("failed to delete link").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return richerror.New(op).
			SetMessage("failed to get affected rows").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	if count == 0 {
		return richerror.New(op).
			SetMessage("link not found").
			SetKind(richerror.KindNotFound)
	}

	return nil
}