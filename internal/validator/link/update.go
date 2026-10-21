package link

import (
	"context"
	"errors"
	dto "linkpulse/internal/dto/link"
	"linkpulse/internal/pkg/richerror"
	"unicode/utf8"
)

func (v Validator) UpdateLink(ctx context.Context, req dto.UpdateRequest) error {
	const op = "link-validator.UpdateLink"
	if err := ctx.Err(); err != nil {
		return err
	}
	fields := make(map[string]any)
	if !req.OriginalURL.Set && !req.CustomCode.Set && !req.ExpiresAt.Set {
		fields["body"] = "at least one field is required"
	}
	if req.OriginalURL.Set {
		var err error
		if req.OriginalURL.Null || utf8.RuneCountInString(req.OriginalURL.Value) > 2048 {
			err = errors.New("a URL of at most 2048 characters is required")
		} else {
			err = v.validateOriginalURL(req.OriginalURL.Value)
		}
		if err != nil {
			fields["original-url"] = err.Error()
		}
	}
	if req.CustomCode.Set && (req.CustomCode.Null || !CODE_REGEX.MatchString(req.CustomCode.Value)) {
		fields["custom-code"] = "code must contain 3 to 32 letters, digits, underscores or hyphens"
	}
	if req.ExpiresAt.Set && !req.ExpiresAt.Null {
		if err := v.validateExpiresAt(&req.ExpiresAt.Value); err != nil {
			fields["expires-at"] = err.Error()
		}
	}
	if len(fields) > 0 {
		return richerror.New(op).SetMessage("invalid inputs").SetKind(richerror.KindBadRequest).SetMeta(fields)
	}
	return nil
}
