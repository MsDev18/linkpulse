package link

import (
	"context"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/richerror"
)

func (s Service) GetAll(ctx context.Context, userID uint) ([]dto.GetOneResponse, error) {
	const op = "link-service.GetAll"

	linkListsEntity, err := s.repository.GetAll(ctx, userID)
	if err != nil {
		return []dto.GetOneResponse{}, err
	}

	if len(linkListsEntity) == 0 {
		return nil, richerror.New(op).
			SetMessage("not found link's").
			SetKind(richerror.KindNotFound)
	}

	linklist := make([]dto.GetOneResponse, len(linkListsEntity))
	for i, v := range linkListsEntity {
		linklist[i] = s.entityLinkMapToDto(
			v.ID,
			v.OriginalURL,
			v.Code,
			v.ExpiresAt,
			v.CreatedAt,
		)
	}

	return linklist, nil
}
