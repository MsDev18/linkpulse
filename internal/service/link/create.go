package link

import (
	"context"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
)

func (s Service) Create(ctx context.Context, userID uint, req dto.CreateRequest) (dto.GetOneResponse, error) {
	const op = "link-service.Create"

	if err := ctx.Err(); err != nil {
		return dto.GetOneResponse{}, richerror.New(op).
			SetMessage("timeout err").
			SetKind(richerror.KindTimeout).
			SetErr(err)
	}

	var code = req.CustomCode
	if req.CustomCode == "" {
		generatedCode, err := s.generateCode()
		if err != nil {
			return dto.GetOneResponse{}, err
		}
		code = generatedCode
	}

	exists, err := s.repository.IsExistsCode(ctx, code)
	if err != nil {
		return dto.GetOneResponse{}, err
	}
	if exists {
		return dto.GetOneResponse{}, richerror.New(op).
			SetMessage("url code is not unique try again").
			SetKind(richerror.KindConflict)
	}

	link, err := s.repository.Create(ctx, entity.Link{
		UserID:      userID,
		OriginalURL: req.OriginalURL,
		Code:        code,
		ExpiresAt:   req.ExpiresAt,
	})

	if err != nil {
		return dto.GetOneResponse{}, err
	}

	return s.entityLinkMapToDto(
		link.ID ,
		link.OriginalURL,
		link.Code,
		link.ExpiresAt,
		link.CreatedAt,
	), nil
}






