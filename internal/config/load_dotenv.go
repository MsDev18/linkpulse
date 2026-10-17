package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
)

func (k KoanfConfig) LoadDotEnv(path string) {
	parser := dotenv.ParserEnv("", ".", func(s string) string {
		lower := strings.ToLower(s)
		return strings.ReplaceAll(lower, "__", ".")
	})

	if err := k.koanf.Load(file.Provider(path), parser); err != nil {
		log.Fatalf("cna't read data from %s file, err -> %v", path, err)
	}
}
