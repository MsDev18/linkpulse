package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	authclient "linkpulse/internal/client/auth"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	token := os.Getenv("ACCESS_TOKEN")
	if token == "" {
		return fmt.Errorf("ACCESS_TOKEN is required")
	}

	client, err := authclient.New(
		"127.0.0.1:50051",
		3*time.Second,
	)
	if err != nil {
		return err
	}
	defer client.Close()

	identity, err := client.ValidateAccessToken(
		context.Background(),
		token,
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"user_id=%d session_id=%d\n",
		identity.UserID,
		identity.SessionID,
	)

	return nil
}