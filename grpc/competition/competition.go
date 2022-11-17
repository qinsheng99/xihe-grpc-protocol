package competition

type SubmissionIndex struct {
	Id            string
	Phase         string
	CompetitionId string
}

type SubmissionInfo struct {
	Score float32
}

type CompetitionService interface {
	SetSubmissionInfo(*SubmissionIndex, *SubmissionInfo) error
}
