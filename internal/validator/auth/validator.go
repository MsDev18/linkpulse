package auth

import "regexp"

type Validator struct {

}

var (
	PHONE_NUMBER_REGEX = regexp.MustCompile(`^09[0-9]{9}$`)
)

func New () Validator {
	return Validator{}
}