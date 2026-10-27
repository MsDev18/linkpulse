package link

import (
	"context"

	pb "linkpulse/internal/pb/link"
	"linkpulse/internal/pkg/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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

	if err := h.service.DeleteByID(ctx, userID, uint(id)); err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.DeleteResponse{}, nil
}
