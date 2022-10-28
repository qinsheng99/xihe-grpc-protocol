package training

type TrainingIndex struct {
	Id        string
	User      string
	ProjectId string
}

type TrainingInfo struct {
	Duration      int
	Status        string
	LogPath       string
	AimZipPath    string
	OutputZipPath string
}

type TrainingService interface {
	SetTrainingInfo(*TrainingIndex, *TrainingInfo) error
}
