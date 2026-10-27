package link

import (
	"context"

	authclient "linkpulse/internal/client/auth"
	dto "linkpulse/internal/dto/link"
	pb "linkpulse/internal/pb/link"
	validator "linkpulse/internal/validator/link"
)

type Service interface {
	Create(ctx context.Context, userID uint, req dto.CreateRequest) (dto.GetOneResponse, error)
	GetAll(ctx context.Context, userID uint) ([]dto.GetOneResponse, error)
	GetOneByID(ctx context.Context, id uint) (dto.GetOneResponse, error)
	Resolve(ctx context.Context, code string) (string, error)
	Update(ctx context.Context, userID uint, linkID uint, req dto.UpdateRequest) (dto.GetOneResponse, error)
	DeleteByID(ctx context.Context, userID uint, linkID uint) error
}

type AuthClient interface {
	ValidateAccessToken(ctx context.Context, token string) (authclient.Identity, error)
}

type Handler struct {
	pb.UnimplementedLinkServer
	service   Service
	auth      AuthClient
	validator validator.Validator
}

func New(service Service, auth AuthClient) Handler {
	return Handler{
		service:   service,
		auth:      auth,
		validator: validator.New(),
	}
}
