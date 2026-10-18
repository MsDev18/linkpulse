package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/entity"
	"linkpulse/internal/pkg/richerror"
	"math/big"
	"time"

)

func (s Service) SendOtp(ctx context.Context, req dto.SendOtpRequest) (dto.SendOtpResponse, error) {
	const op = "auth-service.SendOtp"

	// get user by id
	user, err := s.getOrCreateUser(ctx, req.PhoneNumber)
	if err != nil {
		return dto.SendOtpResponse{}, err
	}

	otp, err := s.updateOrCreateOtp(ctx, user.ID)
	if err != nil {
		return dto.SendOtpResponse{}, err
	}

	fmt.Println(otp.Code)

	return dto.SendOtpResponse{}, nil
}

func (s Service) getOrCreateUser(ctx context.Context, phoneNumber string) (*entity.User, error) {
	const op = "auth-service.getOrCreateUser"

	var user *entity.User
	var err error

	user, err = s.repository.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	record := entity.User{
		PhoneNumber: phoneNumber,
	}
	user, err = s.repository.CreateUser(ctx, record)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s Service) updateOrCreateOtp(ctx context.Context, userID uint) (*entity.Otp, error) {
	const op = "auth-service.updateOrCreateOtp"

	var otp *entity.Otp
	var err error

	otp, err = s.repository.GetOtpByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	code, err := s.generateOtpCode()
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(s.config.OtpCodeDuration)
	
	if otp != nil {
		otp.Code = code
		otp.ExpiresAt = expiresAt
		err := s.repository.UpdateOtpByUserID(ctx, *otp)
		if err != nil {
			return nil, err
		}
		return otp, err
	}

	otp = &entity.Otp{
		UserID:    userID,
		Code:      code,
		ExpiresAt: expiresAt,
	}
	createdOtp, err := s.repository.CreateOtp(ctx, *otp)
	if err != nil {
		return nil, err
	}
	return createdOtp, nil

}

func (s Service) generateOtpCode() (string, error) {
	const op = "auth-service.generateOtpCode"

	n, err := rand.Int(rand.Reader, big.NewInt(90000))
	if err != nil {
		return "", richerror.New(op).
			SetMessage("unexpected error in generate otp code").
			SetKind(richerror.KindUnexpected).
			SetErr(err)
	}

	return fmt.Sprintf("%d", n.Int64()+10000), nil
}
