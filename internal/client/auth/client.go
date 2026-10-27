package auth

import (
	"context"
	"fmt"
	pb "linkpulse/internal/pb/auth"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Identity struct {
	UserID    uint64
	SessionID uint64
}

type Client struct {
	conn    *grpc.ClientConn
	service pb.AuthClient
	timeout time.Duration
}

func New(address string, timeout time.Duration) (*Client, error) {
	if strings.TrimSpace(address) == "" {
		return nil, fmt.Errorf("auth address is required")
	}

	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be possitive")
	}

	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("create auth client err: %w", err)
	}

	return &Client{
		conn:    conn,
		service: pb.NewAuthClient(conn),
		timeout: timeout,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) ValidateAccessToken(ctx context.Context, token string) (Identity, error) {
	if strings.TrimSpace(token) == "" {
		return Identity{}, status.Error(
			codes.Unauthenticated,
			"access token is required",
		)
	}

	callCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	callCtx = metadata.NewOutgoingContext(
		callCtx,
		metadata.Pairs(
			"authorization",
			"Bearer "+token,
		))

	res, err := c.service.ValidateAccessToken(callCtx, &pb.ValidateAccessTokenRequest{})
	if err != nil {
		return Identity{}, err
	}

	if res.GetUserId() == 0 || res.GetSessionId() == 0 {
		return Identity{}, status.Error(
			codes.Internal,
			"invalid identity returned by auth server",
		)
	}

	return Identity{
		UserID:    res.GetUserId(),
		SessionID: res.GetSessionId(),
	}, nil
}
