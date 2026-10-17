package richerror

import "fmt"

type Op string

type RichErr struct {
	op      Op
	message string
	kind    Kind
	err     error
	meta    map[string]any
}

func New(op Op) *RichErr {
	return &RichErr{
		op: op,
	}
}

func (r *RichErr) Error() string {
	err := fmt.Sprintf("op: %s \nmessage: %s \nerr: %v", r.op, r.message, r.err)
	return err
}

func (r *RichErr) Unwrap() error {
	return r.err
}