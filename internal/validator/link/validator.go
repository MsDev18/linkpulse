package link

import "regexp"

type Validator struct{}

var CODE_REGEX = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,32}$`)

func New() Validator {
	return Validator{}
}

