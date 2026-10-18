package config

import (
	"linkpulse/internal/api/server"
	"linkpulse/internal/repository/mysql"
	authservice "linkpulse/internal/service/auth"

	"github.com/knadh/koanf/v2"
)

type KoanfConfig struct {
	koanf *koanf.Koanf
}

type Config struct {
	MySQL  mysql.Config        `koanf:"mysql"`
	Server server.Config       `koanf:"server"`
	Auth   authservice.Config `koanf:"auth"`
}

func New() KoanfConfig {
	return KoanfConfig{
		koanf: koanf.New("."),
	}
}
