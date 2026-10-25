package server

type Env string

const (
	ProductionEnv  Env = "production"
	DevelopmentEnv Env = "development"
)


func (e Env) isValid () bool {
	if e == ProductionEnv || e == DevelopmentEnv {
		return true 
	}
	return false 
}