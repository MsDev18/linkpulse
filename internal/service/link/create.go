package link

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
	"strings"
)

func (s Service) Create(ctx context.Context, userID uint, req dto.CreateRequest) (dto.CreateResponse, error) {
	const op = "link-service.Create"

	if err := ctx.Err(); err != nil {
		return dto.CreateResponse{}, richerror.New(op).
			SetMessage("timeout err").
			SetKind(richerror.KindTimeout).
			SetErr(err)
	}

	var code = req.CustomCode
	if req.CustomCode == "" {
		generatedCode, err := s.generateCode()
		if err != nil {
			return dto.CreateResponse{}, err
		}
		code = generatedCode
	}

	exists, err := s.repository.IsExistsCode(ctx, code)
	if err != nil {
		return dto.CreateResponse{}, err
	}
	if exists {
		return dto.CreateResponse{}, richerror.New(op).
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
		return dto.CreateResponse{}, err
	}

	return dto.CreateResponse{
		ID:          link.ID,
		OriginalURL: link.OriginalURL,
		Code:        link.Code,
		ShortURL:    strings.TrimRight(s.config.PublicBaseURL, "/") + "/r/" + link.Code,
		ExpiresAt:   link.ExpiresAt,
		CreatedAt:   link.CreatedAt,
	}, nil
}

func (s Service) generateCode() (string, error) {
	const op = "link-service.generateCode"
	var randomBytes [6]byte

	if _, err := rand.Read(randomBytes[:]); err != nil {
		return "", richerror.New(op).
			SetMessage("can't generate random code").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	code := base64.RawURLEncoding.EncodeToString(randomBytes[:])

	return code, nil
}
