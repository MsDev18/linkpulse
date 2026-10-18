package router

func (r Router) RegisterAuthRoute() {
	authG := r.e.Group("auth")
	authG.POST("send-otp", r.auth.SendOtp)
}