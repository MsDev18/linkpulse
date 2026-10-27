package auth

import (
	"context"
	pb "linkpulse/internal/pb/auth"
	"linkpulse/internal/pkg/mapper"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h Handler) ValidateAccessToken(ctx context.Context, _ *pb.ValidateAccessTokenRequest) (*pb.ValidateAccessTokenResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "authorization metadata required")
	}

	values := md.Get("authorization")
	if len(values) != 1 {
		return nil, status.Error(
			codes.Unauthenticated,
			"exactly one authorization value required",
		)
	}

	parts := strings.Fields(values[0])
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, status.Error(
			codes.Unauthenticated,
			"invalid authorization format",
		)
	}

	res, err := h.service.ValidateAccessToken(ctx , parts[1])
	if err != nil {
		return nil , mapper.Error(err)
	}

	return &pb.ValidateAccessTokenResponse{
		UserId:    uint64(res.UserID),
		SessionId: uint64(res.SessionID),
	} , nil
}
