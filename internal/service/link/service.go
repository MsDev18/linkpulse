package link

import (
	"context"
	"linkpulse/internal/entity"
)

type Service struct {
	config Config
	repository Repository
}

type Config struct {
	PublicBaseURL string `koanf:"public_base_url"`
}

type Repository interface {
	Create(ctx context.Context, link entity.Link) (*entity.Link, error)
	IsExistsCode(ctx context.Context, code string) (bool, error)
}

func New(config Config, repository Repository) Service {
	return Service{
		config: config,
		repository: repository,
	} 
}