package mapper

import (
	"context"
	"errors"
	"linkpulse/internal/pkg/richerror"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.Canceled) {
		return status.Error(codes.Canceled, "request canceled")
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, "deadline exceeded")
	}

	var richErr *richerror.RichErr
	if errors.As(err, &richErr) {
		switch richErr.GetKind() {
		case richerror.KindBadRequest:
			return status.Error(codes.InvalidArgument, "invalid input")
		case richerror.KindNotFound:
			return status.Error(codes.NotFound, "resource not found")
		case richerror.KindConflict:
			return status.Error(codes.AlreadyExists, "resource already exists")
		case richerror.KindUnauthorize:
			return status.Error(codes.Unauthenticated, "authentication required")
		case richerror.KindForbidden:
			return status.Error(codes.PermissionDenied, "permission denied")
		case richerror.KindUnavailable:
			return status.Error(codes.Unavailable, "service unavailable")
		case richerror.KindTimeout:
			return status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
	}

	return status.Error(codes.Internal, "internal server error")
}