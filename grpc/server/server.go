package server

import (
	"context"
	"errors"
	"net"

	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/grpc/training"
	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

type Server interface {
	Run(string) error

	RegisterTrainingServer(training.TrainingService) error
}

func NewServer() Server {
	return serverImpl{grpc.NewServer()}
}

type serverImpl struct {
	server *grpc.Server
}

func (impl serverImpl) RegisterTrainingServer(s training.TrainingService) error {
	if s == nil {
		return errors.New("invliad service")
	}

	if impl.server == nil {
		return errors.New("no server")
	}

	protocol.RegisterTrainingServer(impl.server, &trainingServer{s: s})

	return nil
}

func (impl serverImpl) Run(port string) error {
	if impl.server == nil {
		return errors.New("no server")
	}

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	defer interrupts.WaitForGracefulShutdown()

	interrupts.OnInterrupt(func() {
		logrus.Errorf("grpc server exit...")
		impl.server.Stop()
	})

	return impl.server.Serve(listen)
}

type trainingServer struct {
	s training.TrainingService
	protocol.UnimplementedTrainingServer
}

func (t *trainingServer) SetTrainingInfo(ctx context.Context, v *protocol.TrainingInfo) (
	*protocol.Result, error,
) {
	index := training.TrainingIndex{
		Id:        v.GetId(),
		User:      v.GetUser(),
		ProjectId: v.GetProjectId(),
	}

	info := training.TrainingInfo{
		OutputZipPath: v.GetOutputZipPath(),
		AimZipPath:    v.GetAimZipPath(),
		LogPath:       v.GetLogPath(),
		Status:        v.GetStatus(),
		Duration:      int(v.GetDuration()),
	}

	return nil, t.s.SetTrainingInfo(&index, &info)
}
