package server

import (
	"fmt"
	health "linkpulse/internal/api/handler/health"
	auth "linkpulse/internal/api/handler/auth"
	authmiddleware "linkpulse/internal/api/middleware/auth"
	"linkpulse/internal/api/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host         string        `koanf:"host"`
	Port         int           `koanf:"port"`
	ReadTimeout  time.Duration `koanf:"read_timeout"`
	WriteTimeout time.Duration `koanf:"write_timeout"`
	IdleTimeout  time.Duration `koanf:"idle_timeout"`
	Env          Env
}

func (c Config) isValid() bool {
	if c.IdleTimeout == 0 || c.ReadTimeout == 0 || c.WriteTimeout == 0 {
		return false
	}
	return true
}

type Server struct {
	httpServer *http.Server
}

func New(config Config, health health.Handler, auth auth.Hander, authMiddleware authmiddleware.Middleware) Server {
	// new instance from gin engine
	engine := gin.Default()
	// validation configs
	if !config.isValid() {
		log.Fatalf("please implement configurations")
	}

	// validation config.env
	env := Env(config.Env)
	if !env.isValid() {
		log.Fatalf("please enter correct env config")
	}
	switch env {
	case ProductionEnv:
		gin.SetMode(gin.ReleaseMode)
	case DevelopmentEnv:
		gin.SetMode(gin.DebugMode)
	}

	// routre package
	appRouter := router.New(engine, health, auth , authMiddleware)
	appRouter.RegisterRoute()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:      engine,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return Server{
		httpServer: httpServer,
	}
}

func (s Server) Run() {
	if err := s.httpServer.ListenAndServe() ; err != nil {
		log.Fatalf("can't run server err: %v" , err)
	}
}
