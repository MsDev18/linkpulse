package main

import (
	"context"
	"fmt"
	"linkpulse/internal/config"
	authhandler "linkpulse/internal/delivery/grpc/auth"
	authpb "linkpulse/internal/pb/auth"
	"linkpulse/internal/repository/mysql"
	authrepository "linkpulse/internal/repository/mysql/auth"
	authservice "linkpulse/internal/service/auth"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// load configuration
	loader := config.New()
	loader.LoadYaml("config.yml")
	loader.LoadDotEnv(".env")
	cfg := loader.GetConfig()

	if cfg.Auth.AccessTokenSecret == "" {
		return fmt.Errorf("access token secret is required")
	}

	if cfg.Auth.RefreshTokenSecret == "" {
		return fmt.Errorf("refresh token secret is required")
	}

	if cfg.Auth.AccessTokenDuration <= 0 ||
		cfg.Auth.RefreshTokenDuration <= 0 ||
		cfg.Auth.OtpCodeDuration <= 0 {
		return fmt.Errorf("auth token and OTP durations must be positive")
	}

	mysqlRepo := mysql.New(cfg.MySQL)
	defer mysqlRepo.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := mysqlRepo.PingContext(ctx); err != nil {
		return fmt.Errorf("connect to mysql: %w", err)
	}

	repository := authrepository.New(mysqlRepo)
	service := authservice.New(cfg.Auth, repository)
	handler := authhandler.New(service)

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServer(grpcServer, &handler)

	reflection.Register(grpcServer)

	log.Printf("auth grpc listening on %s", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("serve grpc: %w", err)
	}
	return nil
}
