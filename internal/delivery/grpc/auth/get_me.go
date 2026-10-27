package auth

import (
	"context"
	pb "linkpulse/internal/pb/auth"
	"linkpulse/internal/pkg/mapper"
)

func (h Handler) GetMe(ctx context.Context, _ *pb.GetMeRequest) (*pb.GetMeResponse, error) {
	identity, err := h.ValidateAccessToken(
		ctx,
		&pb.ValidateAccessTokenRequest{},
	)
	if err != nil {
		return nil, err
	}

	res, err := h.service.GetMe(ctx, uint(identity.UserId))
	if err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.GetMeResponse{
		Id:          uint64(res.ID),
		Name:        res.Name,
		PhoneNumber: res.PhoneNumber,
	}, nil
}
