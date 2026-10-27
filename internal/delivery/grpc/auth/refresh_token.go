package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	pb "linkpulse/internal/pb/auth"
	"linkpulse/internal/pkg/mapper"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	if req == nil || strings.TrimSpace(req.GetRefreshToken()) == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"refresh token is required",
		)
	}

	result, err := h.service.RefreshToken(
		ctx,
		dto.RefreshTokenRequest{
			RefreshToken: req.GetRefreshToken(),
		},
	)
	if err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.RefreshTokenResponse{
		AccessToken: result.Token.AccessToken,
	}, nil
}
