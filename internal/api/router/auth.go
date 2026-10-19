package router

func (r Router) RegisterAuthRoute() {
	authG := r.e.Group("auth")
	authG.POST("send-otp", r.auth.SendOtp)
	authG.POST("check-otp" , r.auth.CheckOtp)
	authG.GET("get-me" , r.authMiddleware.Auth() , r.auth.GetMe)
	authG.POST("refresh-token" , r.auth.RefreshToken)
	authG.POST("logout", r.authMiddleware.Auth(), r.auth.Logout)
}