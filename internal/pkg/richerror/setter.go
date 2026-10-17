package richerror

func (r *RichErr) SetMessage(message string) *RichErr {
	r.message = message
	return r
}

func (r *RichErr) SetKind(kind Kind) *RichErr {
	r.kind = kind
	return r
}

func (r *RichErr) SetErr(err error) *RichErr {
	r.err = err 
	return r
}

func (r *RichErr) SetMeta(meta map[string]any) *RichErr {
	r.meta = meta
	return r
}