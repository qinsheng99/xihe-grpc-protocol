package inference

type InferenceIndex struct {
	Id         string
	User       string
	ProjectId  string
	LastCommit string
}

type InferenceInfo struct {
	Error     string
	AccessURL string
}

type InferenceService interface {
	SetInferenceInfo(*InferenceIndex, *InferenceInfo) error
}
