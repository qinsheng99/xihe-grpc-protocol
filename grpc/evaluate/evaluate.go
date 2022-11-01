package evaluate

type EvaluateIndex struct {
	Id         string
	User       string
	ProjectId  string
	TrainingID string
}

type EvaluateInfo struct {
	Error     string
	AccessURL string
}

type EvaluateService interface {
	SetEvaluateInfo(*EvaluateIndex, *EvaluateInfo) error
}
