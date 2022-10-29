package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/grpc/training"
	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		cli:  protocol.NewTrainingClient(conn),
	}, nil
}

type Client struct {
	conn *grpc.ClientConn
	cli  protocol.TrainingClient
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) SetTrainingInfo(index *training.TrainingIndex, info *training.TrainingInfo) error {
	_, err := c.cli.SetTrainingInfo(
		context.Background(),
		&protocol.TrainingInfo{
			Id:            index.Id,
			User:          index.User,
			Status:        info.Status,
			LogPath:       info.LogPath,
			Duration:      int32(info.Duration),
			ProjectId:     index.ProjectId,
			AimZipPath:    info.AimZipPath,
			OutputZipPath: info.OutputZipPath,
		},
	)

	return err
}
