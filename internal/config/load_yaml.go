package config

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func (k KoanfConfig) LoadYaml(path string) {
	if err := k.koanf.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Fatalf("can't read data from %s file, err -> %v", path, err)
	}
}
