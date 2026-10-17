package main

import "linkpulse/internal/pkg/migrator"

func main () {
	config := LoadConfig()

	m := migrator.New(config.MySQL)
	defer m.Close()
	m.Up()
}