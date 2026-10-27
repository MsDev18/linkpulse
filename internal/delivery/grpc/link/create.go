package link

import (
	"context"
	dto "linkpulse/internal/dto/link"
	pb "linkpulse/internal/pb/link"
	"linkpulse/internal/pkg/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.GetOneResponse, error) {
	if req == nil {
		return nil, status.Error(
			codes.InvalidArgument , 
			"request is required" ,
		)
	}

	userID, err := h.Authenticate(ctx)
	if err != nil {
		return nil , err
	}

	input := dto.CreateRequest{
		OriginalURL: req.OriginalUrl,
		CustomCode:  req.CustomCode,
	}

	if req.ExpiresAt != nil {
		if err := req.ExpiresAt.CheckValid() ; err != nil {
			return nil , status.Error(
				codes.InvalidArgument , 
				"invalid expiration timestamp" ,
			)
		}

		expiresAt := req.ExpiresAt.AsTime()
		input.ExpiresAt = &expiresAt
	}

	if err := h.validator.CreateLink(ctx , input) ; err != nil {
		return nil , mapper.Error(err)
	}

	result , err := h.service.Create(ctx , userID , input)
	if err != nil {
		return nil , mapper.Error(err)
	}

	return toProtoResponse(result)
}