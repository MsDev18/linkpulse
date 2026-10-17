package main

import (
	"linkpulse/internal/config"
)

func LoadConfig() config.Config {
	k := config.New()
	k.LoadDotEnv(".env")
	k.LoadYaml("config.yml")
	cfg := k.GetConfig()
	return cfg
}
