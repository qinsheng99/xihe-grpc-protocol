package client

import (
	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:           conn,
		TrainingClient: protocol.NewTrainingClient(conn),
	}, nil
}

type Client struct {
	protocol.TrainingClient

	conn *grpc.ClientConn
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}
