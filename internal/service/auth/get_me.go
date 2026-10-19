package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/richerror"
)

func (s Service) GetMe(ctx context.Context, userID uint) (dto.GetMeResponse, error) {
	const op = "auth-service.GetMe"

	user, err := s.repository.GetUserByID(ctx, userID)
	if err != nil {
		return dto.GetMeResponse{}, err
	}

	if user == nil {
		return dto.GetMeResponse{}, richerror.New(op).
			SetMessage("not found user").
			SetKind(richerror.KindNotFound)
	}

	name := ""
	if user.Name != nil {
		name = *user.Name
	}
	

	return dto.GetMeResponse{
		ID:          user.ID,
		Name:        name,
		PhoneNumber: user.PhoneNumber,
	}, err
}
