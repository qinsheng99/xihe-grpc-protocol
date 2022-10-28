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

func (t *trainingServer) toTrainingInfo(v *protocol.TrainingInfo) TrainingInfo {
	return TrainingInfo{
		Id:        v.Id,
		User:      v.User,
		ProjectId: v.ProjectId,
	}
}

func (t *trainingServer) SetTrainingStatus(ctx context.Context, v *protocol.TrainingStatus) (
	*protocol.Result, error,
) {
	info := t.toTrainingInfo(v.Info)

	return nil, t.s.SetTrainingStatus(&info, v.Status)
}

func (t *trainingServer) SetTrainingOutput(ctx context.Context, v *protocol.TrainingOutput) (
	*protocol.Result, error,
) {
	info := t.toTrainingInfo(v.Info)

	output := TrainingOutput{
		OutputZipPath: v.OutputZipPath,
		AimPath:       v.AimPath,
	}

	return nil, t.s.SetTrainingOutput(&info, &output)
}

func (t *trainingServer) SetTrainingLogPath(ctx context.Context, v *protocol.TrainingLogPath) (
	*protocol.Result, error,
) {
	info := t.toTrainingInfo(v.Info)

	return nil, t.s.SetTrainingLogPath(&info, v.Path)
}
