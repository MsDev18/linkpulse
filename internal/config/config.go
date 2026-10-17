package config

import (
	"linkpulse/internal/repository/mysql"

	"github.com/knadh/koanf/v2"
)

type KoanfConfig struct {
	koanf *koanf.Koanf
}

type Config struct {
	MySQL mysql.Config `koanf:"mysql"`
}

func New() KoanfConfig {
	return KoanfConfig{
		koanf: koanf.New("."),
	}
}