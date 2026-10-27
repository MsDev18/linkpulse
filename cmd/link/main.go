package main

import (
	"context"
	"fmt"
	"linkpulse/internal/client/auth"
	"linkpulse/internal/config"
	linkhandler "linkpulse/internal/delivery/grpc/link"
	linkpb "linkpulse/internal/pb/link"
	"linkpulse/internal/repository/mysql"
	linkrepository "linkpulse/internal/repository/mysql/link"
	linkservice "linkpulse/internal/service/link"
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
	loader := config.New()
	loader.LoadYaml("config.yml")
	loader.LoadDotEnv(".env")
	cfg := loader.GetConfig()

	if cfg.Link.PublicBaseURL == "" {
		return fmt.Errorf("link public base URL is required")
	}

	mysqlRepo := mysql.New(cfg.MySQL)
	defer mysqlRepo.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err := mysqlRepo.PingContext(ctx)
	cancel()

	if err != nil {
		return fmt.Errorf("connect to mysql: %w", err)
	}

	auth, err := auth.New("0.0.0.0:50051", time.Second*3)
	if err != nil {
		return fmt.Errorf("create auth client err: %w", err)
	}
	defer auth.Close()

	repo := linkrepository.New(mysqlRepo)
	service := linkservice.New(cfg.Link, repo)
	handler := linkhandler.New(service, auth)

	listener, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		return fmt.Errorf("listen err: %w", err)
	}
	defer listener.Close()

	server := grpc.NewServer()
	linkpb.RegisterLinkServer(server, &handler)
	reflection.Register(server)

	log.Print("link grpc listen on 0.0.0.0:50052")
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("serve grpc: %w", err)
	}
	return nil
}
