package link

import (
	"context"
	"errors"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) UpdateLink(ctx context.Context, req dto.UpdateRequest) error {
	const op = "link-validator.UpdateLink"

	err := validation.ValidateStructWithContext(
		ctx,
		&req,
		validation.Field(&req.OriginalURL, validation.When(req.OriginalURL != "",
			validation.Length(1, 2048), validation.By(v.validateOriginalURL),
		)),
		validation.Field(&req.CustomCode, validation.Match(CODE_REGEX)),
		validation.Field(&req.ExpiresAt, validation.By(v.validateExpiresAt)),
	)
	if err == nil {
		return nil
	}

	var validationErr validation.Errors
	if errors.As(err, &validationErr) {
		meta := make(map[string]any, len(validationErr))
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
		SetMessage("unexpected err").
		SetKind(richerror.KindUnexpected).
		SetErr(err)
}
