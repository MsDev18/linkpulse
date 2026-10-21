package link

import "context"

func (s Service) DeleteByID(ctx context.Context, userID, linkID uint) error {
	return s.repository.DeleteByID(ctx, userID, linkID)
}