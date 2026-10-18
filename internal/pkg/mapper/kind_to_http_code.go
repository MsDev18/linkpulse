package mapper

import (
	"net/http"

	"linkpulse/internal/pkg/richerror"
)

// KindToHttpError maps an error kind to its HTTP status code.
// Unknown and unset kinds are treated as internal server errors.
func KindToHttpError(kind richerror.Kind) int {
	switch kind {
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	case richerror.KindBadRequest:
		return http.StatusBadRequest
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindConflict:
		return http.StatusConflict
	case richerror.KindUnauthorize:
		return http.StatusUnauthorized
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindInternal:
		return http.StatusInternalServerError
	case richerror.KindUnavailable:
		return http.StatusServiceUnavailable
	case richerror.KindTimeout:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}
