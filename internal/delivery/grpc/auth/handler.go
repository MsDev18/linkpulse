package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	pb "linkpulse/internal/pb/auth"
	validator "linkpulse/internal/validator/auth"
)

type Service interface {
	ValidateAccessToken(ctx context.Context, token string) (dto.ValidateAccessTokenResponse, error)
	SendOtp(ctx context.Context, req dto.SendOtpRequest) (dto.SendOtpResponse, error)
	CheckOtp(ctx context.Context, req dto.CheckOtpRequest) (dto.CheckOtpResponse, error)
	GetMe(ctx context.Context, userID uint) (dto.GetMeResponse, error)
	Logout(ctx context.Context, sessionID uint) error
	RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)
}

type Handler struct {
	pb.UnimplementedAuthServer
	service   Service
	validator validator.Validator
}

func New(service Service) Handler {
	return Handler{
		service:   service,
		validator: validator.New(),
	}
}
