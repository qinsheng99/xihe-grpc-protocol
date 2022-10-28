package server

type TrainingInfo struct {
	Id        string
	User      string
	ProjectId string
}

type TrainingOutput struct {
	AimPath       string
	OutputZipPath string
}

type TrainingService interface {
	SetTrainingStatus(*TrainingInfo, string) error
	SetTrainingOutput(*TrainingInfo, *TrainingOutput) error
	SetTrainingLogPath(*TrainingInfo, string) error
}
