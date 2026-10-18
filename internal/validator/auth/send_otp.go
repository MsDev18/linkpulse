package auth

import (
	"context"
	"errors"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/richerror"

	"github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) SendOtp(ctx context.Context, req dto.SendOtpRequest) error {
	const op = "auth-validator.SendOtp"

	err := validation.ValidateStructWithContext(
		ctx,
		&req,
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(PHONE_NUMBER_REGEX)),
	)

	if err == nil {
		return nil
	}

	var fieldErr *validation.Errors
	if errors.As(err, &fieldErr) {
		meta := make(map[string]any, len(*fieldErr))
		for field, fieldErr := range *fieldErr {
			meta[field] = fieldErr.Error()
		}
		return richerror.New(op).
			SetMessage("invalid input").
			SetKind(richerror.KindBadRequest).
			SetErr(err)
	}

	return richerror.
		New(op).
		SetMessage("unexpected error").
		SetKind(richerror.KindUnexpected).
		SetErr(err)
}
