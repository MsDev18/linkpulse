package router

func (r Router) RegisterHealthRoute() {
	healthG := r.e.Group("/health")
	healthG.GET("" , r.health.Health)
}
