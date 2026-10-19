package auth

import (
	"context"
	"linkpulse/internal/entity"
	"time"
)

type Service struct {
	repository Repository
	config     Config
}

type Config struct {
	OtpCodeDuration      time.Duration `koanf:"otp_code_duration"`
	AccessTokenSecret    string        `koanf:"access_token_secret"`
	AccessTokenDuration  time.Duration `koanf:"access_token_duration"`
	RefreshTokenSecret   string        `koanf:"refresh_token_secret"`
	RefreshTokenDuration time.Duration `koanf:"refresh_token_duratioon"`
}

type Repository interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumner string) (*entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	GetOtpByUserID(ctx context.Context, userID uint) (*entity.Otp, error)
	UpdateOtpByUserID(ctx context.Context, otp entity.Otp) error
	CreateOtp(ctx context.Context, otp entity.Otp) (*entity.Otp, error)
	GetUserByID(ctx context.Context, id uint) (*entity.User, error)
}

func New(config Config, repository Repository) Service {
	return Service{
		repository: repository,
		config:     config,
	}
}
