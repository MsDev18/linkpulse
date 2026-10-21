package link

import (
	"context"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/richerror"
)

func (s Service) Update(ctx context.Context, userID, linkID uint, req dto.UpdateRequest) (dto.GetOneResponse, error) {
	const op = "link-service.Update"

	if err := ctx.Err(); err != nil {
		return dto.GetOneResponse{}, richerror.New(op).
			SetMessage("timeout err").
			SetKind(richerror.KindTimeout).
			SetErr(err)
	}

	item, err := s.repository.GetOneByID(ctx, linkID)
	if err != nil {
		return dto.GetOneResponse{}, err
	}
	if item == nil || item.UserID != userID {
		return dto.GetOneResponse{}, richerror.New(op).
			SetMessage("link not found").
			SetKind(richerror.KindNotFound)
	}
	if req.OriginalURL != "" {
		item.OriginalURL = req.OriginalURL
	}
	if req.CustomCode != "" {
		item.Code = req.CustomCode
	}
	if req.ExpiresAt != nil {
		value := req.ExpiresAt.UTC()
		item.ExpiresAt = &value
	}
	item, err = s.repository.Update(ctx, userID, linkID, *item)
	if err != nil {
		return dto.GetOneResponse{}, err
	}
	return s.entityLinkMapToDto(item.ID, item.OriginalURL, item.Code, item.ExpiresAt, item.CreatedAt), nil
}
