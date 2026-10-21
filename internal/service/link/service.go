package link

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
	"strings"
	"time"
)

type Service struct {
	config     Config
	repository Repository
}

type Config struct {
	PublicBaseURL string `koanf:"public_base_url"`
}

type Repository interface {
	Create(ctx context.Context, link entity.Link) (*entity.Link, error)
	IsExistsCode(ctx context.Context, code string) (bool, error)
	GetAll(ctx context.Context, userID uint) ([]*entity.Link, error)
	GetOneByID(ctx context.Context, id uint) (*entity.Link, error)
	DeleteByID(ctx context.Context, userID uint, linkID uint) error
}

func New(config Config, repository Repository) Service {
	return Service{
		config:     config,
		repository: repository,
	}
}

func (s Service) entityLinkMapToDto(id uint, originalURL string, code string, expiresAt *time.Time, createdAt time.Time) dto.GetOneResponse {
	return dto.GetOneResponse{
		ID:          id,
		OriginalURL: originalURL,
		Code:        code,
		ShortURL:    s.getShortURL(code),
		ExpiresAt:   expiresAt,
		CreatedAt:   createdAt,
	}
}

func (s Service) getShortURL(code string) string {
	return strings.TrimRight(s.config.PublicBaseURL, "/") + "/r/" + code
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
