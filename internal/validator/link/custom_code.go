package link

import "errors"

func (v Validator) validateCustomCode(value any) error {
	code, ok := value.(string)
	if !ok {
		return errors.New("invalid custom code")
	}

	switch code {
	case "auth", "health", "link":
		return errors.New("custom code is reserved for API routes")
	}

	return nil
}
