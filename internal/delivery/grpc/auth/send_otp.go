package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	pb "linkpulse/internal/pb/auth"
	"linkpulse/internal/pkg/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) SendOtp(ctx context.Context, req *pb.SendOtpRequest) (*pb.SendOtpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is required")
	}

	input := dto.SendOtpRequest{
		PhoneNumber: req.GetPhoneNumber(),
	}

	if err := h.validator.SendOtp(ctx, input); err != nil {
		return nil, mapper.Error(err)
	}

	_, err := h.service.SendOtp(ctx, input)
	if err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.SendOtpResponse{}, nil
}
