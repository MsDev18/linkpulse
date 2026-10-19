package auth

import (
	"context"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/claims"
	"linkpulse/internal/pkg/richerror"
	"time"
)

func (s Service) CheckOtp(ctx context.Context, req dto.CheckOtpRequest) (dto.CheckOtpResponse, error) {
	const op = "auth-service.CheckOtp"

	user, err := s.repository.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return dto.CheckOtpResponse{}, err
	}

	if user == nil {
		return dto.CheckOtpResponse{}, richerror.New(op).
			SetMessage("not found user").
			SetKind(richerror.KindNotFound).
			SetErr(err)
	}

	otp, err := s.repository.GetOtpByUserID(ctx, user.ID)
	if err != nil {
		return dto.CheckOtpResponse{}, err
	}
	if otp == nil {
		return dto.CheckOtpResponse{}, richerror.New(op).
			SetMessage("not found otp").
			SetKind(richerror.KindNotFound).
			SetErr(err)
	}

	if time.Now().After(otp.ExpiresAt) {
		return dto.CheckOtpResponse{}, richerror.New(op).
			SetMessage("otp code expired").
			SetKind(richerror.KindUnauthorize).
			SetErr(err)
	}

	if otp.Code != req.Code {
		return dto.CheckOtpResponse{}, richerror.New(op).
			SetMessage("incorrect code").
			SetKind(richerror.KindUnauthorize)
	}

	sessionRecord := entity.Session{
		UserID: user.ID,
		ExpiresAt: time.Now().Add(s.config.RefreshTokenDuration),
	}
	session , err := s.repository.CreateSession(ctx, sessionRecord)
	if err != nil {
		return dto.CheckOtpResponse{},err
	}

	// jwt code 
	accessToken , err := claims.GenAccessToken(s.config.AccessTokenSecret , user.ID, session.ID , s.config.AccessTokenDuration)
	if err != nil {
		return dto.CheckOtpResponse{} , err
	}
	refreshToken, err := claims.GenRefreshToken(s.config.RefreshTokenSecret , user.ID, session.ID , s.config.RefreshTokenDuration) 
	if err != nil {
		return dto.CheckOtpResponse{}, err
	}

	return dto.CheckOtpResponse{
		Token : dto.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	} , nil
}
