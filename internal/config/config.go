package config

import "github.com/knadh/koanf/v2"

type KoanfConfig struct {
	koanf *koanf.Koanf
}

type Config struct {

}

func New() KoanfConfig {
	return KoanfConfig{
		koanf: koanf.New("."),
	}
}


