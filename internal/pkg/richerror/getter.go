package richerror

import "errors"

func (r *RichErr) GetMessage() string {
	if r.message != "" {
		return r.message
	}

	if r.err == nil {
		return ""
	}

	var rErr *RichErr
	if errors.As(r.err, &rErr) {
		return rErr.GetMessage()
	}

	return r.err.Error()
}

func (r *RichErr) GetKind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	if r.err == nil {
		return 0
	}

	var rErr *RichErr
	if errors.As(r.err , &rErr) {
		return rErr.GetKind()
	}
	return 0
}

func (r *RichErr) GetMeta() map[string]any {
	return r.meta
}
