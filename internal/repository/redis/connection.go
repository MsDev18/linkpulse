package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Redis struct {
	Conn *goredis.Client
}

type Config struct {
	Host         string        `koanf:"host"`
	Port         uint          `koanf:"port"`
	DialTimeout  time.Duration `koanf:"dial_timeout"`
	ReadTimeout  time.Duration `koanf:"read_timeout"`
	WriteTimeout time.Duration `koanf:"write_timeout"`
}


func New(config Config) Redis {
	client := goredis.NewClient(&goredis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	})

	ctx, cancel := context.WithTimeout(
		context.Background(), 
		config.DialTimeout ,
	)
	defer cancel()

	if err := client.Ping(ctx).Err() ; err != nil {
		_ = client.Close()
		log.Fatalf("connect to redis: %v", err)
	}

	return Redis{
		Conn: client,
	}
}
