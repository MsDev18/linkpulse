package auth

import (
	"context"
	"linkpulse/internal/entity"
)

const (
	USER_ID_KEY = "user-id"
	SESSION_ID_KEY = "session-id"
)

type Middleware struct {
	config     Config
	repository Repository
}

type Repository interface {
	GetSessionByID(ctx context.Context, id uint) (*entity.Session, error)
}

type Config struct {
	accessTokenSecret string
}

func New(repository Repository, accessSecret string) Middleware {
	return Middleware{
		config: Config{
			accessTokenSecret: accessSecret,
		},
		repository: repository,
	}
}
