package link

import (
	"context"
	pb "linkpulse/internal/pb/link"
	"linkpulse/internal/pkg/mapper"
)

func (h Handler) GetAll(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	userID, err := h.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.service.GetAll(ctx, userID)
	if err != nil {
		return nil, mapper.Error(err)
	}

	links := make([]*pb.GetOneResponse, 0, len(res))

	for _, r := range res {
		item, err := toProtoResponse(r)
		if err != nil {
			return nil, err
		}

		links = append(links, item)
	}

	return &pb.GetAllResponse{
		Links: links,
	}, nil
}
