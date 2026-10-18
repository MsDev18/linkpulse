package richerror

type Kind uint8

const (
	KindUnexpected Kind = iota + 1
	KindBadRequest
	KindNotFound
	KindConflict
	KindUnauthorize
	KindForbidden
	KindInternal
	KindUnavailable
	KindTimeout
)