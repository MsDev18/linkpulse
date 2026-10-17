package config

import "log"

func (k KoanfConfig) GetConfig () Config {
	var cfg Config 

	if err := k.koanf.Unmarshal("" , &cfg) ; err != nil {
		log.Fatalf("can't unmarshal data, err -> %v", err)
	}
	return cfg
}