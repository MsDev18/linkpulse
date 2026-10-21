package link

import (
	"context"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/richerror"
)

func (s Service) GetOneByID(ctx context.Context, id uint) (dto.GetOneResponse, error) {
	const op = "link-GetOneByID"

	link, err := s.repository.GetOneByID(ctx, id)
	if err != nil {
		return dto.GetOneResponse{}, err
	}

	if link == nil {
		return dto.GetOneResponse{}, richerror.New(op).
			SetMessage("not found link with this id").
			SetKind(richerror.KindNotFound)
	}

	return s.entityLinkMapToDto(
		link.ID,
		link.OriginalURL,
		link.Code,
		link.ExpiresAt,
		link.CreatedAt,
	), nil
}
