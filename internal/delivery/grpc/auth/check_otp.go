package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	pb "linkpulse/internal/pb/auth"
	"linkpulse/internal/pkg/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) CheckOtp(ctx context.Context, req *pb.CheckOtpRequest) (*pb.CheckOtpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is required")
	}

	input := dto.CheckOtpRequest{
		PhoneNumber: req.GetPhoneNumber(),
		Code:        req.Code,
	}

	if err := h.validator.CheckOtp(ctx, input); err != nil {
		return nil, mapper.Error(err)
	}

	res, err := h.service.CheckOtp(ctx, input)
	if err != nil {
		return nil, mapper.Error(err)
	}

	return &pb.CheckOtpResponse{
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
	}, nil
}
