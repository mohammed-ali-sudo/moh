package execs

// internal/adapters/grpc/execs_client.go
import (
	"context"
	"time"

	gapi "moh/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	api  gapi.ExecsServiceClient
}

// New: uses grpc.NewClient (lazy connect). Add Connect() below if you want to block.
func New(addr string) (*Client, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn, api: gapi.NewExecsServiceClient(conn)}, nil
}

func (c *Client) Close() error { return c.conn.Close() }

// Optional: call once at startup if you want to actively connect instead of lazy.
func (c *Client) Connect() { c.conn.Connect() }

// RPC helpers with short per-call timeouts.
func (c *Client) AddDrug(ctx context.Context, d *gapi.DrugOut) (*gapi.ConfirmMessage, error) {
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.api.AddDrug(cctx, d)
}
func (c *Client) UpdateDrug(ctx context.Context, d *gapi.DrugOut) (*gapi.ConfirmMessage, error) {
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return c.api.UpdateDrug(cctx, d)
}
