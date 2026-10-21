package router

func (r Router) RegisterLinkRoute() {
	linkG := r.e.Group("link")
	linkG.POST("create", r.authMiddleware.Auth(), r.link.Create)
}
