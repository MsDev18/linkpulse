package auth

import (
	"context"
	pb "linkpulse/internal/pb/auth"
	"linkpulse/internal/pkg/mapper"
)

func (h Handler) Logout(ctx context.Context, _ *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	identity, err := h.ValidateAccessToken(ctx, &pb.ValidateAccessTokenRequest{})
	if err != nil {
		return nil, err
	}

	if err := h.service.Logout(ctx, uint(identity.SessionId)); err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.LogoutResponse{}, nil

}
