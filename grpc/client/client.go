package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/grpc/inference"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/training"
	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

func NewTrainingClient(endpoint string) (*TrainingClient, error) {
	c, err := newConn(endpoint)
	if err != nil {
		return nil, err
	}

	return &TrainingClient{
		clientConn: &c,
		cli:        protocol.NewTrainingClient(c.conn),
	}, nil
}

func NewInferenceClient(endpoint string) (*InferenceClient, error) {
	c, err := newConn(endpoint)
	if err != nil {
		return nil, err
	}

	return &InferenceClient{
		clientConn: &c,
		cli:        protocol.NewInferenceClient(c.conn),
	}, nil
}

func newConn(endpoint string) (c clientConn, err error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err == nil {
		c.conn = conn
	}

	return
}

type clientConn struct {
	conn *grpc.ClientConn
}

func (c *clientConn) Disconnect() error {
	if c == nil || c.conn == nil {
		return nil
	}

	err := c.conn.Close()
	c.conn = nil

	return err
}

type TrainingClient struct {
	*clientConn

	cli protocol.TrainingClient
}

func (c *TrainingClient) SetTrainingInfo(index *training.TrainingIndex, info *training.TrainingInfo) error {
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

type InferenceClient struct {
	*clientConn

	cli protocol.InferenceClient
}

func (c *InferenceClient) SetInferenceInfo(index *inference.InferenceIndex, info *inference.InferenceInfo) error {
	_, err := c.cli.SetInferenceInfo(
		context.Background(),
		&protocol.InferenceInfo{
			Id:         index.Id,
			User:       index.User,
			ProjectId:  index.ProjectId,
			LastCommit: index.LastCommit,
			Error:      info.Error,
			AccessUrl:  info.AccessURL,
		},
	)

	return err
}
