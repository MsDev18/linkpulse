package response

import (
	"errors"
	"linkpulse/internal/pkg/mapper"
	"linkpulse/internal/pkg/richerror"
)

func (r Responder) Error(err error) {
	var rErr *richerror.RichErr
	if errors.As(err, &rErr) {
		code := mapper.KindToHttpError(rErr.GetKind())
		message := rErr.GetMessage()
		r.c.AbortWithStatusJSON(code, Body{
			Code:    code,
			Message: message,
			Errors:  rErr.GetMeta(),
		})
		return 	
	}
	r.c.AbortWithStatusJSON(mapper.KindToHttpError(richerror.KindUnexpected), Body{
		Code:    mapper.KindToHttpError(richerror.KindUnexpected),
		Message: "for this package you should be use richerror",
	})
}
