package server

import (
	"context"
	"net"

	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

func Start(port string, s TrainingService) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	protocol.RegisterTrainingServer(server, &trainingServer{s: s})

	return run(server, listen)
}

func run(server *grpc.Server, listen net.Listener) error {
	defer interrupts.WaitForGracefulShutdown()

	interrupts.OnInterrupt(func() {
		logrus.Errorf("grpc server exit...")
		server.Stop()
	})

	return server.Serve(listen)
}

type trainingServer struct {
	s TrainingService
	protocol.UnimplementedTrainingServer
}

func (t *trainingServer) SetTrainingInfo(ctx context.Context, v *protocol.TrainingInfo) (
	*protocol.Result, error,
) {
	index := TrainingIndex{
		Id:        v.GetId(),
		User:      v.GetUser(),
		ProjectId: v.GetProjectId(),
	}

	info := TrainingInfo{
		OutputZipPath: v.GetOutputZipPath(),
		AimZipPath:    v.GetAimZipPath(),
		LogPath:       v.GetLogPath(),
		Status:        v.GetStatus(),
		Duration:      int(v.GetDuration()),
	}

	return nil, t.s.SetTrainingInfo(&index, &info)
}
