package richerror

type Kind uint8

const (
	KindUnknown Kind = iota + 1
	KindInvalid
	KindNotFound
	KindConflict
	KindUnauthenticated
	KindForbidden
	KindInternal
	KindUnavailable
	KindTimeout
)