package server

import (
	"context"
	"errors"
	"net"

	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/opensourceways/xihe-grpc-protocol/grpc/competition"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/evaluate"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/finetune"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/inference"
	"github.com/opensourceways/xihe-grpc-protocol/grpc/training"
	"github.com/opensourceways/xihe-grpc-protocol/protocol"
)

type Server interface {
	Run(string) error

	RegisterFinetuneServer(finetune.FinetuneService) error
	RegisterTrainingServer(training.TrainingService) error
	RegisterEvaluateServer(evaluate.EvaluateService) error
	RegisterInferenceServer(inference.InferenceService) error
	RegisterCompetitionServer(competition.CompetitionService) error
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

func (impl serverImpl) RegisterFinetuneServer(s finetune.FinetuneService) error {
	if s == nil {
		return errors.New("invliad service")
	}

	if impl.server == nil {
		return errors.New("no server")
	}

	protocol.RegisterFinetuneServer(impl.server, &finetuneServer{s: s})

	return nil
}

func (impl serverImpl) RegisterInferenceServer(s inference.InferenceService) error {
	if s == nil {
		return errors.New("invliad service")
	}

	if impl.server == nil {
		return errors.New("no server")
	}

	protocol.RegisterInferenceServer(impl.server, &inferenceServer{s: s})

	return nil
}

func (impl serverImpl) RegisterEvaluateServer(s evaluate.EvaluateService) error {
	if s == nil {
		return errors.New("invliad service")
	}

	if impl.server == nil {
		return errors.New("no server")
	}

	protocol.RegisterEvaluateServer(impl.server, &evaluateServer{s: s})

	return nil
}

func (impl serverImpl) RegisterCompetitionServer(s competition.CompetitionService) error {
	if s == nil {
		return errors.New("invliad service")
	}

	if impl.server == nil {
		return errors.New("no server")
	}

	protocol.RegisterCompetitionServer(impl.server, &competitionServer{s: s})

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

// train
type trainingServer struct {
	s training.TrainingService
	protocol.UnimplementedTrainingServer
}

func (t *trainingServer) SetTrainingInfo(ctx context.Context, v *protocol.TrainingInfo) (
	*protocol.TrainingResult, error,
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

	// Must return new(protocol.Result), or grpc will failed.
	return new(protocol.TrainingResult), t.s.SetTrainingInfo(&index, &info)
}

// finetune
type finetuneServer struct {
	s finetune.FinetuneService
	protocol.UnimplementedFinetuneServer
}

func (t *finetuneServer) SetFinetuneInfo(ctx context.Context, v *protocol.FinetuneInfo) (
	*protocol.FinetuneResult, error,
) {
	index := finetune.FinetuneIndex{
		Id:   v.GetId(),
		User: v.GetUser(),
	}

	info := finetune.FinetuneInfo{
		Status:   v.GetStatus(),
		Duration: int(v.GetDuration()),
	}

	// Must return new(protocol.Result), or grpc will failed.
	return new(protocol.FinetuneResult), t.s.SetFinetuneInfo(&index, &info)
}

// inference
type inferenceServer struct {
	s inference.InferenceService
	protocol.UnimplementedInferenceServer
}

func (t *inferenceServer) SetInferenceInfo(ctx context.Context, v *protocol.InferenceInfo) (
	*protocol.InferenceResult, error,
) {
	index := inference.InferenceIndex{
		Id:         v.GetId(),
		User:       v.GetUser(),
		ProjectId:  v.GetProjectId(),
		LastCommit: v.GetLastCommit(),
	}

	info := inference.InferenceInfo{
		Error:     v.GetError(),
		AccessURL: v.GetAccessUrl(),
	}

	// Must return new(protocol.Result), or grpc will failed.
	return new(protocol.InferenceResult), t.s.SetInferenceInfo(&index, &info)
}

// evaluate
type evaluateServer struct {
	s evaluate.EvaluateService
	protocol.UnimplementedEvaluateServer
}

func (t *evaluateServer) SetEvaluateInfo(ctx context.Context, v *protocol.EvaluateInfo) (
	*protocol.EvaluateResult, error,
) {
	index := evaluate.EvaluateIndex{
		Id:         v.GetId(),
		User:       v.GetUser(),
		ProjectId:  v.GetProjectId(),
		TrainingID: v.GetTrainingId(),
	}

	info := evaluate.EvaluateInfo{
		Error:     v.GetError(),
		AccessURL: v.GetAccessUrl(),
	}

	// Must return new(protocol.Result), or grpc will failed.
	return new(protocol.EvaluateResult), t.s.SetEvaluateInfo(&index, &info)
}

// competition
type competitionServer struct {
	s competition.CompetitionService
	protocol.UnimplementedCompetitionServer
}

func (t *competitionServer) SetSubmissionInfo(ctx context.Context, v *protocol.SubmissionInfo) (
	*protocol.SubmissionResult, error,
) {
	info := competition.SubmissionInfo{
		Id:       v.GetId(),
		Status:   v.GetStatus(),
		Score:    v.GetScore(),
		Phase:    v.GetPhase(),
		PlayerId: v.GetPlayerId(),
	}

	// Must return new(protocol.Result), or grpc will failed.
	return new(protocol.SubmissionResult), t.s.SetSubmissionInfo(v.GetCompetitionId(), &info)
}
