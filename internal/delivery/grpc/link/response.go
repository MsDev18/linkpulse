package link

import (
	dto "linkpulse/internal/dto/link"
	pb "linkpulse/internal/pb/link"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProtoResponse(result dto.GetOneResponse) (*pb.GetOneResponse, error) {
	createdAt := timestamppb.New(result.CreatedAt)
	if err := createdAt.CheckValid(); err != nil {
		return nil, status.Error(
			codes.Internal,
			"invalid creation timestamp",
		)
	}

	response := &pb.GetOneResponse{
		Id:          uint64(result.ID),
		OriginalUrl: result.OriginalURL,
		Code:        result.Code,
		ShortUrl:    result.ShortURL,
		CreatedAt:   createdAt,
	}

	if result.ExpiresAt != nil {
		expiresAt := timestamppb.New(*result.ExpiresAt)
		if err := expiresAt.CheckValid(); err != nil {
			return nil, status.Error(
				codes.Internal,
				"invalid expiration time",
			)
		}
		response.ExpiresAt = expiresAt
	}

	return response, nil
}
