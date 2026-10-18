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
	OtpCodeDuration time.Duration `koanf:"otp_code_duration"`
}

type Repository interface {
	GetUserByPhoneNumber(ctx context.Context, userID string) (*entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	GetOtpByUserID(ctx context.Context, userID uint) (*entity.Otp, error)
	UpdateOtpByUserID(ctx context.Context, otp entity.Otp) error
	CreateOtp(ctx context.Context, otp entity.Otp) (*entity.Otp, error)
}

func New(config Config, repository Repository) Service {
	return Service{
		repository: repository,
		config:     config,
	}
}
