package competition

type CompetitionIndex struct {
	Id    string
	Phase string
}

type SubmissionInfo struct {
	Id     string
	Status string
	Score  float32
}

type CompetitionService interface {
	SetSubmissionInfo(*CompetitionIndex, *SubmissionInfo) error
}
