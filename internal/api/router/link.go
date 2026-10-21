package router

func (r Router) RegisterLinkRoute() {
	linkG := r.e.Group("link")
	linkG.POST("", r.authMiddleware.Auth(), r.link.Create)
	linkG.GET("", r.authMiddleware.Auth(), r.link.GetAll)
	linkG.DELETE(":id" , r.authMiddleware.Auth() , r.link.Delete)
	linkG.GET(":id", r.link.GetOneByID)
}
