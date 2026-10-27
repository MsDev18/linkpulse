package link

import (
	"context"

	dto "linkpulse/internal/dto/link"
	pb "linkpulse/internal/pb/link"
	"linkpulse/internal/pkg/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) Update(ctx context.Context,req *pb.UpdateRequest) (*pb.GetOneResponse, error) {
	if req == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"request is required",
		)
	}

	userID, err := h.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	id := req.GetId()

	if req.OriginalUrl == nil &&
		req.CustomCode == nil &&
		req.ExpiresAt == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"at least one update field is required",
		)
	}

	if req.OriginalUrl != nil && req.GetOriginalUrl() == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"original URL must not be empty",
		)
	}

	if req.CustomCode != nil && req.GetCustomCode() == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"custom code must not be empty",
		)
	}

	input := dto.UpdateRequest{
		OriginalURL: req.GetOriginalUrl(),
		CustomCode:  req.GetCustomCode(),
	}

	if req.ExpiresAt != nil {
		if err := req.ExpiresAt.CheckValid(); err != nil {
			return nil, status.Error(
				codes.InvalidArgument,
				"invalid expiration timestamp",
			)
		}

		expiresAt := req.ExpiresAt.AsTime()
		input.ExpiresAt = &expiresAt
	}

	if err := h.validator.UpdateLink(ctx, input); err != nil {
		return nil, mapper.Error(err)
	}

	result, err := h.service.Update(
		ctx,
		userID,
		uint(id),
		input,
	)
	if err != nil {
		return nil, mapper.Error(err)
	}

	return toProtoResponse(result)
}