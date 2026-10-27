package router

func (r Router) RegisterRedirectRoute() {
	r.e.GET("/:code", r.link.Redirect)
}
