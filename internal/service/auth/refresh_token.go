package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/claims"
)

func (s Service) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	const op = "auth-service.RefreshToken"
	
	claim, err := claims.ParseRefreshToken(s.config.RefreshTokenSecret, req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	accessToken, err := claims.GenAccessToken(s.config.AccessTokenSecret, claim.UserID, s.config.AccessTokenDuration)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	return dto.RefreshTokenResponse{
		Token: dto.Token{
			AccessToken: accessToken,
		},
	}, nil
}
