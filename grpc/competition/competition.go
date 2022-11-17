package competition

type SubmissionIndex struct {
	Id            string
	Phase         string
	CompetitionId string
}

type SubmissionInfo struct {
	Status string
	Score  float32
}

type CompetitionService interface {
	SetSubmissionInfo(*SubmissionIndex, *SubmissionInfo) error
}
