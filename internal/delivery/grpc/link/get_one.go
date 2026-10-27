package link

import (
	"context"
	pb "linkpulse/internal/pb/link"
	"linkpulse/internal/pkg/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) GetOne(ctx context.Context, req *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	if req == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"request is required",
		)
	}

	id := req.GetId()

	result, err := h.service.GetOneByID(ctx, uint(id))
	if err != nil {
		return nil, mapper.Error(err)
	}

	return toProtoResponse(result)
}
