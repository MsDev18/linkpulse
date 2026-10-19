package auth

import (
	"context"
	"errors"
	dto "linkpulse/internal/dto/auth"
	"linkpulse/internal/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) CheckOtp(ctx context.Context, req dto.CheckOtpRequest) error {
	const op = "auth-validator.CheckOtp"

	err := validation.ValidateStructWithContext(
		ctx,
		&req,
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(PHONE_NUMBER_REGEX)),
		validation.Field(&req.Code, validation.Required, validation.Length(5, 5)),
	)

	if err == nil {
		return nil
	}

	var validationErr validation.Errors
	if errors.As(err, &validationErr) {
		meta := make(map[string]any)
		for field, fieldErr := range validationErr {
			meta[field] = fieldErr.Error()
		}
		return richerror.New(op).
			SetMessage("invalid inputs").
			SetKind(richerror.KindBadRequest).
			SetErr(err).
			SetMeta(meta)
	}
	return richerror.New(op).
		SetMessage("unexpected err in validation").
		SetKind(richerror.KindUnexpected).
		SetErr(err)
}
