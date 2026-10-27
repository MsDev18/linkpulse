package link

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h Handler) Authenticate(ctx context.Context) (uint, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Error(
			codes.Unauthenticated,
			"authorization metadata require",
		)
	}

	values := md.Get("Authorization")
	if len(values) != 1 {
		return 0, status.Error(
			codes.Unauthenticated,
			"exactly one authorization value required",
		)
	}

	parts := strings.Fields(values[0])
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return 0, status.Error(
			codes.Unauthenticated,
			"invalid authorization format",
		)
	}

	identity, err := h.auth.ValidateAccessToken(ctx, parts[1])
	if err != nil {
		return 0, err
	}

	return uint(identity.UserID), nil
}
