package link

import (
	"context"
	pb "linkpulse/internal/pb/link"
	"linkpulse/internal/pkg/mapper"
	validator "linkpulse/internal/validator/link"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) Resolve(ctx context.Context, req *pb.ResolveRequest) (*pb.ResolveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is required")
	}

	code := req.GetCode()
	if !validator.CODE_REGEX.MatchString(code) {
		return nil, status.Error(codes.InvalidArgument, "code must contain 3 to 32 letters, digits, underscores or hyphens")
	}

	originalURL, err := h.service.Resolve(ctx, code)
	if err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.ResolveResponse{OriginalUrl: originalURL}, nil
}
