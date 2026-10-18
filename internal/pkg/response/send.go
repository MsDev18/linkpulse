package response


func (r Responder) Send(code int , message string, data any) {
	r.c.JSON(code , Body{
		Code: code,
		Message: message,
		Data: data,
	})
}