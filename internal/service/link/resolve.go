package link

import (
	"context"
	"linkpulse/internal/pkg/richerror"
	"time"
)

func (s Service) Resolve(ctx context.Context, code string) (string, error) {
	const op = "link-service.Resolve"

	link, err := s.repository.GetOneByCode(ctx, code)
	if err != nil {
		return "", err
	}

	if link == nil || (link.ExpiresAt != nil && !time.Now().Before(*link.ExpiresAt)) {
		return "", richerror.New(op).
			SetMessage("link not found or expired").
			SetKind(richerror.KindNotFound)
	}

	return link.OriginalURL, nil
}
