package link

import (
	"context"
	"errors"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/richerror"
	"net/url"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) CreateLink(ctx context.Context, req dto.CreateRequest) error {
	const op = "link-validator.CreateLink"

	err := validation.ValidateStructWithContext(
		ctx,
		&req,
		validation.Field(&req.OriginalURL, validation.Required, validation.Length(1, 2048), validation.By(v.validateOriginalURL)),
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

func (v Validator) validateOriginalURL(value any) error {
	rawURL, ok := value.(string)
	if !ok {
		return errors.New("invalid URL")
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return errors.New("invalid URL")
	}

	if !strings.EqualFold(parsed.Scheme, "http") &&
		!strings.EqualFold(parsed.Scheme, "https") {
		return errors.New("URL scheme must be http or https")
	}

	if parsed.Hostname() == "" {
		return errors.New("URL must have a host")
	}

	if parsed.User != nil {
		return errors.New("URL must not contain credentials")
	}

	return nil
}


func (v Validator) validateExpiresAt(value any) error {
	expiresAt, ok := value.(*time.Time)
	if !ok {
		return errors.New("invalid expiration time")
	}

	if expiresAt == nil {
		return nil
	}

	if !expiresAt.After(time.Now()) {
		return errors.New("expiration time must be future")
	}

	return nil
}